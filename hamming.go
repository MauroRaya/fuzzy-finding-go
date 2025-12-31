package main

import (
	"fmt"
	"math"
)

func Hamming(a, b string) (value int, err error) {
	if len(a) != len(b) {
		return math.MaxInt, fmt.Errorf("hamming: arguments must have same length")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
