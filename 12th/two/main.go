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

	var wx float64 = 10
	var wy float64 = 1
	var sx float64 = 0
	var sy float64 = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		action := line[:1]
		value, err := strconv.ParseFloat(line[1:], 64)
		if err != nil {
			panic(err)
		}

		switch action {
		case "N":
			wy += value
		case "E":
			wx += value
		case "S":
			wy -= value
		case "W":
			wx -= value
		case "L":
			r := 0 - (value * math.Pi / 180)
			newx := wx*math.Cos(r) + wy*math.Sin(r)
			newy := wy*math.Cos(r) - wx*math.Sin(r)
			wx = newx
			wy = newy
		case "R":
			r := value * math.Pi / 180
			newx := wx*math.Cos(r) + wy*math.Sin(r)
			newy := wy*math.Cos(r) - wx*math.Sin(r)
			wx = newx
			wy = newy
		case "F":
			sx += wx * value
			sy += wy * value
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(math.Round(math.Abs(sx)+math.Abs(sy)))
}
