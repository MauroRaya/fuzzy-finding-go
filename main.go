package main

import (
	"fmt"
)

func main() {
	score, err := Hamming("London", "Londen")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Score:", score)
	}

	score, err = Hamming("London", "Lndon")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Score:", score)
	}
}
