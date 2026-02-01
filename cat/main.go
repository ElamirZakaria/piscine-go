package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) == 1 {
	} else {
		for i := 1; i < len(args); i++ {
			fileData, err := os.ReadFile(args[i])
			if err != nil {
				os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			} else {
				for _, a := range string(fileData) {
					z01.PrintRune(a)
				}
			}
		}
	}
}
