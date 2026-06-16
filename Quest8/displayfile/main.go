package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("File name missing")
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fileData, err := os.ReadFile(args[1])
		if err != nil {
		} else {
			fmt.Print(string(fileData))
		}

	}
}
