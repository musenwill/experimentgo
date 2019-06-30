package sample

import (
	"fmt"
	"time"
)

// CoinWithSelect ...
func CoinWithSelect() {
	coin := make(chan int, 2)
	for i := 0; i < 200; i++ {
		select {
		case x := <-coin:
			{
				fmt.Print(x & 0x01)
			}
		case coin <- i:
		}
	}
}

// CoinWithRoutine ...
func CoinWithRoutine() {
	c := make(chan int)
	go func() {
		for {
			c <- 0
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c <- 1
			time.Sleep(50 * time.Millisecond)
		}
	}()
	for i := range c {
		fmt.Print(i)
	}
}
