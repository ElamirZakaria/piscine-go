package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	flag := 0

	for i := 0; i < len(arguments); i++ {
		for j := 0; j < len(arguments)-i-1; j++ {
			if arguments[j] > arguments[j+1] {
				temp := arguments[j]
				arguments[j] = arguments[j+1]
				arguments[j+1] = temp
			}
		}
	}
	for _, a := range arguments {
		// if i == 0 {
		// 	continue
		// }
		for _, b := range a {
			if b == '/' {
				flag = 1
				break
			}
			z01.PrintRune(b)
		}
		if flag == 1 {
			flag = 0
			continue

		}
		z01.PrintRune('\n')
	}
}
