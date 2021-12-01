package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("13th/two/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	timeStrings := strings.Split(lines[1], ",")
	times := make([]*big.Int, 0)
	diffs := make([]*big.Int, 0)

	diff := -1
	for _, ts := range timeStrings {
		diff++
		if ts == "x" {
			continue
		}
		t, err := strconv.Atoi(ts)
		if err != nil {
			panic(err)
		}
		times = append(times, big.NewInt(int64(t)))
		diffs = append(diffs, big.NewInt(int64(diff)))
	}

	ans, err := crt(diffs, times)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans.String())
}

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(big.NewInt(1)) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return p.Sub(p, x.Mod(&x, p)), nil
}
