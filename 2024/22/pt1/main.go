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

	var sum int64 = 0
	for _, secret := range secrets {
		curr := secret
		for i := 0; i < 2000; i++ {
			s := curr
			s = prune(mix(curr*int64(64), s))
			s = prune(mix(s/int64(32), s))
			s = prune(mix(s*int64(2048), s))
			curr = s
		}
		sum += curr
	}

	fmt.Println(sum)
}

func mix(v, secret int64) int64 {
	return v ^ secret
}

func prune(secret int64) int64 {
	return secret % 16777216
}
