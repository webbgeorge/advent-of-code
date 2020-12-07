package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cumulativeCount := 0

	scanner := bufio.NewScanner(f)
	scanner.Split(splitByDoubleNewLine)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)

		questions := make(map[rune]int)
		for _, c := range line {
			if unicode.IsSpace(c) {
				continue
			}
			questions[c]++
		}
		cumulativeCount += len(questions)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(cumulativeCount)
}

func splitByDoubleNewLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n','\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i+1]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
