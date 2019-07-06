package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/musenwill/experimentgo/grpc/idl"
	"google.golang.org/grpc"
)

func main() {
	var server string
	var msg string
	flag.StringVar(&server, "server", "localhost:9000", "server address")
	flag.StringVar(&msg, "message", "musenwill", "message send to server")
	flag.Parse()

	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := idl.NewDemoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Hello(ctx, &idl.HelloRequest{Greet: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reply: %s", r.Reply)
}
