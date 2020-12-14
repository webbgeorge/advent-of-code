package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("14th/two/input.txt")
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
				address: fmt.Sprintf("%036b", addr),
				value:   value,
			})
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	// lets just keep each bit as a single char in string
	data := make(map[string]*big.Int, 0)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, ins := range instructions {
		switch ins.t {
		case "mask":
			mask = ins.mask
		case "mem":
			addrMask := ""
			for i := 0; i < 36; i++ {
				switch mask[i] {
				case '0':
					addrMask += string(ins.address[i])
				case '1':
					addrMask += "1"
				case 'X':
					addrMask += "X"
				}
			}

			addresses := make([]string, 0)
			xs := strings.Count(addrMask, "X")
			combs := math.Pow(2, float64(xs))
			for i := 0; i < int(math.Round(combs)); i++ {
				a := addrMask
				f := fmt.Sprintf("%%0%db", xs)
				bin := fmt.Sprintf(f, i)
				for j := 0; j < len(bin); j++ {
					a = strings.Replace(a, "X", string(bin[j]), 1)
				}
				addresses = append(addresses, a)
			}

			for _, address := range addresses {
				data[address] = big.NewInt(int64(ins.value))
			}
		}
	}

	c := big.NewInt(0)
	for _, v := range data{
		c = c.Add(c, v)
	}

	fmt.Println(c.String())
}

type instruction struct {
	t       string
	mask    string
	address string
	value   int
}
