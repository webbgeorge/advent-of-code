package main

import "fmt"

func main() {
	//z := 2020
	z := 30000000
	//numbers := map[int]int{0: 1, 3: 2, 6: 3}
	numbers := map[int]int{8: 1, 13: 2, 1: 3, 0: 4, 18: 5, 9: 6}
	lastNumberTurn := 0
	for i := len(numbers) + 1; i <= z; i++ {
		n := i - 1 - lastNumberTurn
		if lastNumberTurn == 0 {
			n = 0
		}
		lastNumberTurn = numbers[n]
		numbers[n] = i
		if i == z {
			fmt.Println(n)
		}
	}
}
