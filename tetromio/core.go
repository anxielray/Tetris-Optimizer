package tetromio

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/*
this function takes in an io writer and prints out a 3D slice.
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
		//reset the variables back to fatory settings for the next tetromino
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

func PrintInputs(myArray [][4][4]string) {
	// Prints inputted tetrominoes
	for _, tetromino := range myArray {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				fmt.Printf("%s ", tetromino[j][k])
			}
			fmt.Printf("\n")
		}
		fmt.Println()
	}
}

func SquareCreator(n int) [][]string {
	// initializes a square
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
AlphamConvert converts a 4-character string `str` using `latin` for substitution:
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
