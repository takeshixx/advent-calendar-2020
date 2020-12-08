package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	xg "github.com/takeshixx/xmasgreetings/xmasgreetings"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:50051"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}
	name := "xmas"
	if len(os.Args) > 2 {
		name = os.Args[2]
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := xg.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.XmasGreeting(ctx, &xg.XmasRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.GetMessage())
}
