package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestReadInput(t *testing.T) {
	want := [][]int{
		{79, 82},
		{14, 43},
		{55, 86},
		{13, 35},
	}

	sd := ReadInput("check.txt")
	have := [][]int{}
	for _, seed := range sd.Seeds {
		have = append(have, []int{seed, sd.ChainMap(seed)})
	}

	for i := 0; i < len(want); i++ {
		if !slices.Equal(want[i], have[i]) {
			t.Fatalf("want %v, have %v\n", want, have)
		}
	}
}
