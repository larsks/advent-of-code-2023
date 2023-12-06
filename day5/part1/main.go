package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (s *SeedData) AddMap(mapname string) {
	s.Maps[mapname] = NewMap()
	s.MapSequence = append(s.MapSequence, mapname)
}

func (s *SeedData) ChainMap(seed int) int {
	res := seed
	fmt.Printf("---\n")
	for _, name := range s.MapSequence {
		newres := s.Maps[name].Map(res)
		fmt.Printf("map %s src %d dst %d\n", name, res, newres)
		res = newres
	}
	return res
}

func NewMap() *Map {
	return &Map{}
}

func (m *Map) AddEntry(entry []int) {
	m.Entries = append(m.Entries, MapEntry{entry[0], entry[1], entry[2]})
}

func (m *Map) Map(src int) int {
	for _, entry := range m.Entries {
		dst, err := entry.Map(src)
		if err == nil {
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
	state := 0
	mapname := ""
	sd := NewSeedData()

	for lines.Scan() {
		line := lines.Text()
		switch state {
		case 0:
			if len(line) == 0 {
				state = 1
			} else if line[:6] == "seeds:" {
				parts := strings.Split(line, ": ")
				sd.Seeds = s2i(strings.Fields(parts[1]))
			}
		case 1:
			if len(line) == 0 {
				continue
			} else if line[len(line)-5:] == " map:" {
				mapname = line[:len(line)-5]
				fmt.Printf("reading map %s\n", mapname)
				sd.AddMap(mapname)
				state = 2
			}
		case 2:
			if len(line) == 0 {
				state = 1
			} else {
				entry := s2i(strings.Fields(line))
				sd.Maps[mapname].AddEntry(entry)
			}
		}
	}

	return sd
}

func main() {
	sd := ReadInput("input.txt")
	fmt.Printf("%+v\n", sd)
	lowest := 0

	for _, seed := range sd.Seeds {
		mapped := sd.ChainMap(seed)
		if lowest == 0 || mapped < lowest {
			lowest = mapped
		}
		fmt.Printf("%d -> %d\n", seed, mapped)
	}

	fmt.Printf("lowest location: %d\n", lowest)
}
