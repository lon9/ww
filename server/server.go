package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lon9/ww/game"
	pb "github.com/lon9/ww/proto"
	xcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NumPlayers is the number of players
const NumPlayers = 10

// NumWarewolf is the number of warewolfs
const NumWarewolf = 2

// NumTeller is the number of fortune tellers
const NumTeller = 1

// NumKnight is the number of knights
const NumKnight = 1

type connectionEntry struct {
	Name    string
	ResChan chan game.Personer
}

// Server is struct for server
type Server struct {
	personers           game.Personers
	state               pb.State
	connectionQueue     chan *connectionEntry
	stateQueue          chan chan bool
	stateBroadcastChans []chan bool
	finishActionCh      chan string
	actionMutex         *sync.Mutex
}

// NewTestServer constructor for test server
func NewTestServer() *Server {

	s := make([]int, NumPlayers)
	for i := 0; i < NumPlayers; i++ {
		s[i] = i
	}
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}

	personers := make(game.Personers)
	var idx int
	for i := 0; i < NumWarewolf; i++ {
		personers[s[idx]] = game.NewPersoner(s[idx], "", pb.Kind_WAREWOLF)
		idx++
	}
	for i := 0; i < NumTeller; i++ {
		personers[s[idx]] = game.NewPersoner(s[idx], "", pb.Kind_TELLER)
		idx++
	}
	for i := 0; i < NumKnight; i++ {
		personers[s[idx]] = game.NewPersoner(s[idx], "", pb.Kind_KNIGHT)
		idx++
	}
	for i := 0; i < NumPlayers-NumWarewolf-NumKnight-NumTeller; i++ {
		personers[s[idx]] = game.NewPersoner(s[idx], "", pb.Kind_CITIZEN)
		idx++
	}

	return &Server{
		personers:       personers,
		state:           pb.State_BEFORE,
		connectionQueue: make(chan *connectionEntry),
		stateQueue:      make(chan chan bool),
		actionMutex:     new(sync.Mutex),
	}
}

// Run runs server
func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterWWServer(grpcServer, s)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Waiting for hello request
	go func() {
		defer close(s.connectionQueue)
		var numConnected int
		for entry := range s.connectionQueue {
			s.personers[numConnected].SetName(entry.Name)
			entry.ResChan <- s.personers[numConnected]
			numConnected++
			log.Printf("%s is connected", entry.Name)
			if numConnected == NumPlayers {
				break
			}
		}
	}()

	// Waiting for state request
	go func() {
		defer close(s.stateQueue)
		var numConnected int
		for ch := range s.stateQueue {
			s.stateBroadcastChans = append(s.stateBroadcastChans, ch)
			numConnected++
			if numConnected == NumPlayers {
				break
			}
		}
		log.Println("All players ready")
		s.changeState(pb.State_NIGHT)
		go s.gameLoop()
	}()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) gameLoop() {
	for {
		// Waiting responses from clients
		numAlive := s.personers.NumAlive()
		s.finishActionCh = make(chan string, numAlive)
		personMap := make(map[string]bool, numAlive)
		for _, v := range s.personers {
			if !v.GetIsDead() {
				personMap[v.GetUUID().String()] = false
			}
		}
		for i := 0; i < numAlive; i++ {
			uuid := <-s.finishActionCh
			if v, ok := personMap[uuid]; !ok || v {
				return
			}
			personMap[uuid] = true
		}
		close(s.finishActionCh)

		// Decide dead or alive of the players
		switch s.state {
		case pb.State_NIGHT:
			s.personers.ResolveDeadOrAliveAtNight()
		case pb.State_MORNING:
			s.personers.ResolveDeadOrAliveAtMorning()
		}

		// If the game is finished, transition to after state
		if s.personers.IsFinish() {
			s.changeState(pb.State_AFTER)
			return
		}

		switch s.state {
		case pb.State_BEFORE:
			s.changeState(pb.State_MORNING)
		case pb.State_MORNING:
			s.changeState(pb.State_NIGHT)
		case pb.State_NIGHT:
			s.changeState(pb.State_MORNING)
		}
	}
}

// Hello handles Hello request
func (s *Server) Hello(ctx xcontext.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	entry := &connectionEntry{
		Name:    req.GetName(),
		ResChan: make(chan game.Personer),
	}
	defer close(entry.ResChan)
	s.connectionQueue <- entry
	player := <-entry.ResChan
	fmt.Println(int32(player.GetID()))
	fmt.Println(player.GetKind())
	res := &pb.HelloResponse{
		Id:   int32(player.GetID()),
		Uuid: player.GetUUID().String(),
		Name: player.GetName(),
		Kind: player.GetKind(),
	}
	return res, nil
}

// State handles State request
func (s *Server) State(req *pb.StateRequest, stream pb.WW_StateServer) error {
	ch := make(chan bool)
	defer close(ch)
	s.stateQueue <- ch
	for range ch {
		res := &pb.StateResponse{
			State:   s.state,
			Players: s.personers.ConvertPersoners(),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

// Bite handles Bite request
func (s *Server) Bite(ctx xcontext.Context, req *pb.BiteRequest) (*pb.BiteResponse, error) {
	s.actionMutex.Lock()
	if !s.personers.ValidKind(req.GetSrcUuid(), pb.Kind_WAREWOLF) {
		s.actionMutex.Unlock()
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	id := int(req.GetDstId())
	s.personers[id].IncDeadWill()
	s.actionMutex.Unlock()
	return new(pb.BiteResponse), nil
}

// Vote handles Vote request
func (s *Server) Vote(ctx xcontext.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	s.actionMutex.Lock()
	id := int(req.GetDstId())
	s.personers[id].IncVotes()
	s.actionMutex.Unlock()
	return new(pb.VoteResponse), nil
}

// Protect handles Protect request
func (s *Server) Protect(ctx xcontext.Context, req *pb.ProtectRequest) (*pb.ProtectResponse, error) {
	s.actionMutex.Lock()
	if !s.personers.ValidKind(req.GetSrcUuid(), pb.Kind_KNIGHT) {
		s.actionMutex.Unlock()
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	id := int(req.GetDstId())
	s.personers[id].IncAliveWill()
	s.actionMutex.Unlock()
	return new(pb.ProtectResponse), nil
}

// Tell handles Tell request
func (s *Server) Tell(ctx xcontext.Context, req *pb.TellRequest) (*pb.TellResponse, error) {
	s.actionMutex.Lock()
	if !s.personers.ValidKind(req.GetSrcUuid(), pb.Kind_TELLER) {
		s.actionMutex.Unlock()
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	id := int(req.GetDstId())
	s.actionMutex.Unlock()
	return &pb.TellResponse{
		Camp: s.personers[id].GetCamp(),
	}, nil
}

// Sleep handles Sleep request
func (s *Server) Sleep(ctx xcontext.Context, req *pb.SleepRequest) (*pb.SleepResponse, error) {
	return new(pb.SleepResponse), nil
}

func (s *Server) changeState(state pb.State) {
	// Initializes before changing state
	s.personers.Init()
	s.state = state
	for i := range s.stateBroadcastChans {
		s.stateBroadcastChans[i] <- true
	}
}
