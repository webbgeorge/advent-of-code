package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type bag struct {
	colour   string
	contains []struct {
		colour string
		count  int
	}
}

var bagRE = regexp.MustCompile(`^\w+ \w+`)
var containsRE = regexp.MustCompile(`(\d) (\w+ \w+)`)

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

		colour := bagRE.FindString(line)
		matches := containsRE.FindAllStringSubmatch(line, -1)

		contains := make([]struct {
			colour string
			count  int
		}, 0)
		for _, c := range matches {
			containsColour := c[2]
			count, err := strconv.Atoi(c[1])
			if err != nil {
				panic(err)
			}
			contains = append(contains, struct {
				colour string
				count  int
			}{colour: containsColour, count: count})
		}
		bags[colour] = bag{
			colour:   colour,
			contains: contains,
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	sg := bags["shiny gold"]
	bagCount := countBags(sg, bags)
	fmt.Println(bagCount)
}

func countBags(b bag, bs map[string]bag) int {
	count := 0
	for _, c := range b.contains {
		count += c.count
		count += countBags(bs[c.colour], bs) * c.count
	}
	return count
}
