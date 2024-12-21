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
	f, err := os.Open("2024/01/pt1/input.txt")
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
		sum += diffInt(l1[i], l2[i])
	}
	fmt.Println(sum)
}

func diffInt(n, m int64) int64 {
	if n < m {
		return m - n
	}
	return n - m
}
