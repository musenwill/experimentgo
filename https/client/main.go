package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
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

	app := cli.NewApp()
	app.ErrWriter = os.Stdout
	app.EnableBashCompletion = true
	app.Name = "https client"
	app.Author = "musenwill"
	app.Email = "musenwill@qq.com"
	app.Flags = []cli.Flag{httpsFlag, caFlag, certFlag, keyFlag}
	app.Action = action

	app.RunAndExitOnError()
}

func action(c *cli.Context) error {
	https := c.Bool("https")
	addr := "http://localhost:8080"

	var client *http.Client
	if https {
		addr = "https://localhost:8080"
		tlsConfig, err := tlsConfig(c)
		if err != nil {
			return err
		}
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}
	} else {
		client = &http.Client{}
	}

	resp, err := client.Get(addr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	fmt.Println(resp.Status)
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

	return tlsConfig, nil
}
