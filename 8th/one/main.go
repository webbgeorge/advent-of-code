package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`^(\w{3}) \+?(-?\d+)$`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	instructions := make([]instruction, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		matches := re.FindAllStringSubmatch(line, -1)

		value, err := strconv.Atoi(matches[0][2])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, instruction{
			op:    matches[0][1],
			value: value,
		})

	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	i := 0
	acc := 0
	for {
		if i >= len(instructions) {
			break
		}

		fmt.Printf("%d: %s %d\n", i, instructions[i].op, instructions[i].value)
		if instructions[i].called {
			fmt.Printf("repeat instruction i: '%d', acc: %d\n", i, acc)
			break
		}
		instructions[i].called = true

		switch instructions[i].op {
		case "acc":
			acc += instructions[i].value
			i++
		case "nop":
			i++
		case "jmp":
			i += instructions[i].value
		}
	}

	fmt.Println("nope")
}

type instruction struct {
	op     string
	value  int
	called bool
}
