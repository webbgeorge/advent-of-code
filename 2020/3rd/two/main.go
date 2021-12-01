package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := make([][]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		gridLine := strings.Split(line, "")
		grid = append(grid, gridLine)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	slopes := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	m := 1
	for _, slope := range slopes {
		trees := treeCount(grid, slope[0], slope[1])
		fmt.Printf("i: %d, j: %d, count: %d\n", slope[0], slope[1], trees)
		m = m * trees
	}

	fmt.Println(m)
}

func treeCount(grid [][]string, iSlope, jSlope int) int {
	treeCount := 0
	ii := 0
	for i := 0; i < len(grid); i++ {
		if i%iSlope != 0 {
			continue
		}

		if grid[i][(jSlope*ii)%len(grid[i])] == "#" {
			treeCount++
		}

		ii++
	}
	return treeCount
}
