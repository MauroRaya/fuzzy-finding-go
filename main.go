package main

import (
	"fmt"
)

func main() {
	score, _ := Hamming("London", "Londen")
	fmt.Println(score)
}
