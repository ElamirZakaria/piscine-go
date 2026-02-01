package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

// func PrintInt(a int) {
// 	b := a/10 + '0'
// 	c := a%10 + '0'
// 	z01.PrintRune(rune(b))
// 	z01.PrintRune(rune(c))
// }

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	var s string = "x = 42 , y = 21"
	// var s2 string = ""
	//	var x rune = points.x + '0'
	for _, a := range s {
		// if a == '!' {
		// 	// var b rune = rune(points.x/10 + '0')

		// } else if a == '@' {
		// 	// z01.PrintRune('4')
		// } else {
		z01.PrintRune(a)
		// }
	}

	z01.PrintRune('\n')
}
