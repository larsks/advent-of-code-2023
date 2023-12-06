package main

import "testing"

func TestReadParts(t *testing.T) {
	want := 4361
	have := ReadParts("check.txt")
	if want != have {
		t.Fatalf("Want %d, have %d", want, have)
	}
}
