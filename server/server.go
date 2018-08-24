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

const NumPlayers = 10
const NumWarewolf = 2
const NumTeller = 1
const NumKnight = 1

type Server struct {
	people []game.Personer
}

func NewTestServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterWWServer(grpcServer, s)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Hello(ctx xcontext.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	player := game.NewPersoner(1, req.GetName(), pb.Kind_CITIZEN)
	s.people = append(s.people, player)
	res := &pb.HelloResponse{
		Uuid:  player.GetUUID().String(),
		Name:  player.GetName(),
		Kind:  player.GetKind(),
		State: s.getStateFromPeople(),
	}
	return res, nil
}

func (s *Server) Bite(ctx xcontext.Context, req *pb.BiteRequest) (*pb.BiteResponse, error) {
	return nil, nil
}

func (s *Server) Vote(ctx xcontext.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	return nil, nil
}

func (s *Server) Protect(ctx xcontext.Context, req *pb.ProtectRequest) (*pb.ProtectResponse, error) {
	return nil, nil
}

func (s *Server) Tell(ctx xcontext.Context, req *pb.TellRequest) (*pb.TellResponse, error) {
	return nil, nil
}

func (s *Server) Sleep(ctx xcontext.Context, req *pb.SleepRequest) (*pb.SleepResponse, error) {
	return nil, nil
}

func (s *Server) getStateFromPeople() *pb.State {
	state := new(pb.State)
	for _, v := range s.people {
		player := &pb.Player{
			Id:     int32(v.GetID()),
			Name:   v.GetName(),
			IsDead: v.GetIsDead(),
		}
		state.Players = append(state.Players, player)
	}
	return state
}
