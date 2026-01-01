package main

import (
	"math"
)

func Hamming(a, b string) (value int) {
	if len(a) != len(b) {
		return math.MaxInt
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance
}
