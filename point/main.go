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
	var s string = "x = ! , y = 21"
	str := ""
	num := points.x / 10
	switch num {
	case 1:
		str += "1"
	case 2:
		str += "2"
	case 3:
		str += "3"
	case 4:
		str += "4"
	case 5:
		str += "5"
	case 6:
		str += "6"
	case 7:
		str += "7"
	case 8:
		str += "8"
	case 9:
		str += "9"
	case 0:
		str += "0"

	}
	num = points.x % 10
	switch num {
	case 1:
		str += "1"
	case 2:
		str += "2"
	case 3:
		str += "3"
	case 4:
		str += "4"
	case 5:
		str += "5"
	case 6:
		str += "6"
	case 7:
		str += "7"
	case 8:
		str += "8"
	case 9:
		str += "9"
	case 0:
		str += "0"

	}
	str2 := ""
	num = points.y / 10
	switch num {
	case 1:
		str2 += "1"
	case 2:
		str2 += "2"
	case 3:
		str2 += "3"
	case 4:
		str2 += "4"
	case 5:
		str2 += "5"
	case 6:
		str2 += "6"
	case 7:
		str2 += "7"
	case 8:
		str2 += "8"
	case 9:
		str2 += "9"
	case 0:
		str2 += "0"

	}
	num = points.y % 10
	switch num {
	case 1:
		str2 += "1"
	case 2:
		str2 += "2"
	case 3:
		str2 += "3"
	case 4:
		str2 += "4"
	case 5:
		str2 += "5"
	case 6:
		str2 += "6"
	case 7:
		str2 += "7"
	case 8:
		str2 += "8"
	case 9:
		str2 += "9"
	case 0:
		str2 += "0"

	}
	for _, a := range s {
		if a == '!' {
			for _, b := range str {
				z01.PrintRune(b)
			}
		} else if a == '@' {
			for _, b := range str2 {
				z01.PrintRune(b)
			}
		} else {
			z01.PrintRune(a)
		}
	}

	z01.PrintRune('\n')
}
