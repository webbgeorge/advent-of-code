package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("expected args")
		os.Exit(1)
	}

	fp := os.Args[1]

	f, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rules := make([][]int, 0)
	updates := make([][]int, 0)

	isRules := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRules = false
			continue
		}
		if isRules {
			ruleStr := strings.Split(line, "|")
			rule := make([]int, 0)
			for _, rs := range ruleStr {
				n, _ := strconv.Atoi(rs)
				rule = append(rule, n)
			}
			rules = append(rules, rule)
		} else {
			updateStr := strings.Split(line, ",")
			update := make([]int, 0)
			for _, us := range updateStr {
				n, _ := strconv.Atoi(us)
				update = append(update, n)
			}
			updates = append(updates, update)
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(countMiddles(rules, updates))
}

func countMiddles(rules, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		fmt.Println(update)
		valid := true
		for _, rule := range rules {
			r1i := slices.Index(update, rule[0])
			r2i := slices.Index(update, rule[1])
			if r1i != -1 && r2i != -1 && r1i > r2i {
				valid = false
			}
		}
		fmt.Println(update, valid)
		if valid {
			sum += update[int(math.Ceil(float64(len(update)/2)))]
		}
		fmt.Println("")
	}
	return sum
}
