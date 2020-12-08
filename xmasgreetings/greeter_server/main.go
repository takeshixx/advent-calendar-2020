package main

import (
	"context"
	"log"
	"net"
	"os"

	xg "github.com/takeshixx/xmasgreetings/xmasgreetings"
	"google.golang.org/grpc"
)

const (
	port    = ":50051"
	message = `We wish you a merry Christmas
We wish you a merry Christmas
We wish you a merry Christmas and a happy new secret: `
)

type server struct {
	xg.UnimplementedGreeterServer
}

func (s *server) XmasGreeting(ctx context.Context, in *xg.XmasRequest) (*xg.XmasReply, error) {
	if in.GetName() == "xmas" {
		return &xg.XmasReply{Message: message + os.Getenv("XMAS_SECRET")}, nil
	}
	return &xg.XmasReply{Message: "What is " + in.GetName() + "?!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	xg.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
