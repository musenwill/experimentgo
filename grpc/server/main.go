package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/musenwill/experimentgo/grpc/idl"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) Hello(ctx context.Context, in *idl.HelloRequest) (*idl.HelloResponse, error) {
	log.Printf("received: %v", in.Greet)
	return &idl.HelloResponse{
		Code:  idl.HelloResponse_OK,
		Msg:   "",
		Reply: "welcome " + in.Greet,
	}, nil
}

func main() {
	var port int
	flag.IntVar(&port, "port", 9000, "set port of server")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	idl.RegisterDemoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
