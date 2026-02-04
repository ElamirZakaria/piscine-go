package main

import (
	"fmt"
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

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 3 {
		firstNumber, err1 := stoi(arguments[0])

		secondNumber, err2 := stoi(arguments[2])
		if err1 == false || err2 == false {
		} else {
			if arguments[1] == "+" {
				fmt.Println(firstNumber + secondNumber)
			} else if arguments[1] == "-" {
				fmt.Println(firstNumber - secondNumber)
			} else if arguments[1] == "*" {
				fmt.Println(firstNumber * secondNumber)
			} else if arguments[1] == "/" {
				if secondNumber == 0 {
					fmt.Println("No division by 0")
				} else {
					fmt.Println(firstNumber / secondNumber)
				}
			} else if arguments[1] == "%" {
				if secondNumber == 0 {
					fmt.Println("No modulo by 0")
				} else {
					fmt.Println(firstNumber % secondNumber)
				}
			}
		}
	}
}
