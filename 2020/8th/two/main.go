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
	f, err := os.Open("8th/two/input.txt")
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
	for {
		acc, err := tryChangeOp(i, instructions)
		if err != nil {
			i++
			continue
		}
		fmt.Printf("acc: %d, attempt: %d\n", acc, i)
		break
	}
}

type instruction struct {
	op    string
	value int
}

func tryChangeOp(attempt int, instructions []instruction) (int, error) {
	i := 0
	ii := 0
	acc := 0
	triedInstructions := make(map[int]bool)
	for {
		if i >= len(instructions) {
			break
		}

		if triedInstructions[i] {
			return 0, fmt.Errorf("repeat instruction i: '%d', acc: %d\n", i, acc)
		}
		triedInstructions[i] = true

		switch instructions[i].op {
		case "acc":
			acc += instructions[i].value
			i++
		case "nop":
			if attempt == ii {
				// do jmp instead
				i += instructions[i].value
				ii++
			} else {
				i++
				ii++
			}
		case "jmp":
			if attempt == ii {
				// do nop instead
				i++
				ii++
			} else {
				i += instructions[i].value
				ii++
			}
		}
	}
	return acc, nil
}
