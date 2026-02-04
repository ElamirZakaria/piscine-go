package main

import "github.com/01-edu/z01"

func main() {
	j := 0
	i := '0'
	for j != 2 {
		z01.PrintRune(i)
		i++
		if i > '9' || j == 1 {
			i = '\n'
			j++
		}

	}
}
