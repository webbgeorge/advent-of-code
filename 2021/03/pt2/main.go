package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("2021/03/pt2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	o2Str := bitCheck(lines, 0, o2Check)
	co2Str := bitCheck(lines, 0, co2Check)

	o2, err := strconv.ParseInt(o2Str, 2, 64)
	if err != nil {
		panic(err)
	}
	co2, err := strconv.ParseInt(co2Str, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("o2: %d, co2: %d, product: %d", o2, co2, o2*co2)
}

func bitCheck(lines []string, bitIndex int, getCheckBit func(zeros, ones int) rune) string {
	zeros := 0
	ones := 0
	for _, line := range lines {
		switch line[bitIndex] {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}

	check := getCheckBit(zeros, ones)

	newLines := make([]string, 0)
	for _, line := range lines {
		if rune(line[bitIndex]) == check {
			newLines = append(newLines, line)
		}
	}

	if len(newLines) == 1 {
		return newLines[0]
	}

	return bitCheck(newLines, bitIndex + 1, getCheckBit)
}

func o2Check(zeros, ones int) rune {
	if zeros > ones {
		return '0'
	} else if zeros < ones {
		return '1'
	}
	return '1'
}

func co2Check(zeros, ones int) rune {
	if zeros > ones {
		return '1'
	} else if zeros < ones {
		return '0'
	}
	return '0'
}
