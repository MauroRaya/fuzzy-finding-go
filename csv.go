package main

// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go

// I know encoding/csv exists.
// It panics when reading unformatted .csv files.
// I'll try again later.

import (
	"bufio"
	"io"
	"strings"
)

type Scanner struct {
	Reader *bufio.Reader
	Header map[string]int
	Values []string
	Delim  byte
	Sep    string
}

func NewScannerCSV(r io.Reader, delim byte, sep string) Scanner {
	reader := bufio.NewReader(r)

	row, err := reader.ReadString(delim)
	if err != nil {
		panic(err)
	}

	values := strings.Split(row, sep)

	header := map[string]int{}
	for index, value := range values {
		header[value] = index
	}

	return Scanner{
		Reader: reader,
		Header: header,
		Delim:  delim,
		Sep:    sep,
	}
}

func (scanner *Scanner) Scan() bool {
	row, err := scanner.Reader.ReadString(scanner.Delim)
	if err == io.EOF {
		return false
	}

	if err != nil {
		panic(err)
	}

	values := strings.Split(row, scanner.Sep)
	scanner.Values = values

	return true
}

func (scanner *Scanner) Text(name string) string {
	index := scanner.Header[name]
	return scanner.Values[index]
}
