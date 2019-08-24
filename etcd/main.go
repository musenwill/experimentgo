package main

import (
	"log"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/musenwill/experimentgo/etcd/impl"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 1 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		cli.Close()
		log.Println("close etcd connection")
	}()

	log.Println("connected to etcd")

	key := "foo"
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		impl.RandPut(cli, key)
	}()
	go func() {
		defer wg.Done()
		impl.Watch(cli, key)
	}()
	go func() {
		defer wg.Done()
		impl.RandDelete(cli, key)
	}()

	wg.Wait()
}
