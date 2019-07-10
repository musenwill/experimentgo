package sample

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRangeChannel(t *testing.T) {
	consumeChannel(func(msg <-chan string, msgCount, sleepMillisecond int) {
		count := 0
		for range msg {
			count++
			time.Sleep(time.Duration(sleepMillisecond) * time.Microsecond)
		}
		if exp, act := msgCount, count; exp != count {
			t.Errorf("failed receive all messages from channel, %v != %v", exp, act)
		}
	})
}

func TestForLoopChannel(t *testing.T) {
	consumeChannel(func(msg <-chan string, msgCount, sleepMillisecond int) {
		count := 0
		for {
			_, open := <-msg
			if !open {
				break
			}
			count++
			time.Sleep(time.Duration(sleepMillisecond) * time.Microsecond)
		}
		if exp, act := msgCount, count; exp != count {
			t.Errorf("failed receive all messages from channel, %v != %v", exp, act)
		}
	})
}

func TestSelectBlockChannel(t *testing.T) {
	consumeChannel(func(msg <-chan string, msgCount, sleepMillisecond int) {
		count := 0
	loop:
		for {
			select {
			case _, open := <-msg:
				if !open {
					break loop
				}
				count++
				time.Sleep(time.Duration(sleepMillisecond) * time.Microsecond)
			}
		}
		if exp, act := msgCount, count; exp != count {
			t.Errorf("failed receive all messages from channel, %v != %v", exp, act)
		}
	})
}

func TestSelectUnBlockChannel(t *testing.T) {
	consumeChannel(func(msg <-chan string, msgCount, sleepMillisecond int) {
		count := 0
	loop:
		for {
			select {
			case _, open := <-msg:
				if !open {
					break loop
				}
				count++
			default:
				time.Sleep(time.Duration(sleepMillisecond) * time.Microsecond)
			}
		}
		if exp, act := msgCount, count; exp != count {
			t.Errorf("failed receive all messages from channel, %v != %v", exp, act)
		}
	})
}

func consumeChannel(f func(msg <-chan string, msgCount, sleepMillisecond int)) {
	var msg = make(chan string, 10)
	var wg sync.WaitGroup
	msgCount := 5
	wg.Add(1)
	go func() {
		defer wg.Done()
		f(msg, msgCount, 2)
	}()

	for i := 0; i < msgCount; i++ {
		msg <- fmt.Sprintf("the %vth message", i)
		time.Sleep(1 * time.Millisecond)
	}
	close(msg)
	wg.Wait()
}
