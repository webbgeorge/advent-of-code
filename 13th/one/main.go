package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("13th/one/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	start, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	times := strings.Split(lines[1], ",")

	bestWait := -1
	bestTime := -1
	for _, time := range times {
		if time == "x" {
			continue
		}
		t, err := strconv.Atoi(time)
		if err != nil {
			panic(err)
		}
		wait := t - (start % t)
		if bestWait == -1 || wait < bestWait{
			bestWait = wait
			bestTime = t
		}
	}

	fmt.Println(bestWait*bestTime)
}
