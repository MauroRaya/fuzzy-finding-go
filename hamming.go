package main

import (
	"math"
)

func Hamming(a, b string) float32 {
	if len(a) != len(b) {
		return math.MaxFloat32
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return float32(distance)
}
