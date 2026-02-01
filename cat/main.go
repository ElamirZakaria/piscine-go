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
				os.Stderr.WriteString("Error: " + err.Error() + "\n")
				os.Exit(1)
				// fmt.Println("Error:", err)
				// os.Exit(1)
			} else {
				//	Data,err1 := os.ReadFile(args[1])
				// fmt.Print(string(fileData))
				for _, a := range string(fileData) {
					z01.PrintRune(a)
				}
			}
		}
	}
}
