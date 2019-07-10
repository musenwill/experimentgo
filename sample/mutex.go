package sample

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var msg string
var once sync.Once
var myOnce = &MyOnce{}

type MyOnce struct {
	initFlag bool
	sync.RWMutex
}

func (o *MyOnce) Do(setup func()) {
	o.RLock()
	initted := o.initFlag
	if initted {
		o.RUnlock()
		return
	}
	o.RUnlock()

	o.Lock()
	defer o.Unlock()
	if !o.initFlag {
		setup()
	}
	o.initFlag = true
}

func setUp() {
	msg = uuid.New().String()
}

func doPrint() {
	fmt.Println(msg)
}

func doPrintSetupOnce() {
	once.Do(setUp)
	fmt.Println(msg)
}

func doPrintSetupMyOnce() {
	myOnce.Do(setUp)
	fmt.Println(msg)
}

func Once() {
	var wg sync.WaitGroup
	count := 100
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			doPrintSetupMyOnce()
		}()
	}
	wg.Wait()
}
