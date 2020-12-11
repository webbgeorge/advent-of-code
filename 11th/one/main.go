package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("11th/one/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	seats := make([][]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, "")
		seats = append(seats, splitLine)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	for {
		newSeats := make([][]string, 0)
		for i := 0; i < len(seats); i++ {
			newRow := make([]string, 0)
			for j := 0; j < len(seats[i]); j++ {
				s := seats[i][j]
				n := occupiedAdjacentSeats(i, j, seats)
				switch seats[i][j] {
				case "L":
					if n == 0 {
						s = "#"
					}
				case "#":
					if n >= 4 {
						s = "L"
					}
				}
				newRow = append(newRow, s)
			}
			newSeats = append(newSeats, newRow)
		}
		if compareSeats(seats, newSeats) {
			break
		}
		seats = newSeats
	}

	n := 0
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == "#" {
				n++
			}
		}
	}
	fmt.Println(n)
}

func occupiedAdjacentSeats(i, j int, seats [][]string) int {
	n := 0
	for k := i-1; k <= i+1; k++ {
		for l := j-1; l <= j+1; l++ {
			// check for out of bounds
			if k < 0 || k >= len(seats) || l < 0 || l >= len(seats[0]) {
				continue
			}
			// don't count middle chair
			if k == i && l == j {
				continue
			}
			if seats[k][l] == "#" {
				n++
			}
		}
	}
	return n
}

func compareSeats(seats, newSeats [][]string) bool {
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] != newSeats[i][j] {
				return false
			}
		}
	}
	return true
}
