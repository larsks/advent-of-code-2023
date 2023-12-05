package main

import "testing"

func TestSumDigits(t *testing.T) {
	want := 142
	have := SumDigits("check.txt")
	if want != have {
		t.Fatalf("Wanted %d, got %d", want, have)
	}
}
