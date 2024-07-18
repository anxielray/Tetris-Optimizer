package tetromio

import (
	"fmt"
	"os"
)

func FileStats(filename string) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}
	if fileInfo.Size() == 0 {
		fmt.Println("ERROR")
		os.Exit(0)
	}
}
