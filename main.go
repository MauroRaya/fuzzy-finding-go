package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getColumnIndex(scanner *bufio.Scanner, delim string, name string) int {
	scanner.Scan()
	row := scanner.Text()

	cols := strings.Split(row, delim)

	for idx, col := range cols {
		if col == name {
			return idx
		}
	}

	return -1
}

func getSliceColumn(scanner *bufio.Scanner, index int) []string {
	slice := make([]string, 0)

	for scanner.Scan() {
		row := scanner.Text()
		values := strings.Split(row, ";")
		slice = append(slice, values[index])
	}

	return slice
}

type FuzzyScore struct {
	Key   string
	Value int
}

func main() {
	file, err := os.Open("data/BRAZIL_CITIES.csv")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	idx := getColumnIndex(scanner, ";", "CITY")
	cities := getSliceColumn(scanner, idx)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	scores := make([]FuzzyScore, 0)

	for _, city := range cities {
		score := Hamming(input, city)
		scores = append(scores, FuzzyScore{city, score})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Value < scores[j].Value
	})

	limit := 20

	for i, j := range scores {
		fmt.Printf("%s %d\n", j.Key, j.Value)

		if i == limit {
			break
		}
	}
}
