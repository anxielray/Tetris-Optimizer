package main

import (
	"fmt"
	"os"
)

func main() {
	//check cmd arguments
	if len(os.Args) != 2 {
		fmt.Println("ERROR!")
		os.Exit(0)
	}
	//read test file
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}
	defer file.Close()
	/*validation of tetrominoes
	* 1. check if all tetrominoes are unique
	* 3. check if all tetrominoes are correct
	  - It should have exactly 4 rows.
	  - Each row should have exactly 4 characters ('#' or .).
	  - There should be exactly 4 '#' characters in total.
	  - Ensure that each tetromino has exactly 6 sides of '#' cells touching.
	*/

	//board calculation

	//backtracking algorithm

}
