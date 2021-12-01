package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	validCount := 0
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	scanner := bufio.NewScanner(f)
	scanner.Split(splitByDoubleNewLine)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)

		valid := true
		for _, requiredField := range requiredFields {
			if !strings.Contains(line, requiredField+":") {
				valid = false
				break
			}
		}
		if valid {
			validCount++
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(validCount)
}

func splitByDoubleNewLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0 : i+1]), nil
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
