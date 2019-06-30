package sample

import (
	"fmt"
	"time"
)

// CountDown ...
func CountDown() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("\r%2d", countdown)
		<-tick
	}
	fmt.Println("launch")
}
