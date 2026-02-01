package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintInt(a int) {
	b := a/10 + '0'
	c := a%10 + '0'
	z01.PrintRune(rune(b))
	z01.PrintRune(rune(c))
}

func main() {
	points := point{}

	setPoint(&points)
	var s string = "x = ! , y = @"
	// var s2 string = ""
	for _, a := range s {
		if a == '!' {
			PrintInt(points.x)
		} else if a == '@' {
			PrintInt(points.y)
		} else {
			z01.PrintRune(a)
		}
	}

	z01.PrintRune('\n')
}
