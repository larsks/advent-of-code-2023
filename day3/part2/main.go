package main

import (
	"fmt"
)

func ReadGears(filename string) int {
	partlist := NewPartList(filename)
	total := 0
	for _, gear := range partlist.Gears() {
		total += gear[0] * gear[1]
	}
	return total
}

func main() {
	have := ReadGears("input.txt")
	fmt.Printf("have: %d\n", have)
}
