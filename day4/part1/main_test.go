package main

import "testing"

func TestReadCards(t *testing.T) {
	want := 13
	have := TotalPoints("check.txt")

	if want != have {
		t.Fatalf("Want %d, have %d", want, have)
	}
}
