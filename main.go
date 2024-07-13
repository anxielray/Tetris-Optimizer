package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR!")
		os.Exit(0)
	}
}