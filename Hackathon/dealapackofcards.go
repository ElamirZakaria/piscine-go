package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	length := 0

	for range deck {
		length++
	}
	numberOfCardsofeachPlayer := length / 4
	fmt.Printf("Player 1:")
	for i := 0; i < numberOfCardsofeachPlayer; i++ {
		if i == numberOfCardsofeachPlayer-1 {
			fmt.Printf(" %d", deck[i])
		} else {
			fmt.Printf(" %d,", deck[i])
		}
	}
	fmt.Printf("\nPlayer 2:")
	for i := numberOfCardsofeachPlayer; i < numberOfCardsofeachPlayer*2; i++ {
		if i == numberOfCardsofeachPlayer*2-1 {
			fmt.Printf(" %d", deck[i])
		} else {
			fmt.Printf(" %d,", deck[i])
		}
	}
	fmt.Printf("\nPlayer 3:")
	for i := numberOfCardsofeachPlayer * 2; i < numberOfCardsofeachPlayer*3; i++ {
		if i == numberOfCardsofeachPlayer*3-1 {
			fmt.Printf(" %d", deck[i])
		} else {
			fmt.Printf(" %d,", deck[i])
		}
	}
	fmt.Printf("\nPlayer 4:")
	for i := numberOfCardsofeachPlayer * 3; i < numberOfCardsofeachPlayer*4; i++ {
		if i == numberOfCardsofeachPlayer*4-1 {
			fmt.Printf(" %d", deck[i])
		} else {
			fmt.Printf(" %d,", deck[i])
		}
	}
	fmt.Printf("\n")
}
