package server

import (
	"strconv"
	"sync"
	"testing"

	pb "github.com/lon9/ww/proto"
	xcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func prepareServer() {
	s := NewTestServer()
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
			t.Log(res)
			wg.Done()
		}(i, t)
	}
	wg.Wait()
}
