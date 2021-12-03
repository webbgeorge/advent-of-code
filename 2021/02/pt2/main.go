package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("2021/02/pt2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	x := 0
	aim := 0
	z := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		words := strings.Split(line, " ")
		if len(words) != 2 {
			panic("invalid instruction")
		}
		inst := words[0]
		n, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}

		switch inst {
		case "forward":
			x += n
			z += aim*n
		case "up":
			aim -= n
		case "down":
			aim += n
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("X: %d, Z: %d, product: %d", x, z, z*x)
}
