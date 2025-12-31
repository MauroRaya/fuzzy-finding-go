package main

import (
	"bufio"
	"fmt"
	"os"
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

	idx := getColumnIndex(scanner, ";", "STATE")
	slice := getSliceColumn(scanner, idx)

	fmt.Println(slice)
}
