package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	x := 0
	y := 0
	dir := 90

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		action := line[:1]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch action {
		case "N":
			y += value
		case "E":
			x += value
		case "S":
			y -= value
		case "W":
			x -= value
		case "L":
			dir = (dir - value) % 360
		case "R":
			dir = (dir + value) % 360
		case "F":
			y += int(math.Cos(float64(dir) * math.Pi / 180) * float64(value))
			x += int(math.Sin(float64(dir) * math.Pi / 180) * float64(value))
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(math.Abs(float64(x))+math.Abs(float64(y)))
}
