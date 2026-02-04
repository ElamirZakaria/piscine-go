package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	numberOfCardsofeachPlayer := len(deck) / 4

	fmt.Print("Player 1: ")
	for i := 0; i < numberOfCardsofeachPlayer; i++ {
		if i == numberOfCardsofeachPlayer-1 {
			fmt.Print(deck[i])
		} else {
			fmt.Print(deck[i], ",")
		}
	}
	fmt.Println()
	fmt.Print("Player 2: ")
	for i := numberOfCardsofeachPlayer; i < numberOfCardsofeachPlayer*2; i++ {
		if i == numberOfCardsofeachPlayer*2-1 {
			fmt.Print(deck[i])
		} else {
			fmt.Print(deck[i], ",")
		}
	}
	fmt.Println()
	fmt.Print("Player 3: ")
	for i := numberOfCardsofeachPlayer * 2; i < numberOfCardsofeachPlayer*3; i++ {
		if i == numberOfCardsofeachPlayer*3-1 {
			fmt.Print(deck[i])
		} else {
			fmt.Print(deck[i], ",")
		}
	}
	fmt.Println()
	fmt.Print("Player 4: ")
	for i := numberOfCardsofeachPlayer * 3; i < numberOfCardsofeachPlayer*4; i++ {
		if i == numberOfCardsofeachPlayer*4-1 {
			fmt.Print(deck[i])
		} else {
			fmt.Print(deck[i], ",")
		}
	}
	fmt.Println()
}
