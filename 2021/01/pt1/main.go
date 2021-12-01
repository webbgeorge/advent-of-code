package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("2021/01/pt1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numbers := make([]int, 0)

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
		numbers = append(numbers, n)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	increaseCount := 0
	// Start at 2nd number
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			increaseCount++
		}
	}

	fmt.Println(increaseCount)
}
