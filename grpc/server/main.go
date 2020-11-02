package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/musenwill/experimentgo/grpc/idl"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	portFlag := cli.IntFlag{
		Name:  "port",
		Value: 9000,
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
	verifyClientFlag := cli.BoolFlag{
		Name: "verify-client",
	}

	app := cli.NewApp()
	app.ErrWriter = os.Stdout
	app.EnableBashCompletion = true
	app.Name = "grpc server"
	app.Author = "musenwill"
	app.Email = "musenwill@qq.com"
	app.Flags = []cli.Flag{portFlag, tlsFlag, caFlag, certFlag, keyFlag, verifyClientFlag}
	app.Action = action

	app.RunAndExitOnError()
}

func action(c *cli.Context) error {
	port := c.Int("port")
	tls := c.Bool("tls")
	var s *grpc.Server
	if tls {
		tlsConfig, err := tlsConfig(c)
		if err != nil {
			return err
		}
		credent := credentials.NewTLS(tlsConfig)
		s = grpc.NewServer(grpc.Creds(credent))
	} else {
		s = grpc.NewServer()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	idl.RegisterDemoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func tlsConfig(c *cli.Context) (*tls.Config, error) {
	caPath := c.String("ca")
	certPath := c.String("cert")
	keyPath := c.String("key")
	verifyClient := c.Bool("verify-client")

	if caPath == "" {
		return nil, errors.New("ca file required")
	}
	if certPath == "" {
		return nil, errors.New("cert file required")
	}
	if keyPath == "" {
		return nil, errors.New("key file required")
	}

	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)

	tlsConfig := &tls.Config{}
	tlsConfig.ClientCAs = pool
	if verifyClient {
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}
	tlsConfig.Certificates = []tls.Certificate{cert}
	return tlsConfig, nil
}

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
		// log.Printf("received %d bytes", len(req.Data))
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
		// log.Printf("sent %d bytes", n)
	}

	return nil
}
