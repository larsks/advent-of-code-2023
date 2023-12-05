package main

import (
	"testing"
)

func TestRunGames(t *testing.T) {
	bag := NewBag(12, 13, 14)
	want := 2286
	have := CheckGames("check.txt", bag)
	if have != want {
		t.Fatalf("Want %v, have %v", want, have)
	}
}
