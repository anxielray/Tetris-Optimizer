package tetromio

import (
	"fmt"
	"math"
)

var grid [][]string

/*
The back-tracking algorithm takes the full grid that holds the tetrominones
with the number of that specific tetromino from the file holding them and
checks if the current tetromino can be placed in the grid. If it can, it will
place it and then check if the next tetromino can be placed. If it can, it
will place it and so on. If it can't, it will remove the last tetromino
and try to place it in a different position. If it can't place it in any
position, it will remove it and try to place the next tetromino. If it can't
place the next tetromino, it will remove it and try to place the next one
and so on. If it can't place any tetromino, it will return false.
*/
func BackTrackAlgorithm(tetrominoes [][4][4]string, n int) bool {

	//give the command if all tetrominoes are contained in the board
	if n == len(tetrominoes) {
		return true
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if ValidInsertion(i, j, tetrominoes[n]) {
				Insert(i, j, tetrominoes[n])
				if BackTrackAlgorithm(tetrominoes, n+1) {
					return true
				}
				Remove(i, j, tetrominoes[n])
			}
		}
	}
	return false
}

// the function places a tetromino to a board
func Insert(i, j int, tetro [4][4]string) {
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				c++
				grid[i+a][j+b] = tetro[a][b]
				if c == 4 {
					break
				}
			}
			b++
		}
		b = 0
		a++
	}
}

// the function removes a tetromino
// basically meaning, it replaces them with a "."
func Remove(i, j int, tetro [4][4]string) {
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				if c == 4 {
					break
				}
				grid[i+a][j+b] = "."
			}
			b++
		}
		b = 0
		a++
	}
}

/*
The function creates the mathematical board with the data received
from the GridMaker(). If the initialized board size isn't rational
then the function will add 1 to the original size, and generate a
board with a size +1, meaning 1 extra row and column.
*/
func GeniusSolver(tetrominoes [][4][4]string) [][]string {
	boardCoreIntSize := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominoes)))))
	grid = GridMaker(boardCoreIntSize)
	for !BackTrackAlgorithm(tetrominoes, 0) {
		boardCoreIntSize++
		grid = GridMaker(boardCoreIntSize)
	}
	return grid
}

// the function prints the final piece-work.
func MagicPrinter() {
	for i := range grid {
		for j := range grid {
			fmt.Printf("%s", grid[i][j])
		}
		println()
	}
}
