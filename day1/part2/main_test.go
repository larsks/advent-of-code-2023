package main

import "testing"

func TestSumDigits(t *testing.T) {
	want := 281
	have := SumDigits("check.txt")
	if want != have {
		t.Fatalf("Wanted %d, got %d", want, have)
	}
}
