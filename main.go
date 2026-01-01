package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Match struct {
	Key   string
	Score int
}

// Not complete but thats fine
type Row struct {
	City    string
	State   string
	Capital string
}

func main() {
	SEP := ";"

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

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Type the name of a Brazilian city: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.TrimSpace(input)

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip headers

	matches := []Match{}
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), SEP)
		row := Row{
			City: values[0],
		}

		score := Hamming(input, row.City)
		matches = append(matches, Match{row.City, score})
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Score < matches[j].Score
	})

	for _, match := range matches[:10] {
		fmt.Printf("%s %d\n", match.Key, match.Score)
	}
}
