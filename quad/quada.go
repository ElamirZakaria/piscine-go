package piscine

import "fmt"

func QuadA(x, y int) {
	// 0) Check if negative
	if x <= 0 || y <= 0 {
		fmt.Println("")
	} else {

		// 1) create matrix
		matrix := make([][]string, y)
		for i := 0; i < y; i++ {
			matrix[i] = make([]string, x)
		}

		// 2) fill matrix
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if (i == 0 && j == 0) || (i == 0 && j == x-1) || (i == y-1 && j == 0) || (i == y-1 && j == x-1) {
					matrix[i][j] = "o"
				} else if (i == 0 && j != 0) || (i == y-1 && j != x-1) {
					matrix[i][j] = "-"
				} else if (i != 0 && j == 0) || (i != y-1 && j == x-1) {
					matrix[i][j] = "|"
				} else {
					matrix[i][j] = " "
				}
			}
		}
		// 3) print matrix (NO [])
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				fmt.Print(matrix[i][j])
			}
			fmt.Println()
		}
	}
}
