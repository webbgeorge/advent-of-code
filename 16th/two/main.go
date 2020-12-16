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
	f, err := os.Open("16th/two/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rules := make(map[string][]int)
	myTicket := make([]int, 0)
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
			myTicket = ticket
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

	// eliminate invalid tickets
	validNearbyTickets := make([][]int, 0)
	for _, ticket := range nearbyTickets {
		isTicketValid := true
		for _, field := range ticket {
			isFieldValid := false
			for _, rule := range rules {
				if (field >= rule[0] && field <= rule[1]) || (field >= rule[2] && field <= rule[3]) {
					isFieldValid = true
				}
			}
			if !isFieldValid {
				isTicketValid = false
			}
		}
		if isTicketValid {
			validNearbyTickets = append(validNearbyTickets, ticket)
		}
	}

	// calculate possible fields
	possibleRulesFields := make([]map[string]bool, 0)
	// loops through columns
	for i := 0; i < len(validNearbyTickets[0]); i++ {
		possibleRules := getPossibleRules(rules)
		// loop through tickets
		for j := 0; j < len(validNearbyTickets); j++ {
			for name, rule := range rules {
				field := validNearbyTickets[j][i]
				if !((field >= rule[0] && field <= rule[1]) || (field >= rule[2] && field <= rule[3])) {
					possibleRules[name] = false
				}
			}
		}
		possibleRulesFields = append(possibleRulesFields, possibleRules)
	}

	// get possibles into better format
	fields := make([][]string, 0)
	answers := make([]string, 0)
	for _, possibles := range possibleRulesFields {
		field := make([]string, 0)
		for name, b := range possibles {
			if b {
				field = append(field, name)
			}
		}
		fields = append(fields, field)
		answers = append(answers, "")
	}

	// reduce down answers

	for {
		done := 0
		for i := 0; i < len(fields); i++ {
			if len(fields[i]) == 0 {
				done++
				continue
			}
			if len(fields[i]) == 1 {
				answers[i] = fields[i][0]
				for j := 0; j < len(fields); j++ {
					newFields := make([]string, 0)
					for k := 0; k < len(fields[j]); k++ {
						if fields[j][k] != answers[i] {
							newFields = append(newFields, fields[j][k])
						}
					}
					fields[j] = newFields
				}
			}
		}
		if done == len(fields) {
			break
		}
	}

	m := 1
	for i, answer := range answers {
		if strings.HasPrefix(answer, "departure") {
			m *= myTicket[i]
		}
	}

	fmt.Printf("%+v\n", m)
	fmt.Printf("%+v\n", answers)
}

func getPossibleRules(rules map[string][]int) map[string]bool {
	possibleRules := make(map[string]bool)
	for name := range rules {
		possibleRules[name] = true
	}
	return possibleRules
}
