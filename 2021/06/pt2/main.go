package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := "3,1,4,2,1,1,1,1,1,1,1,4,1,4,1,2,1,1,2,1,3,4,5,1,1,4,1,3,3,1,1,1,1,3,3,1,3,3,1,5,5,1,1,3,1,1,2,1,1,1,3,1,4,3,2,1,4,3,3,1,1,1,1,5,1,4,1,1,1,4,1,4,4,1,5,1,1,4,5,1,1,2,1,1,1,4,1,2,1,1,1,1,1,1,5,1,3,1,1,4,4,1,1,5,1,2,1,1,1,1,5,1,3,1,1,1,2,2,1,4,1,3,1,4,1,2,1,1,1,1,1,3,2,5,4,4,1,3,2,1,4,1,3,1,1,1,2,1,1,5,1,2,1,1,1,2,1,4,3,1,1,1,4,1,1,1,1,1,2,2,1,1,5,1,1,3,1,2,5,5,1,4,1,1,1,1,1,2,1,1,1,1,4,5,1,1,1,1,1,1,1,1,1,3,4,4,1,1,4,1,3,4,1,5,4,2,5,1,2,1,1,1,1,1,1,4,3,2,1,1,3,2,5,2,5,5,1,3,1,2,1,1,1,1,1,1,1,1,1,3,1,1,1,3,1,4,1,4,2,1,3,4,1,1,1,2,3,1,1,1,4,1,2,5,1,2,1,5,1,1,2,1,2,1,1,1,1,4,3,4,1,5,5,4,1,1,5,2,1,3"

	lfs := make(map[int]int)
	for _, lf := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(lf)
		lfs[i]++
	}

	start := time.Now()
	for day := 1; day <= 256; day++ {
		fmt.Printf("day %d, count %d, after %f\n", day, len(lfs), time.Now().Sub(start).Seconds())
		newLfs := make(map[int]int)
		for t, n := range lfs {
			if t == 0 {
				newLfs[6] += n
				newLfs[8] += n
			} else {
				newLfs[t-1] += n
			}
		}
		lfs = newLfs
	}

	sum := 0
	for _, n := range lfs {
		sum += n
	}

	fmt.Println(sum)
}