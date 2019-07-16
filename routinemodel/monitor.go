package routinemodel

type monitor struct {
	valueGetter chan int
	valueSetter chan int
}

func newMonitor() *monitor {
	m := &monitor{make(chan int), make(chan int)}
	go m.teller()
	return m
}

func (m *monitor) get() int {
	return <-m.valueGetter
}

func (m *monitor) set(val int) {
	m.valueSetter <- val
}

func (m *monitor) teller() {
	var value int
	for {
		select {
		case val := <-m.valueSetter:
			value = val
		case m.valueGetter <- value:
		}
	}
}
