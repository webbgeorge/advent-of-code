package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("2021/01/pt2/input.txt")
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
	// Start at 4th number (2nd measurement window
	for i := 3; i < len(numbers); i++ {
		m1 := numbers[i-3] + numbers[i-2] + numbers[i-1]
		m2 := numbers[i-2] + numbers[i-1] + numbers[i]
		if m2 > m1 {
			increaseCount++
		}
	}

	fmt.Println(increaseCount)
}
