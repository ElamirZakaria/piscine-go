package main

import (
	"os"
)

func stoi(s string) (int, bool) {
	sign := 1
	a := 0
	err := true
	for _, b := range s {
		if b == '-' {
			sign = -1
		} else if b >= '0' && b <= '9' {
			a = a*10 + int(b-'0')
		} else {

			err = false

			break
		}
	}
	if a > 9223372036854775806 || a <= -9223372036854775807 {
		err = false
	}

	return a * sign, err
}

func itos(a int) string {
	str := ""
	if a < 0 {
		str += string('-')
		a = a * -1
	}
	if a == 0 {
		return "0"
	} else {
		arrofintegers := []int{}

		for a != 0 {
			arrofintegers = append(arrofintegers, a%10)
			a = a / 10
		}

		for i := len(arrofintegers) - 1; i >= 0; i-- {
			str += string(rune(arrofintegers[i] + '0'))
		}
	}

	return str
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 3 {
		firstNumber, err1 := stoi(arguments[0])

		secondNumber, err2 := stoi(arguments[2])
		if err1 == false || err2 == false {
		} else {
			if arguments[1] == "+" {
				os.Stdout.WriteString(itos(firstNumber+secondNumber) + "\n")
			} else if arguments[1] == "-" {
				os.Stdout.WriteString(itos(firstNumber-secondNumber) + "\n")
			} else if arguments[1] == "*" {
				os.Stdout.WriteString(itos(firstNumber*secondNumber) + "\n")
			} else if arguments[1] == "/" {
				if secondNumber == 0 {
					os.Stdout.WriteString("No division by 0" + "\n")
				} else {
					os.Stdout.WriteString(itos(firstNumber/secondNumber) + "\n")
				}
			} else if arguments[1] == "%" {
				if secondNumber == 0 {
					os.Stdout.WriteString("No modulo by 0" + "\n")
				} else {
					os.Stdout.WriteString(itos(firstNumber%secondNumber) + "\n")
				}
			}
		}
	}
}
