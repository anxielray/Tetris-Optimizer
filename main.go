package main

import (
	"fmt"
	"os"
	"time"

	tetro "tetraGame/tetromio"
)

func main() {
	startTime := time.Now()
	args := os.Args[1]
	if args != "" {
		file, err := os.Open(args)
		if err != nil {
			fmt.Println("ERROR")
			os.Exit(0)
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
		end := time.Now()
		elapsed := end.Sub(startTime)
		fmt.Println(elapsed)
	}
}
