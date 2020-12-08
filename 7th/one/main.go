package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type bag struct {
	colour   string
	contains []string
}

var re = regexp.MustCompile(`(\w+ \w+) bags?`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bags := make(map[string]bag)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		matches := re.FindAllStringSubmatch(line, -1)

		contains := make([]string, 0)
		for _, c := range matches[1:] {
			if c[1] == "no other" {
				continue
			}
			contains = append(contains, c[1])
		}
		bags[matches[0][1]] = bag{
			colour:   matches[0][1],
			contains: contains,
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	bagColours := 0
	for _, bag := range bags {
		if bag.colour == "shiny gold" {
			continue
		}
		if containsColour("shiny gold", bag, bags) {
			bagColours++
		}
	}
	fmt.Println(bagColours)
}

func containsColour(colour string, b bag, bs map[string]bag) bool {
	for _, c := range b.contains {
		if c == colour {
			return true
		}
		if containsColour(colour, bs[c], bs) {
			return true
		}
	}
	return false
}
