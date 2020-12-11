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
				switch seats[i][j] {
				case "L":
					n := occupiedAdjacentSeats(i, j, seats)
					if n == 0 {
						s = "#"
					}
				case "#":
					n := occupiedAdjacentSeats(i, j, seats)
					if n >= 5 {
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
	w := len(seats)
	if len(seats[0]) > w {
		w = len(seats[0])
	}
	w = w/2

	positions := make(map[string]bool)
	for o := 1; o <= w; o++ {
		for k := i-o; k <= i+o; k+=o {
			for l := j-o; l <= j+o; l+=o {
				// check for out of bounds
				if k < 0 || k >= len(seats) || l < 0 || l >= len(seats[0]) {
					continue
				}
				// don't count middle chair
				if k == i && l == j {
					continue
				}
				dir := getDir(k-i, l-j)
				_, ok := positions[dir]
				if !ok {
					if seats[k][l] == "#" {
						positions[dir] = true
					}
					if seats[k][l] == "L" {
						positions[dir] = false
					}
				}
			}
		}
	}

	n := 0
	for _, t := range positions {
		if t {
			n++
		}
	}
	return n
}

func getDir(n, m int) string {
	s := ""
	if n < 0 {
		s = s+"-1,"
	}
	if n == 0 {
		s = s+"0,"
	}
	if n > 0 {
		s = s+"1,"
	}
	if m < 0 {
		s = s+"-1"
	}
	if m == 0 {
		s = s+"0"
	}
	if m > 0 {
		s = s+"1"
	}
	return s
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
