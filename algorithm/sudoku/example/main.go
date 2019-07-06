package main

import "github.com/musenwill/experimentgo/algorithm/sudoku"

func main() {
	board := sudoku.Board{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 1, 5, 0, 0, 0, 0, 0},
		{5, 6, 4, 0, 3, 0, 0, 0, 0},
		{8, 9, 7, 0, 0, 1, 0, 0, 0},
		{3, 1, 2, 0, 0, 0, 5, 0, 0},
		{6, 4, 5, 0, 0, 0, 0, 3, 0},
		{9, 7, 8, 0, 0, 0, 0, 0, 1},
	}
	sudoku := sudoku.New(board)
	sudoku.Solve()
}
