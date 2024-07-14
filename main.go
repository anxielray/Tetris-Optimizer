package main

import (
	"fmt"
	"os"

	tetro "tetraGame/tetromio"
)

func main() {
	args := os.Args[1]
	if args != "" {
		file, err := os.Open(args)
		if err != nil {
			fmt.Println("ERROR")
			os.Exit(1)
		}
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("ERROR")
				os.Exit(1)
			}
		}()
		var myArray = tetro.ReadInput(file)
		tetro.Solve(myArray)
		tetro.PrintSol()
	}
}
