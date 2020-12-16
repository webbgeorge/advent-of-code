package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("16th/one/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rules := make(map[string][]int)
	//myTicket := make([]int, 0)
	nearbyTickets := make([][]int, 0)

	ruleRE := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)

	section := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if line == "your ticket:" {
			section++
			continue
		}
		if line == "nearby tickets:" {
			section++
			continue
		}

		switch section {
		case 0:
			matches := ruleRE.FindAllStringSubmatch(line, -1)
			a, _ := strconv.Atoi(matches[0][2])
			b, _ := strconv.Atoi(matches[0][3])
			c, _ := strconv.Atoi(matches[0][4])
			d, _ := strconv.Atoi(matches[0][5])
			rules[matches[0][1]] = []int{a, b, c, d}
		case 1:
			ticket := make([]int, 0)
			tstrs := strings.Split(line, ",")
			for _, tstr := range tstrs {
				t, _ := strconv.Atoi(tstr)
				ticket = append(ticket, t)
			}
			//myTicket = ticket
		case 2:
			ticket := make([]int, 0)
			tstrs := strings.Split(line, ",")
			for _, tstr := range tstrs {
				t, _ := strconv.Atoi(tstr)
				ticket = append(ticket, t)
			}
			nearbyTickets = append(nearbyTickets, ticket)
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, ticket := range nearbyTickets {
		for _, field := range ticket {
			isValid := false
			for _, rule := range rules {
				if (field >= rule[0] && field <= rule[1]) || (field >= rule[2] && field <= rule[3]) {
					isValid = true
				}
			}
			if !isValid {
				sum += field
			}
		}
	}
	fmt.Println(sum)
}
