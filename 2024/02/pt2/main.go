package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("2024/02/pt2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	safeCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		report := strings.Fields(line)
		levels := make([]int, 0)
		for _, levelStr := range report {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		isSafe := isSafeTolerant(levels)
		if isSafe {
			// fmt.Println("safe", levels)
			safeCount++
		}

	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(safeCount)
}

func diffInt(n, m int) int {
	if n < m {
		return m - n
	}
	return n - m
}

func isSafeTolerant(levels []int) bool {
	isSafe := isSafeFn(levels)
	if isSafe {
		return true
	}
	for i := range levels {
		tLevels := make([]int, 0)
		for j, level := range levels {
			if j != i {
				tLevels = append(tLevels, level)
			}
		}
		if isSafeFn(tLevels) {
			return true
		}
	}
	return false
}

func isSafeFn(levels []int) bool {
	isSafe := true
	isInc := true
	for i := 0; i < len(levels); i++ {
		if i == 0 {
			if levels[i] > levels[i+1] {
				isInc = false
			}
			continue
		}

		if isInc && levels[i] < levels[i-1] {
			isSafe = false
			break
		} else if !isInc && levels[i] > levels[i-1] {
			isSafe = false
			break
		}

		diff := diffInt(levels[i], levels[i-1])
		if diff < 1 || diff > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}
