package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/musenwill/experimentgo/grpc/idl"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var cache = make(map[string][]byte)

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

func (s *server) Write(ws idl.DemoService_WriteServer) error {
	data := []byte{}
	for {
		req, err := ws.Recv()
		if err == io.EOF {
			break
		}
		if nil != err {
			return errors.WithStack(err)
		}
		data = append(data, req.Data...)
		log.Printf("received %d bytes", len(req.Data))
	}
	id := fmt.Sprintf("%x", sha256.Sum256(data))
	cache[id] = data

	return ws.SendAndClose(&idl.WriteResponse{
		Code:   idl.WriteResponse_OK,
		Msg:    "",
		DataID: id,
	})
}

func (s *server) Read(in *idl.ReadRequest, rd idl.DemoService_ReadServer) error {
	data, ok := cache[in.DataID]
	if !ok {
		rd.Send(&idl.ReadResponse{
			Code: idl.ReadResponse_ERROR,
			Msg:  "no data",
		})
		return nil
	}

	buf := make([]byte, 8)
	reader := bytes.NewReader(data)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if nil != err {
			return errors.WithStack(err)
		}

		err = rd.Send(&idl.ReadResponse{
			Code: idl.ReadResponse_OK,
			Msg:  "",
			Data: buf[:n],
		})
		if nil != err {
			return errors.WithStack(err)
		}
		log.Printf("sent %d bytes", n)
	}

	return nil
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
