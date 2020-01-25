package app

import (
	"context"
	"net"
	"testing"

	pb "github.com/muizidn/todo-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// https://qiita.com/castaneai/items/8f975204a79e9783ecc3

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	setupLogger()

	lis = bufconn.Listen(bufSize)
	s := testRegisterServers()
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

// You may want use testClientConnect2
func testClientConnect(t *testing.T, function func(context.Context, *grpc.ClientConn)) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	function(ctx, conn)
}

func testClientConnect2(t *testing.T) (context.Context, *grpc.ClientConn) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	return ctx, conn
}

func testRegisterServers() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &serverGreeter{})
	return s
}
