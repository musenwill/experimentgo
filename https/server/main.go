package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	httpsFlag := cli.BoolFlag{
		Name: "https",
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
	app.Name = "https server"
	app.Author = "musenwill"
	app.Email = "musenwill@qq.com"
	app.Flags = []cli.Flag{httpsFlag, caFlag, certFlag, keyFlag, verifyClientFlag}
	app.Action = action

	app.RunAndExitOnError()
}

func action(c *cli.Context) error {
	https := c.Bool("https")

	var err error
	var listener net.Listener
	var handler http.Handler

	if https {
		handler = &HttpsHandler{}
		tlsConfig, err := tlsConfig(c)
		if err != nil {
			return err
		}

		listener, err = tls.Listen("tcp", ":8080", tlsConfig)
		if err != nil {
			return err
		}
	} else {
		handler = &HttpHandler{}
		listener, err = net.Listen("tcp", ":8080")
		if err != nil {
			return err
		}
	}

	err = http.Serve(listener, handler)
	if err != nil {
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
	tlsConfig.ServerName = "falcontsdb"
	return tlsConfig, nil
}

type HttpHandler struct{}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello http")
}

type HttpsHandler struct{}

func (h *HttpsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello https")
}
