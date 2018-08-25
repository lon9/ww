package server

import (
	"log"
	"net"

	"github.com/lon9/ww/game"
	pb "github.com/lon9/ww/proto"
	xcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NumPlayers is the number of players
const NumPlayers = 1

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
	people              []game.Personer
	state               pb.State
	connectionQueue     chan *connectionEntry
	stateQueue          chan chan bool
	stateBroadcastChans []chan bool
}

// NewTestServer constructor for test server
func NewTestServer() *Server {
	return &Server{
		state:           pb.State_AFTER,
		connectionQueue: make(chan *connectionEntry),
		stateQueue:      make(chan chan bool),
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
		numConnected := 0
		for entry := range s.connectionQueue {
			player := game.NewPersoner(numConnected, entry.Name, pb.Kind_CITIZEN)
			s.people = append(s.people, player)
			entry.ResChan <- player
			numConnected++
			log.Printf("%s is connected", player.GetName())
			if numConnected == NumPlayers {
				break
			}
		}
	}()

	// Waiting for state request
	go func() {
		defer close(s.stateQueue)
		numConnected := 0
		for ch := range s.stateQueue {
			s.stateBroadcastChans = append(s.stateBroadcastChans, ch)
			numConnected++
			if numConnected == NumPlayers {
				break
			}
		}
		log.Println("All players ready")
		s.changeState(pb.State_NIGHT)
	}()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
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
	for range ch {
		res := &pb.StateResponse{
			State:   s.state,
			Players: s.convertPeople2Players(),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

// Bite handles Bite request
func (s *Server) Bite(ctx xcontext.Context, req *pb.BiteRequest) (*pb.BiteResponse, error) {
	return nil, nil
}

// Vote handles Vote request
func (s *Server) Vote(ctx xcontext.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	return nil, nil
}

// Protect handles Protect request
func (s *Server) Protect(ctx xcontext.Context, req *pb.ProtectRequest) (*pb.ProtectResponse, error) {
	return nil, nil
}

// Tell handles Tell request
func (s *Server) Tell(ctx xcontext.Context, req *pb.TellRequest) (*pb.TellResponse, error) {
	return nil, nil
}

// Sleep handles Sleep request
func (s *Server) Sleep(ctx xcontext.Context, req *pb.SleepRequest) (*pb.SleepResponse, error) {
	return nil, nil
}

func (s *Server) convertPeople2Players() []*pb.Player {
	var players []*pb.Player
	for _, v := range s.people {
		player := &pb.Player{
			Id:     int32(v.GetID()),
			Name:   v.GetName(),
			IsDead: v.GetIsDead(),
		}
		players = append(players, player)
	}
	return players
}

func (s *Server) changeState(state pb.State) {
	s.state = state
	for i := range s.stateBroadcastChans {
		s.stateBroadcastChans[i] <- true
	}
}
