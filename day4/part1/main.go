package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type (
	Card struct {
		Id      int
		Have    []int
		Winning []int
	}
)

func NewCard(s string) *Card {
	x := strings.Split(s, ": ")
	label := x[0]
	cardid, err := strconv.Atoi(strings.Fields(label)[1])
	if err != nil {
		panic(err)
	}
	x = strings.Split(x[1], " | ")
	winning_s := strings.Fields(x[0])
	have_s := strings.Fields(x[1])

	winning := make([]int, len(winning_s))
	have := make([]int, len(have_s))

	for i, val := range winning_s {
		winning[i], _ = strconv.Atoi(val)
	}

	for i, val := range have_s {
		have[i], _ = strconv.Atoi(val)
	}

	return &Card{
		Id:      cardid,
		Have:    have,
		Winning: winning,
	}
}

func (c *Card) Points() int {
	count := 0
	for _, val := range c.Have {
		if slices.Contains(c.Winning, val) {
			count++
		}
	}

	if count == 0 {
		return 0
	}
	return int(math.Pow(2, float64(count-1)))
}

func ReadCards(filename string) []*Card {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	lines := bufio.NewScanner(fd)
	lines.Split(bufio.ScanLines)

	cards := []*Card{}
	for lines.Scan() {
		card := NewCard(lines.Text())
		cards = append(cards, card)
		fmt.Printf("%+v\n", card)
	}

	return cards
}

func TotalPoints(filename string) int {
	total := 0
	for _, card := range ReadCards(filename) {
		fmt.Printf("card %+v points %d\n", card, card.Points())
		total += card.Points()
	}

	return total
}

func main() {
	total := TotalPoints("input.txt")
	fmt.Printf("total: %d\n", total)
}
