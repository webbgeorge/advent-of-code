package main

import (
	"bufio"
	"fmt"
	"os"
)

var searchWord []rune = []rune("XMAS")

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

	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		gridLine := []rune(line)
		grid = append(grid, gridLine)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	count := 0
	count += getForDirection(grid, 0, 1)
	count += getForDirection(grid, 0, -1)
	count += getForDirection(grid, 1, 0)
	count += getForDirection(grid, -1, 0)
	count += getForDirection(grid, 1, 1)
	count += getForDirection(grid, 1, -1)
	count += getForDirection(grid, -1, 1)
	count += getForDirection(grid, -1, -1)

	fmt.Println(count)
}

func getForDirection(grid [][]rune, dirI, dirJ int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			found := true
			for k := 0; k < len(searchWord); k++ {
				newI := k*dirI + i
				newJ := k*dirJ + j
				if !exists(grid, newI, newJ) || grid[newI][newJ] != searchWord[k] {
					found = false
					break
				}
			}
			if found {
				count++
			}
		}
	}
	return count
}

func exists(grid [][]rune, newI, newJ int) bool {
	if newI >= 0 && newJ >= 0 && len(grid) > newI && len(grid[newI]) > newJ {
		return true
	}
	return false
}
