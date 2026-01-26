package piscine

import "fmt"

func QuadB(x, y int) {
	width := x
	height := y
	if x <= 0 || y <= 0 {
		fmt.Println("")
	} else {
		for i := 1; i <= height; i++ {
			for j := 1; j <= width; j++ {
				if (i == 1 && j == 1) || (i == height && j == width && i != 1 && j != 1) {
					fmt.Print("/")
				} else if (i == 1 && j == width) || (i == height && j == 1) {
					fmt.Print("\\")
				} else if (i != 1 && j == width) || (i == height && j != 1) || (i == 1 && j != width) || (i != height && j == 1) {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}
