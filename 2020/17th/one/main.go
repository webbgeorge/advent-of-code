package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type dimension map[int]map[int]map[int]bool
type plane map[int]map[int]bool
type line map[int]bool

func main() {
	f, err := os.Open("17th/one/test-input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pocketDimension := make(dimension, 0)

	p := make(plane)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		if lineStr == "" {
			continue
		}

		l := make(line)
		for j, s := range strings.Split(lineStr, "") {
			l[j] = false
			if s == "#" {
				l[j] = true
			}
		}
		p[i] = l
		i++
	}
	pocketDimension[0] = p

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("")
	for cycle := 0; cycle < 6; cycle++ {
		newDimension := make(dimension)
		for z, p := range pocketDimension {
			newPlane := make(plane)
			for y, l := range p {
				newLine := make(line)
				for x, cube := range l {
					activeNeighbours := countNeighbours(x, y, z, pocketDimension)
					if cube {
						if activeNeighbours == 2 || activeNeighbours == 3 {
							newLine[x] = true
						} else {
							newLine[x] = false
						}
					} else {
						if activeNeighbours == 3 {
							newLine[x] = true
						} else {
							newLine[x] = false
						}
					}
					setMissingNeighbours(x, y, z, pocketDimension) // this won't work as it doesn't set them on new data structure
				}
				newPlane[y] = newLine
			}
			newDimension[z] = newPlane
		}
		pocketDimension = newDimension
	}

	active := 0
	for _, plane := range pocketDimension {
		for _, line := range plane {
			for _, cube := range line {
				if cube {
					active++
				}
			}
		}
	}

	fmt.Println(active)
}

func countNeighbours(x, y, z int, pocketDimension dimension) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				if pocketDimension[z+i][y+j][x+k] {
					count++
				}
			}
		}
	}
	return count
}

func setMissingNeighbours(x, y, z int, pocketDimension dimension) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				if _, ok := pocketDimension[z+i][y+j][x+k]; !ok {
					pocketDimension[z+i][y+j][x+k] = false
				}
			}
		}
	}
	return count
}

//func printDimension(d dimension) {
//	ds := make([][][]struct{}, 0)
//	for _, plane := range pocketDimension {
//		ps := make([][]string, 0)
//		for _, line := range plane {
//			ls := make([][]string, 0)
//			for _, cube := range line {
//				if cube {
//					active++
//				}
//			}
//			ps = append(ps, ls)
//		}
//		ds = append(ds, ps)
//	}
//}
