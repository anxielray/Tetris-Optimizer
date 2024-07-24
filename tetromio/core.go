package tetromio

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/*
This function takes in an io writer and prints out a 3D slice.
The function scans the file line by line and snatches the tetrominoes
when it finds a valid end of the tetromino. The tetromino is then converted
(the '#' to be specific) are changed by their corresponding latin letters.
After which, the tetromino is validated than optimized, before being used
on the board.
*/
func Brain_Storm(file io.Reader) [][4][4]string {
	var tetrominoBox [][4][4]string
	var tetromino [4][4]string
	scanner := bufio.NewScanner(file)
	var (
		i              = 0
		tetrominoCount = 0
		latin          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		findTetromino  = true
	)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			if findTetromino {
				findTetromino = false
				continue
			} else {
				fmt.Println("ERROR")
				os.Exit(0)
			}
		}
		var converts [4]string
		if len(str) != 4 {
			fmt.Println("ERROR")
			os.Exit(0)
		}
		converts = AlphamConvert(str, latin, converts, tetrominoCount)
		tetromino[i] = converts
		i++
		if i == 4 {
			findTetromino = true
			i = 0
			tetrominoCount++
			if !MatrixMorphologist(tetromino) {
				fmt.Println("ERROR")
				os.Exit(0)
			}
			tetromino = OptimizeTetromino(tetromino)
			tetrominoBox = append(tetrominoBox, tetromino)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return tetrominoBox
}

/*
The function will create the grid/board that will contain
the tetrominoes. It initializes an empty board.
*/
func GridMaker(n int) [][]string {
	var Square [][]string
	var row []string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		Square = append(Square, row)
		row = []string{}
	}
	return Square
}

/*
The function converts a 4-character string `str` using `latin` for substitution:
- '.' is copied directly to `converts`.
- '#' is replaced by `latin[in]`.
Any other character causes an "ERROR" and exits.
*/
func AlphamConvert(str, alpha string, converts [4]string, in int) [4]string {
	for ind := range converts {
		if rune(str[ind]) == '.' {
			converts[ind] = "."
		} else if rune(str[ind]) == '#' {
			converts[ind] = string(alpha[in])
		} else {
			fmt.Println("ERROR")
			os.Exit(0)
		}
	}
	return converts
}
