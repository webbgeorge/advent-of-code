package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	adapters := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, n)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	sort.Ints(adapters)

	// create a list of gaps between adapters
	gaps := make([]int, 0)
	lastAdapter := 0
	for i := 0; i < len(adapters); i++ {
		gap := adapters[i] - lastAdapter
		lastAdapter = adapters[i]
		if gap > 3 {
			panic("adapters not compatible")
		}
		gaps = append(gaps, gap)
	}
	gaps = append(gaps, 3)

	// calculate the possible adapters in front of each adapter for removing
	possibleAdapters := make([]int, 0)
	for i := 0; i < len(gaps); i++ {
		n := 0
		if i+1 != len(gaps) {
			for j := i; j < i+3; j++ {
				if gaps[j] > 1 {
					break
				}
				n++
			}
		}
		if n == 0 {
			n = 1
		}
		possibleAdapters = append(possibleAdapters, n)
	}

	// calculate the possible arrangements for each group of adapters with more than 1 possible arrangements
	arrangements := 1
	groupLen := 0
	for i := 0; i < len(possibleAdapters); i++ {
		if possibleAdapters[i] == 1 {
			arrangements *= arrangementsForGroupLength(groupLen)
			groupLen = 0
			continue
		}
		groupLen++
	}
	arrangements *= arrangementsForGroupLength(groupLen)

	fmt.Printf("%+v\n", possibleAdapters)
	fmt.Printf("%d\n", arrangements)
}

// yikes
// pre-calculated count of possible arrangements for a group of adapters separated by 1 jolt
func arrangementsForGroupLength(groupLen int) int {
	switch groupLen {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 7
	}
	panic("unhandled case")
}
