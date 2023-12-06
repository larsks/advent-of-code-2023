package main

import "testing"

func TestReadCards(t *testing.T) {
	want := 30
	have := TotalCards("check.txt")

	if want != have {
		t.Fatalf("Want %d, have %d", want, have)
	}
}
