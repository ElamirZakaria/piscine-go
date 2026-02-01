package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
	} else {
		for i := 1; i < len(args); i++ {
			fileData, err := os.ReadFile(args[i])
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			} else {
				//	Data,err1 := os.ReadFile(args[1])
				fmt.Print(string(fileData))
			}
		}
	}
}
