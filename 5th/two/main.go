package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	seatIDs := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		seatIDs = append(seatIDs, seatID(line))
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	sort.Ints(seatIDs)

	// skip first
	for i := 1; i < len(seatIDs); i++ {
		if seatIDs[i] - seatIDs[i-1] > 1 {
			fmt.Println(seatIDs[i]-1)
			return
		}
	}

	fmt.Println("nope")
}

func seatID(seat string) int {
	row := strings.ReplaceAll(seat[:7], "F", "0")
	row = strings.ReplaceAll(row, "B", "1")
	rowNumber, err := strconv.ParseInt(row, 2, 0)
	if err != nil {
		panic(err)
	}

	column := strings.ReplaceAll(seat[7:], "L", "0")
	column = strings.ReplaceAll(column, "R", "1")
	columnNumber, err := strconv.ParseInt(column, 2, 0)
	if err != nil {
		panic(err)
	}

	return (int(rowNumber) * 8) + int(columnNumber)
}
