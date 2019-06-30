package sample

import (
	"fmt"
)

func monitor() {
	seal := newSeal()
	fmt.Println(seal.getSignCount())
	seal.addSignCount(10)
	fmt.Println(seal.getSignCount())
	fmt.Println(seal.getSignCount())
}

type seal struct {
	signCount      chan int
	signCountDelta chan int
}

func newSeal() *seal {
	seal := &seal{make(chan int), make(chan int)}
	go seal.teller()
	return seal
}

func (s *seal) getSignCount() int {
	return <-s.signCount
}

func (s *seal) addSignCount(delta int) {
	s.signCountDelta <- delta
}

func (s *seal) teller() {
	var signCount int
	for {
		select {
		case delta := <-s.signCountDelta:
			signCount += delta
		case s.signCount <- signCount:
		}
	}
}
