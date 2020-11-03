package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/musenwill/experimentgo/grpc/idl"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	addrFlag := cli.StringFlag{
		Name:  "addr",
		Usage: "server address, host:port",
		Value: "localhost:9000",
	}
	msgFlag := cli.StringFlag{
		Name:  "msg",
		Usage: "message send to server",
		Value: "how are you?",
	}
	tlsFlag := cli.BoolFlag{
		Name: "tls",
	}
	caFlag := cli.StringFlag{
		Name: "ca",
	}
	certFlag := cli.StringFlag{
		Name: "cert",
	}
	keyFlag := cli.StringFlag{
		Name: "key",
	}

	app := cli.NewApp()
	app.ErrWriter = os.Stdout
	app.EnableBashCompletion = true
	app.Name = "grpc client"
	app.Author = "musenwill"
	app.Email = "musenwill@qq.com"
	app.Flags = []cli.Flag{addrFlag, msgFlag, tlsFlag, caFlag, certFlag, keyFlag}
	app.Action = action

	app.RunAndExitOnError()
}

func action(c *cli.Context) error {
	addr := c.String("addr")
	msg := c.String("msg")
	tls := c.Bool("tls")

	var dialOpt grpc.DialOption
	if tls {
		tlsConfig, err := tlsConfig(c)
		if err != nil {
			return err
		}
		dialOpt = grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))
	} else {
		dialOpt = grpc.WithInsecure()
	}

	conn, err := grpc.Dial(addr, dialOpt)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := idl.NewDemoServiceClient(conn)

	hello(client, msg)
	dataID := write(client)
	fmt.Printf("data id: %s\n", dataID)
	read(client, dataID)
	return nil
}

func tlsConfig(c *cli.Context) (*tls.Config, error) {
	caPath := c.String("ca")
	certPath := c.String("cert")
	keyPath := c.String("key")

	var pool *x509.CertPool
	if caPath != "" {
		ca, err := ioutil.ReadFile(caPath)
		if err != nil {
			return nil, err
		}
		pool = x509.NewCertPool()
		pool.AppendCertsFromPEM(ca)
	}

	var certs []tls.Certificate
	if certPath != "" && keyPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, err
		}
		certs = []tls.Certificate{cert}
	}

	tlsConfig := &tls.Config{
		RootCAs:            pool,
		Certificates:       certs,
		ServerName:         "falcontsdb",
		InsecureSkipVerify: true,
	}
	tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	return tlsConfig, nil
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
		// log.Printf("write %d bytes", n)
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
		// log.Printf("received %d bytes", len(resp.Data))
	}
	fmt.Printf("data: %s\n", string(data))
}
