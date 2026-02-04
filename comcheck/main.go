package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	for _, a := range arguments {
		if (a == "01") || (a == "galaxy") || (a == "galaxy 01") {
			fmt.Println("Alert!!!")
			break
		}
	}
}
