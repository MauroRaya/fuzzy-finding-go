package main

import (
	"fmt"
	"os"
	"sort"
)

type Match struct {
	Key   string
	Score int
}

func main() {
	file, err := os.Open("data/BRAZIL_CITIES.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := NewScannerCSV(file, '\n', ";")

	matches := []Match{}
	for scanner.Scan() {
		text := scanner.Text("CITY")
		score := Hamming("São Paulo", text)
		matches = append(matches, Match{text, score})
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Score < matches[j].Score
	})

	for _, match := range matches[:20] {
		fmt.Printf("%s %d\n", match.Key, match.Score)
	}
}
