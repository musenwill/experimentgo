package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"

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

	hello(c, msg)
	dataID := write(c)
	fmt.Printf("data id: %s\n", dataID)
	read(c, dataID)
}

func hello(c idl.DemoServiceClient, msg string) {
	r, err := c.Hello(context.Background(), &idl.HelloRequest{Greet: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reply: %s", r.Reply)
}

func write(c idl.DemoServiceClient) string {
	data := []byte("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abc")
	wc, err := c.Write(context.Background())
	if err != nil {
		log.Fatalf("failed to write data: %v", err)
	}

	buf := make([]byte, 8)
	reader := bytes.NewReader(data)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		err = wc.Send(&idl.WriteRequest{Data: buf[:n]})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("write %d bytes", n)
	}
	resp, err := wc.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to write data: %v", err)
	}
	if resp.Code != idl.WriteResponse_OK {
		fmt.Printf("response: %#v\n", resp)
		return ""
	}
	return resp.DataID
}

func read(c idl.DemoServiceClient, dataID string) {
	rc, err := c.Read(context.Background(), &idl.ReadRequest{DataID: dataID})
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	for {
		resp, err := rc.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if resp.Code != idl.ReadResponse_OK {
			fmt.Printf("response: %#v\n", resp)
			return
		}
		data = append(data, resp.Data...)
		log.Printf("received %d bytes", len(resp.Data))
	}
	fmt.Printf("data: %s\n", string(data))
}
