package main

import (
	"os"

	"github.com/01-edu/z01"
)

func NumToRune(numStr string) rune {
	switch numStr {
	case "1":
		return 'a'
	case "2":
		return 'b'
	case "3":
		return 'c'
	case "4":
		return 'd'
	case "5":
		return 'e'
	case "6":
		return 'f'
	case "7":
		return 'g'
	case "8":
		return 'h'
	case "9":
		return 'i'
	case "10":
		return 'j'
	case "11":
		return 'k'
	case "12":
		return 'l'
	case "13":
		return 'm'
	case "14":
		return 'n'
	case "15":
		return 'o'
	case "16":
		return 'p'
	case "17":
		return 'q'
	case "18":
		return 'r'
	case "19":
		return 's'
	case "20":
		return 't'
	case "21":
		return 'u'
	case "22":
		return 'v'
	case "23":
		return 'w'
	case "24":
		return 'x'
	case "25":
		return 'y'
	case "26":
		return 'z'
	default:
		return '0' // invalid input
	}
}

func main() {
	arguments := os.Args
	temp := 0
	if arguments[1] == "--upper" {
		temp = 32
	}
	for _, a := range arguments {
		tobeprinted := NumToRune(a)
		if tobeprinted != '0' {
			z01.PrintRune(tobeprinted - rune(temp))
		}
	}
	z01.PrintRune('\n')
}
