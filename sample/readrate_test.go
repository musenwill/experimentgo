package sample

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
const read value 100 concurrently 10000 per routine cost 114.117µs
rwlock read value 100 concurrently 10000 per routine cost 30.362191ms
channel read value 100 concurrently 10000 per routine cost 255.886359ms
buffered channel read value 100 concurrently 10000 per routine cost 95.387394ms

下面四组测试中，都是计算 100 个 goroutine 各自循环读取数据 10000 次所需要的时间。
第一组测试，是直接读取常量，没有任何并发同步问题。
第二组测试，是通过读锁来控制并发问题。
第三组测试，是通过 channel 来控制并发问题。
第四组测试，是通过 buffered channel 来控制并发问题。

很显然的，不用考虑并发控制时，并发读数据的性能是最高的，只需要 114µs
使用读写锁时，性能下降很严重，需要 30ms，慢了 266 倍
使用 channel 时，因为 goroutine 都被串行化了，速度更慢，需要 256ms，比读写锁慢 8 倍
使用大小为 100 的 buffered channel 进行改善，需要 95.387394ms，改善很明星，但还是比读写锁慢 3 倍
*/

func TestConcurrentReadConst(t *testing.T) {
	const val = 99
	start := make(chan struct{})
	var wg = &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 10000; j++ {
				if exp, act := 99, val; exp != act {
					t.Errorf("got %v expected %v", act, exp)
				}
			}
		}()
	}

	s := time.Now()
	close(start)
	wg.Wait()
	fmt.Printf("const read value %v concurrently %v per routine cost %v\n", 100, 10000, time.Since(s))
}

func TestConcurrentReadRWLock(t *testing.T) {
	const val = 99
	var lock = &sync.RWMutex{}

	start := make(chan struct{})
	var wg = &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 10000; j++ {
				lock.RLock()
				v := val
				lock.RUnlock()
				if exp, act := 99, v; exp != act {
					t.Errorf("got %v expected %v", act, exp)
				}
			}
		}()
	}

	s := time.Now()
	close(start)
	wg.Wait()
	fmt.Printf("rwlock read value %v concurrently %v per routine cost %v\n", 100, 10000, time.Since(s))
}

func TestConcurrentReadChannel(t *testing.T) {
	val := make(chan int)
	go func() {
		for {
			val <- 99
		}
	}()

	start := make(chan struct{})
	var wg = &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 10000; j++ {
				if exp, act := 99, <-val; exp != act {
					t.Errorf("got %v expected %v", act, exp)
				}
			}
		}()
	}

	s := time.Now()
	close(start)
	wg.Wait()
	fmt.Printf("channel read value %v concurrently %v per routine cost %v\n", 100, 10000, time.Since(s))
}

func TestConcurrentReadBufferedChannel(t *testing.T) {
	val := make(chan int, 100)
	go func() {
		for {
			val <- 99
		}
	}()

	start := make(chan struct{})
	var wg = &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 10000; j++ {
				if exp, act := 99, <-val; exp != act {
					t.Errorf("got %v expected %v", act, exp)
				}
			}
		}()
	}

	s := time.Now()
	close(start)
	wg.Wait()
	fmt.Printf("buffered channel read value %v concurrently %v per routine cost %v\n", 100, 10000, time.Since(s))
}
