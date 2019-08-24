package impl

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func RandPut(cli *clientv3.Client, key string) error {
	rand.Seed(time.Now().UnixNano())
	for {
		if resp, err := cli.Put(context.TODO(), key, strconv.Itoa(rand.Int())); err != nil {
			return err
		} else {
			log.Println(resp)
		}
		time.Sleep(time.Duration(rand.Int()%2000+1000) * time.Millisecond)
	}
}

func RandDelete(cli *clientv3.Client, key string) error {
	rand.Seed(time.Now().UnixNano())
	for {
		if resp, err := cli.Delete(context.TODO(), key); err != nil {
			return err
		} else {
			log.Println(resp)
		}
		time.Sleep(time.Duration(rand.Int()%2000+1000) * time.Millisecond)
	}
}

func Watch(cli *clientv3.Client, key string) error {
	rch := cli.Watch(context.Background(), key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
	return nil
}
