package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	Map struct {
		Entries []MapEntry
	}

	MapEntry struct {
		DstStart, SrcStart, Length int
	}

	SeedData struct {
		Seeds       []int
		Maps        map[string]*Map
		MapSequence []string
	}
)

func NewSeedData() *SeedData {
	sd := SeedData{
		Maps: make(map[string]*Map),
	}

	return &sd
}

func (s *SeedData) ChainMaps(seed int) int {
	res := seed
	for _, name := range s.MapSequence {
		res = s.Maps[name].Map(res)
	}
	return res
}

func NewMap() *Map {
	return &Map{}
}

func (m *Map) Map(src int) int {
	for _, entry := range m.Entries {
		dst, err := entry.Map(src)
		if err != nil {
			return dst
		}
	}
	return src
}

func (e MapEntry) Map(src int) (int, error) {
	if !e.Contains(src) {
		return 0, fmt.Errorf("value %d out of range", src)
	}

	return e.DstStart + (src - e.SrcStart), nil
}

func (e MapEntry) Contains(src int) bool {
	return src >= e.SrcStart && src <= e.SrcStart+e.Length
}

func ReadInput(filename string) *SeedData {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	lines := bufio.NewScanner(fd)
	lines.Split(bufio.ScanLines)

	for lines.Scan() {
		line := lines.Text()
	}
}

func main() {
	fmt.Println("vim-go")
}
