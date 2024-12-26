package main

import (
	"bufio"
	"cmp"
	"fmt"
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

	codes := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		codes = append(codes, line)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	numpadShortestPaths = make(map[rune]map[rune]string)
	dirpadShortestPaths = make(map[rune]map[rune]string)
	preComputeShortestPaths(getNumpad(), numpadButtons, &numpadShortestPaths)
	preComputeShortestPaths(getDirpad(), dirpadButtons, &dirpadShortestPaths)

	fmt.Println("---")
	for from, fromData := range dirpadShortestPaths {
		for to, path := range fromData {
			fmt.Println(string([]rune{from}), string([]rune{to}), path)
		}
	}
	for from, fromData := range numpadShortestPaths {
		for to, path := range fromData {
			fmt.Println(string([]rune{from}), string([]rune{to}), path)
		}
	}
	fmt.Println("---")

	sum := 0
	for _, code := range codes {
		nCode, _ := strconv.Atoi(strings.ReplaceAll(code, "A", ""))
		length := recursePaths(numpadShortestPaths, 26, code)
		fmt.Println(nCode, length, length*nCode)
		fmt.Println("---")
		sum += length * nCode
	}

	fmt.Println(sum)
}

type lensKey struct {
	path  string
	depth int
}

var lens = map[lensKey]int{}

func recursePaths(pad map[rune]map[rune]string, depth int, path string) int {
	if l, ok := lens[lensKey{path, depth}]; ok {
		return l
	}

	if depth == 0 {
		lens[lensKey{path, depth}] = len(path)
		return len(path)
	}

	currentKey := 'A'
	length := 0
	for _, key := range []rune(path) {
		shortestPath := pad[currentKey][key]
		length += recursePaths(dirpadShortestPaths, depth-1, shortestPath)
		currentKey = key
	}

	lens[lensKey{path, depth}] = length
	return length
}

func getNumpad() [][]int {
	return [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{-1, 0, 0},
	}
}

func getDirpad() [][]int {
	return [][]int{
		{-1, 0, 0},
		{0, 0, 0},
	}
}

var (
	num1 = []int{2, 0}
	num2 = []int{2, 1}
	num3 = []int{2, 2}
	num4 = []int{1, 0}
	num5 = []int{1, 1}
	num6 = []int{1, 2}
	num7 = []int{0, 0}
	num8 = []int{0, 1}
	num9 = []int{0, 2}
	num0 = []int{3, 1}
	numA = []int{3, 2}
	dirU = []int{0, 1}
	dirL = []int{1, 0}
	dirD = []int{1, 1}
	dirR = []int{1, 2}
	dirA = []int{0, 2}
)

var numpadButtons = map[rune][]int{
	'A': numA,
	'0': num0,
	'1': num1,
	'2': num2,
	'3': num3,
	'4': num4,
	'5': num5,
	'6': num6,
	'7': num7,
	'8': num8,
	'9': num9,
}

var dirpadButtons = map[rune][]int{
	'A': dirA,
	'^': dirU,
	'<': dirL,
	'v': dirD,
	'>': dirR,
}

var moves = [][]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}
var pressMove = []int{0, 0}

var (
	numpadShortestPaths map[rune]map[rune]string
	dirpadShortestPaths map[rune]map[rune]string
)

func preComputeShortestPaths(grid [][]int, buttons map[rune][]int, computed *map[rune]map[rune]string) {
	for r, b := range buttons {
		(*computed)[r] = make(map[rune]string)
		for r2, b2 := range buttons {
			if r == r2 {
				(*computed)[r][r2] = pathToString([][]int{pressMove})
				continue
			}
			paths := findPaths(grid, b, b2)
			(*computed)[r][r2] = shortestPath(paths)
		}
	}
}

var sortReplacer = strings.NewReplacer(
	"<", "a",
	"v", "b",
	">", "d",
	"^", "c",
	"A", "e",
)

func shortestPath(paths [][][]int) string {
	var shortest [][]int
	for _, path := range paths {
		if shortest == nil || len(path) < len(shortest) {
			shortest = copyGrid(path)
		}
	}

	allShortest := make([]string, 0)
	for _, path := range paths {
		if len(shortest) == len(path) {
			allShortest = append(allShortest, pathToString(path))
		}
	}

	highest := 0
	newAllShortest := make([]string, 0)
	for _, path := range allShortest {
		n := strings.Count(path, "<<") * 2
		n += strings.Count(path, ">>") * 2
		n += strings.Count(path, "^^") * 2
		n += strings.Count(path, "vv") * 2
		n += strings.Count(path, "vvv") * 3
		n += strings.Count(path, "^^^") * 3
		if n > highest {
			highest = n
			newAllShortest = []string{path}
		} else if n == highest {
			newAllShortest = append(newAllShortest, path)
		}
	}
	allShortest = newAllShortest

	slices.SortFunc(allShortest, func(a, b string) int {
		aa := sortReplacer.Replace(a)
		bb := sortReplacer.Replace(b)
		return cmp.Compare(aa, bb)
	})

	return allShortest[0]
}

func findPaths(grid [][]int, from, to []int) [][][]int {
	pos := []int{from[0], from[1]}
	visited := copyGrid(grid)
	visited[from[0]][from[1]] = -1
	path := make([][]int, 0)
	validPaths := make([][][]int, 0)
	traverse(grid, visited, pos, to, path, &validPaths)
	return validPaths
}

func traverse(grid [][]int, visited [][]int, pos, to []int, path [][]int, validPaths *[][][]int) {
	for _, move := range moves {
		newPos := []int{pos[0] + move[0], pos[1] + move[1]}

		// check bounds
		if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= len(grid) || newPos[1] >= len(grid[0]) {
			continue
		}

		// invalid square, or was already visited
		if grid[newPos[0]][newPos[1]] == -1 || visited[newPos[0]][newPos[1]] == -1 {
			continue
		}

		// update visited
		newVisited := copyGrid(visited)
		newVisited[newPos[0]][newPos[1]] = -1
		newPath := copyGrid(path)
		newPath = append(newPath, move)

		// new pos is the destination, save an continue looking for more paths
		if newPos[0] == to[0] && newPos[1] == to[1] {
			newPath = append(newPath, pressMove)
			*validPaths = append(*validPaths, newPath)
			continue
		}

		traverse(grid, newVisited, newPos, to, newPath, validPaths)
	}
}

func copyGrid(grid [][]int) [][]int {
	newGrid := make([][]int, 0)
	for _, row := range grid {
		newRow := make([]int, 0)
		for _, sq := range row {
			newRow = append(newRow, sq)
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func pathToString(path [][]int) string {
	s := make([]rune, 0)
	for _, mv := range path {
		var c rune
		switch fmt.Sprintf("%d, %d", mv[0], mv[1]) {
		case "-1, 0":
			c = '^'
		case "0, 1":
			c = '>'
		case "1, 0":
			c = 'v'
		case "0, -1":
			c = '<'
		case "0, 0":
			c = 'A'
		}
		s = append(s, c)
	}
	return string(s)
}

func printPathDebugString(path [][]int) {
	fmt.Println(pathToString(path))
}

func moveToDirButton(move []int) rune {
	if move[0] == 0 {
		if move[1] == 1 {
			return '>'
		} else if move[1] == 0 {
			return 'A'
		} else {
			return '<'
		}
	} else {
		if move[0] == 1 {
			return 'v'
		} else {
			return '^'
		}
	}
}
