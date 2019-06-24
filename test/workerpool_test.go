package test

import (
	"fmt"
	"sync"
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

func BenchmarkWorkerPool(b *testing.B) {
	workers := api.NewWorkerPool(1000)
	for i := 0; i < b.N; i++ {
		workers.Excute(func() {
		})
	}
	workers.Wait(nil)
}

func ExampleWorkerPool() {
	workers := api.NewWorkerPool(10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			workers.Excute(func() {
				// do something here
			})
		}
	}()
	wg.Wait()
	workers.Wait(func() { fmt.Println("workers finished") })
	// Output:
	// workers finished
}
