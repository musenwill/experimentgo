package sample

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* for range 会保证 channel 的数据全部读取完 */
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

/* 只要还有空间，buffered channel 就可以无阻塞的继续写入 */
func TestWriteChannelNoneBlock(t *testing.T) {
	timer := time.NewTimer(2 * time.Millisecond)
	val := make(chan int, 1)
	timeout := false

loop:
	for i := 0; i < 1; i++ {
		select {
		case val <- i:
		case <-timer.C:
			timeout = true
			break loop
		}
	}
	close(val)

	if exp, act := 1, len(val); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
	if exp, act := false, timeout; exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}

/* 对于 channel 来说，只要没有接收者，发送者就会阻塞，它并不是 size = 1 的 buffered channel */
func TestWriteChannelBlock(t *testing.T) {
	timer := time.NewTimer(2 * time.Millisecond)
	val := make(chan int)
	timeout := false

loop:
	for i := 0; i < 1; i++ {
		select {
		case val <- i:
		case <-timer.C:
			timeout = true
			break loop
		}
	}
	close(val)

	if exp, act := true, timeout; exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}
