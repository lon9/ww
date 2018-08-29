package server

import (
	"io"
	"strconv"
	"sync"
	"testing"

	"github.com/lon9/ww/consts"
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
	for i := 0; i < consts.NumPlayers; i++ {
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

			stream, err := client.State(xcontext.Background(), &pb.StateRequest{
				Uuid: res.Uuid,
			})
			if err != nil {
				t.Error(err)
			}
			defer stream.CloseSend()
			stateRes, err := stream.Recv()
			if err != nil && err != io.EOF {
				t.Error(err)
			}
			if len(stateRes.GetPlayers()) != consts.NumPlayers {
				t.Errorf("The number of players should be %d: %d", consts.NumPlayers, len(stateRes.GetPlayers()))
			}
			if stateRes.GetState() != pb.State_NIGHT {
				t.Errorf("State should be night: %v", stateRes.GetState())
			}

			if res.GetKind() == pb.Kind_WAREWOLF {
				var wwCount int
				for _, v := range stateRes.GetPlayers() {
					if v.GetKind() == pb.Kind_WAREWOLF {
						wwCount++
					}
				}
				if wwCount != 2 {
					t.Error("should be send warewolf kind to warewolf")
				}
			}

			wg.Done()
		}(i, t)
	}
	wg.Wait()
}
