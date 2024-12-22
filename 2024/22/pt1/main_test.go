package main

import (
	"fmt"
	"testing"
)

func TestMix(t *testing.T) {
	x := mix(15, 42)
	if x != 37 {
		t.Fatalf("expected 37, but got '%d'", x)
	}
}

func TestPrune(t *testing.T) {
	x := prune(100000000)
	if x != 16113920 {
		t.Fatalf("expected 37, but got '%d'", x)
	}
}

func TestPruneReal(t *testing.T) {
	fmt.Println(8650752 * int64(64))
	s := prune(8650752 * int64(64))
	t.Fatal(s)
}
