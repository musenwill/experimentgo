package sample

import (
	"fmt"
	"time"
)

func pipeline() {
	var linears = make(chan int)
	var squares = make(chan int)

	go printer(squares)
	go squarer(squares, linears)
	counter(linears)
}

func counter(out chan<- int) {
	defer close(out)
	for index := 0; index < 10; index++ {
		out <- index
		time.Sleep(200 * time.Millisecond)
	}
}

func squarer(out chan<- int, in <-chan int) {
	defer close(out)
	for {
		lin, ok := <-in
		if !ok {
			break
		}
		out <- lin * lin
	}
}

func printer(in <-chan int) {
	for {
		squ := <-in
		fmt.Print("\r", squ)
	}
	fmt.Println()
}
