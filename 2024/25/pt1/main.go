package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("expected args")
		os.Exit(1)
	}

	fp := os.Args[1]

	f, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	keys := make([][]int, 0)
	locks := make([][]int, 0)

	isLock := true
	var curr []int
	currCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		currCount++

		// is first line of a schematic
		if currCount == 1 {
			if line == "#####" {
				isLock = true
			} else {
				isLock = false
			}
			curr = make([]int, 5)
			continue
		}

		// is last line of a schematic
		if currCount == 7 {
			if isLock {
				locks = append(locks, curr)
			} else {
				keys = append(keys, curr)
			}
			currCount = 0
			continue
		}

		for i, char := range line {
			if char == '#' {
				curr[i]++
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	limit := 5
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for k := 0; k < len(lock); k++ {
				if lock[k]+key[k] > limit {
					fits = false
				}
			}
			if fits {
				count++
			}
		}
	}

	fmt.Println(count)
}
