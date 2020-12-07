package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var highestSeatID int64

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		id := seatID(line)
		if id > highestSeatID {
			highestSeatID = id
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(highestSeatID)
}

func seatID(seat string) int64 {
	row := strings.ReplaceAll(seat[:7], "F", "0")
	row = strings.ReplaceAll(row, "B", "1")
	rowNumber, err := strconv.ParseInt(row, 2, 64)
	if err != nil {
		panic(err)
	}

	column := strings.ReplaceAll(seat[7:], "L", "0")
	column = strings.ReplaceAll(column, "R", "1")
	columnNumber, err := strconv.ParseInt(column, 2, 64)
	if err != nil {
		panic(err)
	}

	return (rowNumber * 8) + columnNumber
}
