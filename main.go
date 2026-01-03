package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Match struct {
	Key   string
	Score float32
}

func main() {
	file, err := os.Open("data/BRAZIL_CITIES.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	col := strings.ToUpper(os.Args[1])
	input := os.Args[2]

	scanner := NewScannerCSV(file, '\n', ";")

	matches := []Match{}
	for scanner.Scan() {
		text := scanner.Text(col)
		score := Levenshtein(input, text)
		matches = append(matches, Match{text, score})
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Score < matches[j].Score
	})

	for _, match := range matches[:20] {
		fmt.Printf("%s %f\n", match.Key, match.Score)
	}
}
