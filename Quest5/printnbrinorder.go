package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var TurnItIntoString string = ""
	if n == 0 {
		z01.PrintRune('0')
	}
	length := 0
	for ; n != 0; n = n / 10 {
		TurnItIntoString += string(n%10 + 48)
	}
	arrayofbytes := []byte(TurnItIntoString)
	for range arrayofbytes {
		length++
	}
	// fmt.Println([]byte(TurnItIntoString))
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arrayofbytes[j] > arrayofbytes[j+1] {
				arrayofbytes[j] = arrayofbytes[j+1] ^ arrayofbytes[j]
				arrayofbytes[j+1] = arrayofbytes[j+1] ^ arrayofbytes[j]
				arrayofbytes[j] = arrayofbytes[j+1] ^ arrayofbytes[j]
			}
		}
	}
	for i := 0; i < length; i++ {
		z01.PrintRune(rune(arrayofbytes[i]))
	}
}
