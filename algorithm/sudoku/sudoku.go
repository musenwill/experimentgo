package sudoku

import (
	"container/heap"
	"fmt"
)

type Sudoku struct {
	Board   Board
	minHeap *minHeap
}

func (s *Sudoku) init() error {
	s.minHeap = &minHeap{}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			crd := coordinate{x, y}
			val := s.Board[x][y]
			if 0 == val {
				options, error := s.Board.options(crd)
				if nil != error {
					return error
				}
				s.minHeap.Push(blankGrid{crd, options})
			}
		}
	}
	heap.Init(s.minHeap)
	return nil
}

func (s *Sudoku) printHeap() {
	for _, val := range *s.minHeap {
		val.print()
	}
}

func (s *Sudoku) getRelatedGrids(hp *minHeap, crd coordinate) []*blankGrid {
	related := []*blankGrid{}

	for index, grid := range *hp {
		if grid.crd.row == crd.row {
			related = append(related, &(*hp)[index])
			continue
		}
		if grid.crd.col == crd.col {
			related = append(related, &(*hp)[index])
			continue
		}

		rowBase := crd.row - crd.row%3
		colBase := crd.col - crd.col%3
		row := grid.crd.row
		col := grid.crd.col
		if row >= rowBase && row < rowBase+3 && col >= colBase && col < colBase+3 {
			related = append(related, &(*hp)[index])
			continue
		}
	}

	return related
}

func (s *Sudoku) recurse(board Board, hp *minHeap, results chan<- Board) {
	if len(*hp) <= 0 {
		results <- board
		return
	}

	heap.Init(hp)
	grid := heap.Pop(hp).(blankGrid)

	for k := range grid.options {
		cp := hp.Copy()
		board.set(grid.crd, k)
		related := s.getRelatedGrids(cp, grid.crd)
		for _, r := range related {
			r.remove(k)
		}
		s.recurse(board.copy(), cp, results)
	}
}

func (s *Sudoku) Solve(results chan<- Board) {
	error := s.init()
	if nil != error {
		fmt.Println(error)
		return
	}

	s.recurse(s.Board, s.minHeap, results)
	close(results)
}

func New(board Board) *Sudoku {
	return &Sudoku{board, nil}
}
