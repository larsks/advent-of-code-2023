package main

import (
	"testing"
)

func TestReadInput(t *testing.T) {
	want := 46
	sd := ReadInput("check.txt")
	have := FindLowestLocation(sd)

	if have != want {
		t.Fatalf("have %d, want %d\n", have, want)
	}
}
