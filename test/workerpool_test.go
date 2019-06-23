package test

import (
	"testing"

	"github.com/musenwill/experimentgo/api"
)

func TestWorkerPoolNormal(t *testing.T) {
	workers := api.NewWorkerPool(10)
	for i := 0; i < 100; i++ {
		index := i
		workers.Excute(func() {
			t.Log("worker num ", index)
		})
	}
	workers.Wait(func() { t.Log("workers finished") })
}

func TestWorkerPoolManyWork(t *testing.T) {
	workers := api.NewWorkerPool(1000)
	for i := 0; i < 100000; i++ {
		index := i
		workers.Excute(func() {
			t.Log("worker num ", index)
		})
	}
	workers.Wait(func() { t.Log("workers finished") })
}
