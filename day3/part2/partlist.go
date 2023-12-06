package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type (
	PartList struct {
		Rows    []string
		Numbers [][][]int
		Symbols [][][]int
	}

	Part struct {
		Id              int
		Row, Start, End int
	}

	Symbol struct {
		Symbol      rune
		Row, Offset int
	}
)

func NewPartList(filename string) *PartList {
	p := &PartList{}

	if len(filename) > 0 {
		p.Read(filename)
	}

	return p
}

func (p *PartList) Read(filename string) error {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	p.Rows = []string{}
	renum := regexp.MustCompile(`\d+`)
	resym := regexp.MustCompile(`[^\d.]`)
	for scanner.Scan() {
		row := scanner.Text()
		p.Rows = append(p.Rows, row)
		nums := renum.FindAllIndex([]byte(row), -1)
		syms := resym.FindAllIndex([]byte(row), -1)
		p.Numbers = append(p.Numbers, nums)
		p.Symbols = append(p.Symbols, syms)
	}
	return nil
}

func (p *PartList) Parts() []int {
	parts := []int{}

	for rownum, row := range p.Numbers {
		for _, match := range row {
		OUTER:
			for i := rownum - 1; i <= rownum+1; i++ {
				if i < 0 || i >= len(p.Numbers) {
					continue
				}
				for _, symbol := range p.Symbols[i] {
					if symbol[0] >= match[0]-1 && symbol[0] <= match[1] {
						fmt.Printf("row %d match %v val %s\n", rownum, match, p.Rows[rownum][match[0]:match[1]])
						val, err := strconv.Atoi(p.Rows[rownum][match[0]:match[1]])
						if err != nil {
							panic(err)
						}

						parts = append(parts, val)
						break OUTER
					}
				}
			}
		}
	}

	return parts
}

func (p *PartList) Gears() [][]int {
	gears := [][]int{}

	for rownum, row := range p.Symbols {
		for _, match := range row {
			fmt.Printf("row %d match %v\n", rownum, match)
			val := p.Rows[rownum][match[0]]
			if val != '*' {
				continue
			}

			found := []int{}
			for i := rownum - 1; i <= rownum+1; i++ {
				if i < 0 || i >= len(p.Symbols) {
					continue
				}

				for _, number := range p.Numbers[i] {
					fmt.Printf("row %d  number %v\n", i, number)
					if match[0] >= number[0]-1 && match[0] <= number[1] {
						val, err := strconv.Atoi(p.Rows[i][number[0]:number[1]])
						if err != nil {
							panic(err)
						}
						found = append(found, val)
					}
				}
			}

			fmt.Printf("found: %v\n", gears)
			if len(found) == 2 {
				gears = append(gears, found)
			}
		}
	}

	return gears
}
