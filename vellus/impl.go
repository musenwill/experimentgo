package vellus

import (
	"sync"
)

/*Vellus is an interface defines the abilities of vellus */
type Vellus interface {
	// add a job to excute
	Excute(func())

	// wait for all goroutines to end and do some clean up after that
	Wait(func())
}

var _ Vellus = &pool{}

type pool struct {
	nWait   sync.WaitGroup
	nWorker chan bool
}

/*NewVellus is a Vellus creator that can limit the max amount of goroutines at any time */
func NewVellus(size int) Vellus {
	pool := &pool{
		nWorker: make(chan bool, size),
	}
	// add 1, prevent worker pool finish too early
	pool.nWait.Add(1)
	return pool
}

/*Excute let you put a job into vellus.
If excute all workers in a goroutine, be sure call Wait() after the goroutine has finished, eg:
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
func (p *pool) Excute(c func()) {
	p.nWorker <- true // acquire token
	p.nWait.Add(1)
	go func() {
		defer p.nWait.Done()
		defer func() { <-p.nWorker }() // release token
		c()
	}()
}

/*Wait for all goroutines to end */
func (p *pool) Wait(c func()) {
	p.nWait.Done() // correspond to nWait.Add(1) in NewVellus()
	p.nWait.Wait()
	if nil != c {
		c()
	}
}
