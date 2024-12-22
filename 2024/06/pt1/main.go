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

	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		row := []rune(line)
		grid = append(grid, row)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(grid))
}

func solve(grid [][]rune) int {
	startI, startJ := findStart(grid)
	distinctPositions := map[string]bool{
		sqStr(startI, startJ): true,
	}

	dir := []int{-1, 0}
	pos := []int{startI, startJ}

	print(grid, dir, pos)

	for {
		newPos := []int{pos[0] + dir[0], pos[1] + dir[1]}
		if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= len(grid) || newPos[1] >= len(grid[0]) {
			break
		} else if grid[newPos[0]][newPos[1]] == '#' {
			dir = []int{dir[1], dir[0] * -1} // rotate 90 right
		} else {
			pos = newPos
			distinctPositions[sqStr(pos[0], pos[1])] = true
		}
		// print(grid, dir, pos)
	}

	return len(distinctPositions)
}

func sqStr(i, j int) string {
	return fmt.Sprintf("%s:%s", i, j)
}

func findStart(grid [][]rune) (int, int) {
	for i, r := range grid {
		for j, s := range r {
			if s == '^' {
				return i, j
			}
		}
	}
	panic("start not found")
}

func print(grid [][]rune, dir, pos []int) {
	for i, row := range grid {
		for j, sq := range row {
			if i == pos[0] && j == pos[1] {
				fmt.Print("^") // TODO dir
			} else if sq == '#' {
				fmt.Print(string([]rune{sq}))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
