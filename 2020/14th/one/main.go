package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("14th/one/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	maskRE := regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRE := regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)

	instructions := make([]instruction, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		switch line[:4] {
		case "mask":
			matches := maskRE.FindAllStringSubmatch(line, -1)
			instructions = append(instructions, instruction{
				t:    "mask",
				mask: matches[0][1],
			})
		case "mem[":
			matches := memRE.FindAllStringSubmatch(line, -1)
			addr, err := strconv.Atoi(matches[0][1])
			if err != nil {
				panic(err)
			}
			value, err := strconv.Atoi(matches[0][2])
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, instruction{
				t:       "mem",
				address: addr,
				value:   value,
			})
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	// lets just keep each bit as a single char in string
	data := make(map[int]string, 0)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, ins := range instructions {
		switch ins.t {
		case "mask":
			mask = ins.mask
		case "mem":
			b := fmt.Sprintf("%036b", ins.value)
			data[ins.address] = ""
			for i := 0; i < 36; i++ {
				switch mask[i] {
				case '0':
					data[ins.address] += "0"
				case '1':
					data[ins.address] += "1"
				case 'X':
					data[ins.address] += string(b[i])
				}
			}
		}
	}

	c := big.NewInt(0)
	for _, v := range data {
		n := &big.Int{}
		n.SetString(v, 2)
		c = c.Add(c, n)
	}

	fmt.Println(c.String())
}

type instruction struct {
	t       string
	mask    string
	address int
	value   int
}
