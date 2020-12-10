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

	gaps := make(map[int]int)
	lastAdapter := 0
	for i := 0; i < len(adapters); i++ {
		gap := adapters[i] - lastAdapter
		lastAdapter = adapters[i]
		if gap > 3 {
			panic("adapters not compatible")
		}
		gaps[gap]++
	}
	gaps[3]++ // inbuilt is always 3 higher

	fmt.Printf("%+v\n", gaps)
	fmt.Printf("%d\n", gaps[1]*gaps[3])
}
