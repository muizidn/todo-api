package app

import (
	"testing"

	pb "github.com/muizidn/todo-api/pb"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	ctx, conn := testClientConnect2(t)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "test"})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Hello test", resp.GetMessage())
}
