package server

import (
	"io"
	"strconv"
	"sync"
	"testing"

	pb "github.com/lon9/ww/proto"
	xcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var s *Server

func prepareServer() {
	s = NewTestServer()
	go s.Run("9999")
}

func TestConnection(t *testing.T) {
	prepareServer()
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int, t *testing.T) {
			conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
			if err != nil {
				t.Error(err)
			}
			defer conn.Close()
			client := pb.NewWWClient(conn)
			req := &pb.HelloRequest{
				Name: strconv.Itoa(idx),
			}
			res, err := client.Hello(xcontext.Background(), req)
			if err != nil {
				t.Error(err)
			}
			if res.GetName() != req.GetName() {
				t.Errorf("Name is invalid %s:%s", res.GetName(), req.GetName())
			}

			stream, err := client.State(xcontext.Background(), new(pb.StateRequest))
			if err != nil {
				t.Error(err)
			}
			defer stream.CloseSend()
			stateRes, err := stream.Recv()
			if err != nil && err != io.EOF {
				t.Error(err)
			}
			if len(stateRes.GetPlayers()) != 10 {
				t.Errorf("The number of players should be 10: %d", len(stateRes.GetPlayers()))
			}
			if stateRes.GetState() != pb.State_NIGHT {
				t.Errorf("State should be night: %v", stateRes.GetState())
			}

			wg.Done()
		}(i, t)
	}
	wg.Wait()
}

func TestBiteOne(t *testing.T) {
	prepareServer()
}
