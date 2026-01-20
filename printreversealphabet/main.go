package main

import "github.com/01-edu/z01"

func main() {
	i := 'z'
	for i != '\t' {

		z01.PrintRune(i)

		if i == 'a' {
			i = '\v'
		}
		i--

	}
}
