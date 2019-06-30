package sample

import (
	"fmt"
	"time"
)

// Spinner ...
func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
