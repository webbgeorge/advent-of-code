package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("2024/01/pt2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	l1 := make([]int64, 0)
	l2 := make([]int64, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		splitLine := strings.Fields(line)
		l1i, _ := strconv.ParseInt(splitLine[0], 10, 64)
		l2i, _ := strconv.ParseInt(splitLine[1], 10, 64)
		l1 = append(l1, l1i)
		l2 = append(l2, l2i)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	if len(l1) != len(l2) {
		panic("lens don't match")
	}

	var sum int64 = 0
	for i := 0; i < len(l1); i++ {
		var count int64 = 0
		for j := 0; j < len(l2); j++ {
			if l2[j] == l1[i] {
				count++
			}
		}
		sum += count * l1[i]
	}
	fmt.Println(sum)
}
