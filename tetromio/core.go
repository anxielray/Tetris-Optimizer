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
func ReadInput(file io.Reader) [][4][4]string {
	var tetrominoArray [][4][4]string
	var tetromino [4][4]string

	// read from the test file line by line
	scanner := bufio.NewScanner(file)
	/*
		Initialize; i - , in -, flag -,
	*/
	i, in, flag, alpha := 0, 0, true, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for scanner.Scan() {

		// call each line str
		str := scanner.Text()
		if str == "" {
			if flag {
				flag = false
				continue
			} else {
				fmt.Println("ERROR")
				os.Exit(0)
			}
		}
		var arr [4]string

		// if the line doesn't have exactly 4 characters
		if len(str) != 4 {
			fmt.Println("ERROR")
			os.Exit(0)
		}
		arr = ConvertToAlphabet(str, alpha, arr, in)
		tetromino[i] = arr
		i++
		if i == 4 {
			flag = true
			i = 0
			in++
			if !ValidateTetro(tetromino) {
				fmt.Println("ERROR")
				os.Exit(1)
			}
			tetromino = OptimizeTetromino(tetromino)
			tetrominoArray = append(tetrominoArray, tetromino)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return tetrominoArray
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

//the function takes a 
func ConvertToAlphabet(str, alpha string, arr [4]string, in int) [4]string {
	for ind := range arr {
		if rune(str[ind]) == '.' {
			arr[ind] = "."
		} else if rune(str[ind]) == '#' {
			arr[ind] = string(alpha[in])
		} else {
			fmt.Println("ERROR")
			os.Exit(0)
		}
	}
	return arr
}
