package sample

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
动态创建 goroutine 并用 WaitGroup 进行同步是有风险的。
因为 Wait 方法在 WaitGroup 计数为 0 时返回，而在动态创建 goroutine 的过程中，很可能某个
时间点 WaitGroup 为 0，后续的 goroutine 都没来得及开始，Wait 方法就返回了。
*/
func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	var count int32

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				atomic.AddInt32(&count, 1)
			}()
			time.Sleep(time.Millisecond)
		}
	}()
	wg.Wait()

	if exp, act := int32(0), count; exp != count {
		t.Errorf("got %v expected %v", act, exp)
	}
}
