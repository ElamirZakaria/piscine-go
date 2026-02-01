package main

import "fmt"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}

	setPoint(&points)
	// var s string = "x = ! , y = @"
	// var s2 string = ""
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
