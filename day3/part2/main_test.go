package main

import "testing"

func TestReadParts(t *testing.T) {
	want := 467835
	have := ReadGears("check.txt")
	if want != have {
		t.Fatalf("Want %d, have %d", want, have)
	}
}
