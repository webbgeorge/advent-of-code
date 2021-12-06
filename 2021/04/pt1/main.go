package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cardNumber struct {
	number   int
	selected bool
}

func main() {
	f, err := os.Open("2021/04/pt1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var pickedNumbers []int
	cards := make([][][]cardNumber, 0)

	scanner := bufio.NewScanner(f)
	scanner.Split(splitByDoubleNewLine)
	for scanner.Scan() {
		block := strings.TrimSpace(scanner.Text())
		if block == "" {
			continue
		}

		if pickedNumbers == nil {
			pickedNumbers = make([]int, 0)
			for _, s := range strings.Split(block, ",") {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				pickedNumbers = append(pickedNumbers, n)
			}
			continue
		}

		card := make([][]cardNumber, 0)
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			numbers := strings.Fields(line)
			cardLine := make([]cardNumber, 0)
			for _, s := range numbers {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				cardLine = append(cardLine, cardNumber{
					number:   n,
					selected: false,
				})
			}
			card = append(card, cardLine)
		}
		cards = append(cards, card)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	for _, pick := range pickedNumbers {
		fmt.Printf("%d\n\n", pick)
		for _, card := range cards {
			updateCard(card, pick)
			printCard(card)
			if checkCard(card) {
				fmt.Printf("score %d", score(card, pick))
				return
			}
		}
		fmt.Print("\n\n\n")
	}
}

func updateCard(card [][]cardNumber, pickedNumber int) {
	for i := 0; i < len(card); i++ {
		for j := 0; j < len(card[i]); j++ {
			if card[i][j].number == pickedNumber {
				card[i][j].selected = true
			}
		}
	}
}

func checkCard(card [][]cardNumber) bool {
	xyLen := len(card) // is a square
	for i := 0; i < xyLen; i++ {
		rowComplete := true
		colComplete := true
		for j := 0; j < xyLen; j++ {
			if !card[i][j].selected {
				rowComplete = false
			}
			if !card[j][i].selected {
				colComplete = false
			}
		}
		if rowComplete {
			return true
		}
		if colComplete {
			return true
		}
	}
	return false
}

func score(card [][]cardNumber, pickedNumber int) int {
	score := 0
	for i := 0; i < len(card); i++ {
		for j := 0; j < len(card[i]); j++ {
			if !card[i][j].selected {
				score += card[i][j].number
			}
		}
	}
	return score*pickedNumber
}

func printCard(card [][]cardNumber) {
	for _, row := range card {
		line := ""
		for _, n := range row {
			line = line + fmt.Sprintf("%2d", n.number)
			if n.selected {
				line = line + "* "
			} else {
				line = line + "  "
			}
		}
		fmt.Println(line)
	}
	fmt.Print("\n")
}

// implementation of a scanner func which splits by 2 new lines instead of by line
func splitByDoubleNewLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0 : i+1]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
