package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

		min, err := strconv.Atoi(submatches[0][1])
		if err != nil {
			panic(fmt.Sprintf("invalid min '%s'", line))
		}
		max, err := strconv.Atoi(submatches[0][2])
		if err != nil {
			panic(fmt.Sprintf("invalid max '%s'", line))
		}
		char := submatches[0][3]
		pw := submatches[0][4]

		count := strings.Count(pw, char)
		if count < min || count > max {
			continue
		}

		validCount++
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("valid count: %d\n", validCount)
}
