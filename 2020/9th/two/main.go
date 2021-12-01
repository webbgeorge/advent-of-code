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

	invalidI := 0
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
			invalidI = i
			fmt.Printf("part one: (%d) %d\n", i, numbers[i])
			break
		}
	}

	for i := 0; i < len(numbers)-1; i++ {
		sum := numbers[i]
		lowest := numbers[i]
		highest := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if numbers[j] < lowest {
				lowest = numbers[j]
			}
			if numbers[j] > highest {
				highest = numbers[j]
			}
			if sum == numbers[invalidI] {
				fmt.Println(highest + lowest)
				return
			}
			if sum > numbers[invalidI] {
				break
			}
		}
	}
}
