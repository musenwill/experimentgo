package sample

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroupTest ...
func WaitGroupTest() {
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < 10; i++ {
			index := i
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Println("runing routine ", index)
			}()
			time.Sleep(300 * time.Millisecond)
		}
	}()
	wg.Wait()
}

/*
动态创建 goroutine 并用 WaitGroup 进行同步是有风险的。
因为 Wait 方法在 WaitGroup 计数为 0 时返回，而在动态创建 goroutine 的过程中，很可能某个
时间点 WaitGroup 为 0，后续的 goroutine 都没来得及开始，Wait 方法就返回了。
*/
