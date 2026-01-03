package main

import (
	"math"
)

func Levenshtein(a, b string) float32 {
	cache := make([][]float32, len(a)+1)
	for i := range cache {
		cache[i] = make([]float32, len(b)+1)
	}

	for i := range len(a) + 1 {
		for j := range len(b) + 1 {
			cache[i][j] = math.MaxFloat32
		}
	}

	for j := range len(b) + 1 {
		cache[len(a)][j] = float32(len(b) - j)
	}
	for i := range len(a) + 1 {
		cache[i][len(b)] = float32(len(a) - i)
	}

	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				cache[i][j] = cache[i+1][j+1]
			} else {
				values := []float32{
					cache[i+1][j],
					cache[i][j+1],
					cache[i+1][j+1],
				}

				min := values[0]
				for _, value := range values {
					if value < min {
						min = value
					}
				}

				cache[i][j] = 1 + min
			}
		}
	}

	return cache[0][0]
}
