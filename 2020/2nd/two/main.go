package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	validCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): (.+)`)
		submatches := re.FindAllStringSubmatch(line, -1)
		if len(submatches) != 1 || len(submatches[0]) != 5 {
			panic(fmt.Sprintf("invalid line '%s'", line))
		}

		pos1, err := strconv.Atoi(submatches[0][1])
		if err != nil {
			panic(fmt.Sprintf("invalid pos1 '%s'", line))
		}
		pos2, err := strconv.Atoi(submatches[0][2])
		if err != nil {
			panic(fmt.Sprintf("invalid pos2 '%s'", line))
		}
		char := submatches[0][3]
		pw := submatches[0][4]

		if pw[pos1-1] == char[0] || pw[pos2-1] == char[0] {
			if pw[pos1-1] == char[0] && pw[pos2-1] == char[0] {
				// not allowed
			} else {
				validCount++
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("valid count: %d\n", validCount)
}
