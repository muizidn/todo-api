package app

import (
	"database/sql"
	"fmt"
	"net"

	pb "github.com/muizidn/todo-api/pb"
	"google.golang.org/grpc"
)

func setupGrpcServer(port int, db *sql.DB) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &serverGreeter{})

	log.Println("server prepared")
	log.Printf("serving at localhost:%d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
