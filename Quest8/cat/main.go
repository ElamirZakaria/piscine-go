package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// Data is being piped in!
			input, _ := io.ReadAll(os.Stdin)
			for _, r := range string(input) {
				z01.PrintRune(r)
			}
		}
	} else {
		for i := 0; i < len(args); i++ {
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
