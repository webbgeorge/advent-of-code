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
		if !checkUpdate(rules, update) {
			fmt.Println(update)
			reorderedUpdate := reorderUpdate(rules, update)
			fmt.Println(reorderedUpdate)
			fmt.Println("")
			sum += reorderedUpdate[int(math.Ceil(float64(len(reorderedUpdate)/2)))]
		}
	}
	return sum
}

func reorderUpdate(rules [][]int, update []int) []int {
	updateList := make([]int, 0)
	for _, page := range update {
		for i := 0; i <= len(updateList); i++ {
			ul := make([]int, 0)
			for j := 0; j <= len(updateList); j++ {
				if i == j {
					ul = append(ul, page)
				} else if i < j {
					ul = append(ul, updateList[j-1])
				} else {
					ul = append(ul, updateList[j])
				}
			}
			if checkUpdate(rules, ul) {
				updateList = ul
				break
			}
		}
	}

	return updateList
}

func checkUpdate(rules [][]int, update []int) bool {
	valid := true
	for _, rule := range rules {
		r1i := slices.Index(update, rule[0])
		r2i := slices.Index(update, rule[1])
		if r1i != -1 && r2i != -1 && r1i > r2i {
			valid = false
		}
	}
	return valid
}
