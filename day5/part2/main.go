package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type (
	Almanac struct {
		Seeds       [][]int
		Maps        map[string][][]int
		Mapsequence []string
	}
)

func ReadInput(filename string) (a *Almanac) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	lines := bufio.NewScanner(fd)
	lines.Split(bufio.ScanLines)

	a = &Almanac{Maps: make(map[string][][]int)}
	mapname := ""

	for lines.Scan() {
		line := lines.Text()
		if len(line) == 0 {
			continue
		}

		if line[:6] == "seeds:" {
			a.Seeds = rangelist2i(strings.Fields(line[6:]))
		} else if line[len(line)-5:] == " map:" {
			mapname = line[:len(line)-5]
			a.Maps[mapname] = [][]int{}
			a.Mapsequence = append(a.Mapsequence, mapname)
		} else {
			a.Maps[mapname] = append(a.Maps[mapname], s2i(strings.Fields(line)))
		}
	}

	return
}

func mapvalue(val int, spec []int) (int, error) {
	if val < spec[1] || val > spec[1]+spec[2] {
		return 0, fmt.Errorf("value %d out of bounds for %v", val, spec)
	}

	return spec[0] + (val - spec[1]), nil
}

func main() {
	a := ReadInput("check.txt")
	lowest := math.MaxInt64

	mapped := a.Seeds
	for _, mapname := range a.Mapsequence {
		newmapped := [][]int{}

		for _, sr := range mapped {
			fmt.Printf("range %v\n", sr)
			thismap := a.Maps[mapname]
			for _, spec := range thismap {
				var start, length int
				fmt.Printf("check range %v spec %v\n", sr, spec)
				// if the seed range is partially contained in this map entry
				start, err := mapvalue(sr[0], spec)
				if err != nil {
					continue
				}
				fmt.Printf("use range %v spec %v start %d\n", sr, spec, start)
				if sr[0]+sr[1] > spec[1]+spec[2] {
					length = spec[2] - (sr[0] - spec[1])
					sr = []int{sr[0] + length, sr[1] - length}
				} else {
					length = sr[1]
					sr = []int{}
				}
				newmapped = append(newmapped, []int{start, length})
				if len(sr) == 0 {
					break
				}
			}

			if len(sr) > 0 {
				newmapped = append(newmapped, sr)
			}
		}

		fmt.Printf("before %v after %v\n\n", mapped, newmapped)
		mapped = newmapped
		newmapped = [][]int{}
	}

	fmt.Printf("mapped: %v\n", mapped)

	for _, x := range mapped {
		if x[0] < lowest {
			lowest = x[0]
		}
	}

	fmt.Printf("lowest: %d\n", lowest)
}
