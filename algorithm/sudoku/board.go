package sudoku

import "fmt"

const (
	rows       = 9
	cols       = 9
	regionSize = 9
)

type coordinate struct {
	row int
	col int
}

func rowCoords(crd coordinate) []coordinate {
	coords := make([]coordinate, cols)
	for y := 0; y < cols; y++ {
		coords[y] = coordinate{crd.row, y}
	}
	return coords
}

func colCoords(crd coordinate) []coordinate {
	coords := make([]coordinate, rows)
	for x := 0; x < rows; x++ {
		coords[x] = coordinate{x, crd.col}
	}
	return coords
}

func regionCoords(crd coordinate) []coordinate {
	coords := make([]coordinate, regionSize)
	rowBase := crd.row - crd.row%3
	colBase := crd.col - crd.col%3

	count := 0
	for x := rowBase; x < rowBase+3; x++ {
		for y := colBase; y < colBase+3; y++ {
			coords[count] = coordinate{x, y}
			count++
		}
	}

	return coords
}

type Board [rows][cols]int

func (b *Board) Print() {
	fmt.Println("-------------------")
	for _, row := range b {
		fmt.Println(row)
	}
	fmt.Println("-------------------")
}

func (b *Board) copy() Board {
	return *b
}

func (b *Board) valueSet(coords []coordinate) (map[int]struct{}, error) {
	set := make(map[int]struct{})
	for _, c := range coords {
		value := b[c.row][c.col]
		if _, ok := set[value]; ok && 0 != value {
			return nil, fmt.Errorf("Invalid board, (%d,%d) %d conflict", c.row, c.col, value)
		}
		set[value] = struct{}{}
	}
	return set, nil
}

func (b *Board) optionSet(crd coordinate) (map[int]struct{}, error) {
	set := make(map[int]struct{})

	rowSet, error := b.valueSet(rowCoords(crd))
	if nil != error {
		return nil, error
	}
	colSet, error := b.valueSet(colCoords(crd))
	if nil != error {
		return nil, error
	}
	regionSet, error := b.valueSet(regionCoords(crd))
	if nil != error {
		return nil, error
	}

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

func (b *Board) options(crd coordinate) (map[int]struct{}, error) {
	options := make(map[int]struct{})
	set, error := b.optionSet(crd)
	if nil != error {
		return nil, error
	}

	for x := 1; x <= rows; x++ {
		if _, ok := set[x]; !ok {
			options[x] = struct{}{}
		}
	}

	return options, nil
}

func (b *Board) set(crd coordinate, value int) {
	(*b)[crd.row][crd.col] = value
}
