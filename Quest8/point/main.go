package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x rune
	y rune
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
	var s string = "x = !, y = @"

	for _, a := range s {
		if a == '!' {
			// These 2 calls print the digits of x
			num := points.x/10 + '0'
			for i := 0; i < 2; i++ {
				z01.PrintRune(num)
				num = points.x%10 + '0'
			}

		} else if a == '@' {
			// These 2 calls print the digits of y
			num := points.y/10 + '0'
			for i := 0; i < 2; i++ {
				z01.PrintRune(num)
				num = points.y%10 + '0'
			}
		} else {
			z01.PrintRune(a)
		}
	}
	z01.PrintRune('\n')
}
