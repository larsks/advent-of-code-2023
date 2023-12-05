package main

import (
	"slices"
	"testing"
)

func TestRunGames(t *testing.T) {
	bag := NewBag(12, 13, 14)
	want := []int{1, 2, 5}
	have := CheckGames("check.txt", bag)
	if !slices.Equal(have, want) {
		t.Fatalf("Want %v, have %v", want, have)
	}
}
