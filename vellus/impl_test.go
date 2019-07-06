package vellus

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestVellusNormal(t *testing.T) {
	workers := NewVellus(10)
	for i := 0; i < 100; i++ {
		index := i
		workers.Excute(func() {
			t.Log("worker num ", index)
		})
	}
	timer := time.NewTimer(500 * time.Millisecond)
	go func() {
		<-timer.C
		t.Error("Wait expected to have returned")
	}()
	workers.Wait(func() { t.Log("workers finished") })
	timer.Stop()
}

func BenchmarkWorkerPool(b *testing.B) {
	workers := NewVellus(1000)
	for i := 0; i < b.N; i++ {
		workers.Excute(func() {
		})
	}
	workers.Wait(nil)
}

func ExampleVellus() {
	workers := NewVellus(10)
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
