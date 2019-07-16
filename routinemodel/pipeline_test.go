package routinemodel

import (
	"testing"
)

func TestPipeLine(t *testing.T) {
	x, y, z := make(chan int), make(chan int), make(chan int)
	go func() {
		for range z {
			// fmt.Println(v)
		}
	}()
	go func() {
		for v := range y {
			z <- v * v
		}
		close(z)
	}()
	go func() {
		for v := range x {
			y <- v + 1
		}
		close(y)
	}()

	for i := 0; i < 10; i++ {
		x <- i
	}
	close(x)
}
