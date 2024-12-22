package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	secrets := make([]int64, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		n, _ := strconv.Atoi(line)
		secrets = append(secrets, int64(n))
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	buyers := make([][]int64, 0)
	for _, secret := range secrets {
		curr := secret
		buyer := make([]int64, 0)
		for i := 0; i < 2000; i++ {
			s := curr
			s = prune(mix(curr*int64(64), s))
			s = prune(mix(s/int64(32), s))
			s = prune(mix(s*int64(2048), s))
			buyer = append(buyer, s%10)
			curr = s
		}
		buyers = append(buyers, buyer)
	}

	type seqData struct {
		total  int64
		buyers map[int]int64
	}

	sequences := make(map[string]seqData)
	for b, buyer := range buyers {
		for i := 4; i < len(buyer); i++ {
			seq := make([]int64, 0)
			for j := i - 3; j <= i; j++ {
				change := buyer[j] - buyer[j-1]
				seq = append(seq, change)
			}

			seqData := seqData{
				total:  0,
				buyers: make(map[int]int64),
			}
			if exSeq, ok := sequences[seqStr(seq)]; ok {
				seqData = exSeq
			}

			if _, ok := seqData.buyers[b]; ok {
				continue
			}
			seqData.total += buyer[i]
			seqData.buyers[b] = buyer[i]

			sequences[seqStr(seq)] = seqData
		}
	}

	var bestSeqValue int64 = 0
	for _, seqData := range sequences {
		if seqData.total > bestSeqValue {
			bestSeqValue = seqData.total
		}
	}

	fmt.Println(bestSeqValue)
}

func seqStr(seq []int64) string {
	return fmt.Sprintf("%v", seq)
}

func mix(v, secret int64) int64 {
	return v ^ secret
}

func prune(secret int64) int64 {
	return secret % 16777216
}
