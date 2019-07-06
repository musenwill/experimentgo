package main

import (
	"sync"

	"github.com/musenwill/experimentgo/algorithm/sudoku"
)

func main() {
	board := sudoku.Board{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 1, 5, 0, 0, 0, 0, 0},
		{5, 6, 4, 0, 0, 0, 0, 0, 0},
		{8, 9, 7, 0, 0, 1, 0, 0, 0},
		{3, 1, 2, 0, 0, 0, 5, 0, 0},
		{6, 4, 5, 0, 0, 0, 0, 3, 0},
		{9, 7, 8, 0, 0, 0, 0, 0, 1},
	}
	results := make(chan sudoku.Board, 100)
	var wg sync.WaitGroup
	wg.Add((1))
	go func() {
		defer wg.Done()
		for r := range results {
			r.Print()
		}
	}()
	sudoku := sudoku.New(board)
	sudoku.Solve(results)
	wg.Wait()
}

/*
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 9, 1, 2, 3, 4, 5, 6},
	{2, 3, 1, 5, 0, 0, 0, 0, 0},
	{5, 6, 4, 0, 3, 0, 0, 0, 0},
	{8, 9, 7, 0, 0, 1, 0, 0, 0},
	{3, 1, 2, 0, 0, 0, 5, 0, 0},
	{6, 4, 5, 0, 0, 0, 0, 3, 0},
	{9, 7, 8, 0, 0, 0, 0, 0, 1},
*/
