package main

import (
	"fmt"
	"os"

	tetra "mathematician/tetromio"
)

func main() {
	// check if there is no argument passed on the cmd...
	if len(os.Args) == 1 {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	// declare the first argument as the test file containing the tetrominoes.
	testFile := os.Args[1]
	file, err := os.Open(testFile)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	// check if the file is empty
	tetra.FileStats(testFile)
	defer file.Close()

	// declaring a 3D slice of string...(the board)
	myArray := tetra.ReadInput(file)

	// the 3D slice is then taken for the Arithmetic solving
	tetra.GeniusSolver(myArray)
	tetra.MagicPrinter()
}
