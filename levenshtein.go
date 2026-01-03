package main

import "slices"

func Levenshtein(a, b string) float32 {
	cache := make([][]float32, len(a)+1)
	for i := range cache {
		cache[i] = make([]float32, len(b)+1)
	}

	for i := 0; i <= len(a); i++ {
		cache[i][0] = float32(i)
	}
	for j := 0; j <= len(b); j++ {
		cache[0][j] = float32(j)
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cache[i][j] = cache[i-1][j-1]
			} else {
				insert := cache[i][j-1]
				delete := cache[i-1][j]
				replace := cache[i-1][j-1]

				cache[i][j] = 1 + slices.Min([]float32{insert, delete, replace})
			}
		}
	}

	return cache[len(a)][len(b)]
}
