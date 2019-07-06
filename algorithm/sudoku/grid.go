package sudoku

import (
	"container/heap"
	"fmt"
)

type blankGrid struct {
	crd     coordinate
	options map[int]struct{}
}

func (bg *blankGrid) print() {
	fmt.Printf("(%d,%d) %d ", bg.crd.row, bg.crd.col, len(bg.options))
	fmt.Println(bg.options)
}

func (bg *blankGrid) remove(value int) {
	delete(bg.options, value)
}

var _ heap.Interface = &minHeap{}

type minHeap []blankGrid

func (hp minHeap) Len() int {
	return len(hp)
}

func (hp minHeap) Less(i, j int) bool {
	return len(hp[i].options) < len(hp[j].options)
}

func (hp minHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp *minHeap) Peek() blankGrid {
	return (*hp)[0]
}

func (hp *minHeap) Push(x interface{}) {
	*hp = append(*hp, x.(blankGrid))
}

func (hp *minHeap) Pop() interface{} {
	old := *hp
	n := len(old)
	item := old[n-1]
	*hp = old[0 : n-1]
	return item
}

func (hp *minHeap) Copy() *minHeap {
	cp := minHeap{}
	for _, val := range *hp {
		options := make(map[int]struct{})
		for k, v := range val.options {
			options[k] = v
		}
		cp = append(cp, blankGrid{val.crd, options})
	}
	return &cp
}
