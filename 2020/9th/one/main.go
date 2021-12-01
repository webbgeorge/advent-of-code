package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const preamble = 25

func main() {
	f, err := os.Open("input.txt")
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

	for i := preamble; i < len(numbers); i++ {
		valid := false
		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if numbers[j]+numbers[k] == numbers[i] {
					valid = true
				}
			}
		}
		if !valid {
			fmt.Printf("invalid number: %d\n", numbers[i])
			return
		}
	}
	fmt.Println("nope")
}
