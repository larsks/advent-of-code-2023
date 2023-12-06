package main

import (
	"fmt"
)

func ReadParts(filename string) int {
	partlist := NewPartList(filename)
	total := 0
	for _, part := range partlist.Parts() {
		total += part
	}
	return total
}

func main() {
	have := ReadParts("input.txt")
	fmt.Printf("have: %d\n", have)
}
