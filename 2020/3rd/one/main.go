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

	treeCount := 0
	for i := 0; i < len(grid); i++ {
		if grid[i][(3*i)%len(grid[i])] == "#" {
			treeCount++
		}
	}

	fmt.Println(treeCount)
}
