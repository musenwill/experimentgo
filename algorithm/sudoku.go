package algorithm

import (
	"fmt"
	"sync"
	"container/heap"
)

const rows = 9
const cols = 9
const regionSize = 9

type coordinate struct {
	R int
	C int
}

func unique(list []int) []int {
	set := make(map[int]struct{})
	u := make([]int, 0, len(list))

	for _, val := range list {
		if _, ok := set[val]; !ok {
			set[val] = struct {}{}
			u = append(u, val)
		}
	}

	return u
}

func rowCoords(coord Coord) []Coord {
	coords := make([]Coord, COLS)
	for y := 0; y < COLS; y++ {
		coords[y] = Coord{coord.R, y}
	}
	return coords
}

func colCoords(coord Coord) []Coord {
	coords := make([]Coord, ROWS)
	for x := 0; x < ROWS; x++ {
		coords[x] = Coord{x, coord.C}
	}
	return coords
}

func regionCoords(coord Coord) []Coord {
	coords := make([]Coord, REGION_SIZE)
	rowBase := coord.R - coord.R % 3
	colBase := coord.C - coord.C % 3

	count := 0
	for x := rowBase; x < rowBase + 3; x++ {
		for y := colBase; y < colBase + 3; y++ {
			coords[count] = Coord {x, y}
			count++
		}
	}

	return coords
}

type Board [ROWS][COLS]int

func (b *Board) print() {
	fmt.Println("-------------------")
	for _, row := range b {
		fmt.Println(row)
	}
	fmt.Println("-------------------")
}

func (b *Board) copy() Board {
	return *b;
}

func (b *Board) valueSet(coords []Coord) (map[int]struct{}, error) {
	set := make(map[int]struct{})
	for _, c := range coords {
		value := b[c.R][c.C]
		if _, ok := set[value]; ok && 0 != value {
			return nil, fmt.Errorf("Invalid board, (%d,%d) %d conflict", c.R, c.C, value)
		}
		set[value] = struct {}{}
	}
	return set, nil
}

func (b *Board) optionSet(coord Coord) (map[int]struct{}, error) {
	set := make(map[int]struct{})

	rowSet, error := b.valueSet(rowCoords(coord))
	if nil != error {return nil, error}
	colSet, error := b.valueSet(colCoords(coord))
	if nil != error {return nil, error}
	regionSet, error := b.valueSet(regionCoords(coord))
	if nil != error {return nil, error}

	for k, v := range rowSet {
		set[k] = v
	}
	for k, v := range colSet {
		set[k] = v
	}
	for k, v := range regionSet {
		set[k] = v
	}

	return set, nil
}

func (b *Board) options(coord Coord) (map[int]struct{}, error) {
	options := make(map[int]struct{})
	set, error := b.optionSet(coord)
	if nil != error {
		return nil, error
	}

	for x := 1; x <= ROWS; x++ {
		if _, ok := set[x]; !ok {
			options[x] = struct{}{}
		}
	}

	return options, nil
}

func (b *Board) set(coord Coord, value int) {
	(*b)[coord.R][coord.C] = value
}

type BlankGrid struct {
	coord Coord
	options map[int]struct{}
}

func (bg *BlankGrid) Print() {
	fmt.Printf("(%d,%d) %d ", bg.coord.R, bg.coord.C, len(bg.options))
	fmt.Println(bg.options)
}

func (bg *BlankGrid) Remove(value int) {
	delete(bg.options, value)
}

type BlankGridMinHeap []BlankGrid

func (hp BlankGridMinHeap) Len() int {
	return len(hp)
}

func (hp BlankGridMinHeap) Less(i, j int) bool {
	return len(hp[i].options) < len(hp[j].options)
}

func (hp BlankGridMinHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp *BlankGridMinHeap) Peek() BlankGrid {
	return (*hp)[0]
}

func (hp *BlankGridMinHeap) Push(x interface{}) {
	*hp = append(*hp, x.(BlankGrid))
}

func (hp *BlankGridMinHeap) Pop() interface{} {
	old := *hp
	n := len(old)
	item := old[n-1]
	*hp = old[0 : n-1]
	return item
}

func (hp *BlankGridMinHeap) Copy() *BlankGridMinHeap {
	cp := BlankGridMinHeap{}
	for _, val := range *hp {
		options := make(map[int]struct{})
		for k, v := range val.options {
			options[k] = v
		}
		cp = append(cp, BlankGrid{Coord{val.coord.R, val.coord.C}, options})
	}
	return &cp
}

type Sudoku struct {
	board Board
	minHeap BlankGridMinHeap
	wg sync.WaitGroup
}

func (s *Sudoku) init() error {
	s.minHeap = BlankGridMinHeap{}
	for x := 0; x < ROWS; x++ {
		for y := 0; y < COLS; y++ {
			val := s.board[x][y]
			if 0 == val {
				options, error := s.board.options(Coord{x, y})
				if nil != error {
					return error;
				}
				s.minHeap.Push(BlankGrid{Coord{x, y}, options})
			}
		}
	}
	heap.Init(&s.minHeap)
	return nil
}

func (s *Sudoku) printHeap() {
	for _, val := range s.minHeap {
		val.Print()
	}
}

func (s *Sudoku) getRelatedGrids(hp *BlankGridMinHeap, coord Coord) []*BlankGrid {
	related := []*BlankGrid{}

	for index, grid := range *hp {
		if grid.coord.R == coord.R {
			related = append(related, &(*hp)[index])
			continue
		}
		if grid.coord.C == coord.C {
			related = append(related, &(*hp)[index])
			continue
		}

		rowBase := coord.R - coord.R % 3
		colBase := coord.C - coord.C % 3
		R := grid.coord.R
		C := grid.coord.C
		if R >= rowBase && R < rowBase + 3 && C >= colBase && C < colBase + 3 {
			related = append(related, &(*hp)[index])
			continue
		}
	}

	return related
}

func (s *Sudoku) recurse(board Board, hp *BlankGridMinHeap) {
	defer s.wg.Done()

	if len(*hp) <= 0 {
		board.print()
		return
	}

	heap.Init(hp)
	grid := heap.Pop(hp).(BlankGrid)
	s.wg.Add(len(grid.options))

	for k, _ := range grid.options {
		hpcpy := hp.Copy()
		board.set(grid.coord, k)
		related := s.getRelatedGrids(hpcpy, grid.coord)
		for _, r := range related {
			r.Remove(k)
		}
		go s.recurse(board.copy(), hpcpy)
	}
}

func (s *Sudoku) solve() {
	error := s.init()
	if nil != error {
		fmt.Println(error)
		return
	}
	s.wg.Add(1)
	go s.recurse(s.board, &s.minHeap)
	s.wg.Wait()
}

func main() {
	board := Board {
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{4, 5, 6, 7, 8, 9, 1, 2, 3},
			{7, 8, 9, 1, 2, 3, 4, 5, 6},
			{2, 3, 1, 9, 0, 0, 0, 0, 0},
			{5, 6, 4, 0, 0, 0, 0, 0, 0},
			{8, 9, 7, 0, 0, 0, 0, 0, 0},
			{3, 1, 2, 0, 0, 0, 0, 0, 0},
			{6, 4, 5, 0, 0, 0, 0, 0, 0},
			{9, 7, 8, 0, 0, 0, 0, 0, 0},
	}
	sudoku := Sudoku{board, nil, sync.WaitGroup{}}
	sudoku.solve()
}
