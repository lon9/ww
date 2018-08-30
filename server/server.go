package server

import (
	"log"
	"math/rand"
	"net"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lon9/ww/consts"
	"github.com/lon9/ww/game"
	pb "github.com/lon9/ww/proto"
	xcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// connectionEntry is used to decide player's job
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
	restartVote         int
	grpcServer          *grpc.Server
}

// NewTestServer constructor for test server
func NewTestServer() *Server {
	return &Server{
		actionMutex: new(sync.Mutex),
	}
}

// Run runs server
func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.grpcServer = grpc.NewServer()
	pb.RegisterWWServer(s.grpcServer, s)
	// Register reflection service on gRPC server.
	// reflection.Register(grpcServer)

	// Start
	s.start()

	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// start set settings of the game and wait for connections
func (s *Server) start() {

	// Initialize server
	indice := make([]int, consts.NumPlayers)
	for i := 0; i < consts.NumPlayers; i++ {
		indice[i] = i
	}
	n := len(indice)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		indice[i], indice[j] = indice[j], indice[i]
	}

	personers := make(game.Personers)
	var idx int
	for i := 0; i < consts.NumWarewolf; i++ {
		personers[indice[idx]] = game.NewPersoner(indice[idx], "", pb.Kind_WAREWOLF)
		idx++
	}
	for i := 0; i < consts.NumTeller; i++ {
		personers[indice[idx]] = game.NewPersoner(indice[idx], "", pb.Kind_TELLER)
		idx++
	}
	for i := 0; i < consts.NumKnight; i++ {
		personers[indice[idx]] = game.NewPersoner(indice[idx], "", pb.Kind_KNIGHT)
		idx++
	}
	for i := 0; i < consts.NumPlayers-consts.NumWarewolf-consts.NumKnight-consts.NumTeller; i++ {
		personers[indice[idx]] = game.NewPersoner(indice[idx], "", pb.Kind_CITIZEN)
		idx++
	}

	s.personers = personers
	s.state = pb.State_BEFORE
	s.connectionQueue = make(chan *connectionEntry)
	s.stateQueue = make(chan chan bool)
	s.restartVote = 0

	// Waiting for hello request
	go func() {
		defer close(s.connectionQueue)
		for i := 0; i < consts.NumPlayers; i++ {
			entry := <-s.connectionQueue
			s.personers[i].SetName(entry.Name)
			entry.ResChan <- s.personers[i]
			log.Printf("%s is connected", entry.Name)
		}
	}()

	// Waiting for state request
	go func() {
		defer close(s.stateQueue)
		for i := 0; i < consts.NumPlayers; i++ {
			ch := <-s.stateQueue
			s.stateBroadcastChans = append(s.stateBroadcastChans, ch)
		}
		log.Println("All players ready")
		s.changeState(pb.State_NIGHT)
		go s.gameLoop()
	}()
}

// reasign reasigns player's job
func (s *Server) reasign() {
	s.restartVote = 0
	n := len(s.personers)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		s.personers[i], s.personers[j] = s.personers[j], s.personers[i]
	}
	for k := range s.personers {
		s.personers[k].SetID(k)
		s.personers[k].SetIsDead(false)
	}
}

// gameLoop sync game state with clients
func (s *Server) gameLoop() {
	for {
		// Waiting responses from clients
		s.finishActionCh = make(chan string, consts.NumPlayers)
		personMap := make(map[string]bool, consts.NumPlayers)
		for _, v := range s.personers {
			personMap[v.GetUUID().String()] = false
		}
		for i := 0; i < consts.NumPlayers; i++ {
			uuid := <-s.finishActionCh
			if v, ok := personMap[uuid]; !ok || v {
				return
			}
			personMap[uuid] = true
		}
		close(s.finishActionCh)

		if s.state == pb.State_AFTER {
			if s.restartVote == consts.NumPlayers {
				// If all player want to restart, restart the server.
				log.Println("Restarting...")
				s.reasign()
				s.changeState(pb.State_NIGHT)
				continue
			}
			log.Println("Shutting down...")
			s.grpcServer.Stop()
			return
		}

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
			continue
		}

		// Transition to next state
		switch s.state {
		case pb.State_BEFORE:
			s.changeState(pb.State_NIGHT)
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
	personer, err := s.personers.FindPersonerByUUID(req.GetUuid())
	if err != nil {
		return err
	}
	for range ch {
		res := &pb.StateResponse{
			State: s.state,
		}
		if s.state == pb.State_AFTER {

			// If the state is after, send all players properties
			res.Players = personer.ConvertAfter(s.personers)
		} else {
			res.Players = personer.ConvertPersoners(s.personers)
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
	s.finishActionCh <- req.GetSrcUuid()
	return new(pb.BiteResponse), nil
}

// Vote handles Vote request
func (s *Server) Vote(ctx xcontext.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	s.actionMutex.Lock()
	id := int(req.GetDstId())
	s.personers[id].IncVotes()
	s.actionMutex.Unlock()
	s.finishActionCh <- req.GetSrcUuid()
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
	s.finishActionCh <- req.GetSrcUuid()
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
	s.finishActionCh <- req.GetSrcUuid()
	return &pb.TellResponse{
		Camp: s.personers[id].GetCamp(),
	}, nil
}

// Sleep handles Sleep request
func (s *Server) Sleep(ctx xcontext.Context, req *pb.SleepRequest) (*pb.SleepResponse, error) {
	s.finishActionCh <- req.GetSrcUuid()
	return new(pb.SleepResponse), nil
}

// Dead handles Dead request
func (s *Server) Dead(ctx xcontext.Context, req *pb.DeadRequest) (*pb.DeadResponse, error) {
	s.finishActionCh <- req.GetSrcUuid()
	return new(pb.DeadResponse), nil
}

// Restart handles Restart request
func (s *Server) Restart(ctx xcontext.Context, req *pb.RestartRequest) (*pb.RestartResponse, error) {
	s.actionMutex.Lock()
	if req.GetIsRestart() {
		s.restartVote++
	}
	s.actionMutex.Unlock()
	s.finishActionCh <- req.GetSrcUuid()
	return new(pb.RestartResponse), nil
}

func (s *Server) changeState(state pb.State) {
	// Initializes before changing state
	log.Println(state)
	s.personers.Init()
	s.state = state
	for i := range s.stateBroadcastChans {
		s.stateBroadcastChans[i] <- true
	}
}
