package piscine

import "fmt"

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		fmt.Println("")
	} else {
		for rows := 1; rows <= y; rows++ {
			for columns := 1; columns <= x; columns++ {
				if (rows == 1) || (rows == y) || (columns == 1) || (columns == x) {
					if (rows == 1 && columns == 1) || (rows == y && columns == 1) {
						fmt.Print("A")
					} else if (rows == y && columns == x) || (rows == 1 && columns == x) {
						fmt.Print("C")
					} else {
						fmt.Print("B")
					}
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
