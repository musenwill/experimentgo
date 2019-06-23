package api

import (
	"sync"
)

type WorkerPool interface {
	// add a task to pool to excute
	Excute(func())

	// wait for all tasks to end and do something after that
	Wait(func())

	// stop all running tasks, this cause Wait() to return immediately
	Cancel()
}

var _ WorkerPool = &workerPool{}

type workerPool struct {
	nWait   sync.WaitGroup
	nWorker chan bool
}

func NewWorkerPool(size int) WorkerPool {
	pool := &workerPool{
		nWorker: make(chan bool, size),
	}
	// add 1, prevent worker pool finish too early
	pool.nWait.Add(1)
	return pool
}

/* if excute all workers in a goroutine, be sure call Wait() after the goroutine has finished, eg:
pool := NewWorkerPool(10)
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	for i:=0; i<100; i++ {
		pool.Excute(func(){
			// do something
		})
	}
}
wg.Wait()	// this make sure go func() has finished, and all workers are be putted into pool
pool.Wait(nil)
*/
func (this *workerPool) Excute(c func()) {
	this.nWorker <- true // acquire token
	this.nWait.Add(1)
	go func() {
		defer this.nWait.Done()
		defer func() { <-this.nWorker }() // release token
		c()
	}()
}

func (this *workerPool) Wait(c func()) {
	this.nWait.Done() // correspond to nWait.Add(1) in NewWorkerPool()
	this.nWait.Wait()
	if nil != c {
		c()
	}
}

/* unimplemented yet */
func (this *workerPool) Cancel() {
}
