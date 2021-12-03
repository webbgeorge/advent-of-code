package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("2021/03/pt1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var bitCounts []map[rune]int

	lineCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if lineCount == 0 {
			lineCount = len(line)
			bitCounts = make([]map[rune]int, lineCount)
		}

		// probably a way to do this with bitwise operators instead
		for i, bit := range line {
			if bitCounts[i] == nil {
				bitCounts[i] = make(map[rune]int)
			}
			bitCounts[i][bit]++
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	gammaBits := ""
	epsilonBits := ""
	for _, bc := range bitCounts {
		if bc['0'] > bc['1'] {
			gammaBits = gammaBits + "0"
			epsilonBits = epsilonBits + "1"
			continue
		}
		gammaBits = gammaBits + "1"
		epsilonBits = epsilonBits + "0"
	}

	gamma, err := strconv.ParseInt(gammaBits, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(epsilonBits, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("gamma: %d, epsilon: %d, product: %d", gamma, epsilon, gamma*epsilon)
}
