package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for i, a := range arguments {
		if i == 0 {
			continue
		}
		for _, b := range a {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')
	}
}
