package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type vector struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	f, err := os.Open("2021/05/pt1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	vectors := make([]vector, 0)
	xMax := 0
	yMax := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
		matches := re.FindAllStringSubmatch(line, -1)

		if len(matches) != 1 {
			panic("unexpected matches")
		}

		x1, _ := strconv.Atoi(matches[0][1])
		y1, _ := strconv.Atoi(matches[0][2])
		x2, _ := strconv.Atoi(matches[0][3])
		y2, _ := strconv.Atoi(matches[0][4])

		if x1 > xMax {
			xMax = x1
		}
		if x2 > xMax {
			xMax = x2
		}
		if y1 > yMax {
			yMax = y1
		}
		if y2 > yMax {
			yMax = y2
		}

		vectors = append(vectors, vector{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	m := make([][]int, 0)
	for i := 0; i < xMax + 1; i++ {
		line := make([]int, 0)
		for j := 0; j < yMax + 1; j++ {
			line = append(line, 0)
		}
		m = append(m, line)
	}

	for _, v := range vectors {
		if v.x1 != v.x2 && v.y1 != v.y2 {
			continue
		}

		var xa, xb, ya, yb int
		if v.x1 <= v.x2 {
			xa = v.x1
			xb = v.x2
		} else {
			xa = v.x2
			xb = v.x1
		}
		if v.y1 <= v.y2 {
			ya = v.y1
			yb = v.y2
		} else {
			ya = v.y2
			yb = v.y1
		}

		for i := xa; i <= xb; i++ {
			for j := ya; j <= yb; j++ {
				m[i][j]++
			}
		}
	}

	//// Print grid
	//for _, n := range m {
	//	line := ""
	//	for _, count := range n {
	//		line = line + strconv.Itoa(count)
	//	}
	//	fmt.Println(line)
	//}

	sum := 0
	for _, n := range m {
		for _, count := range n {
			if count >= 2 {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
