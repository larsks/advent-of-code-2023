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
		Count   int
	}

	ListOfCards []*Card
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
		Count:   1,
	}
}

func (c *Card) Wins() int {
	count := 0
	for _, val := range c.Have {
		if slices.Contains(c.Winning, val) {
			count++
		}
	}

	return count
}

func (c *Card) Points() int {
	count := c.Wins()

	if count == 0 {
		return 0
	}
	return int(math.Pow(2, float64(count-1)))
}

func ReadCards(filename string) ListOfCards {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	lines := bufio.NewScanner(fd)
	lines.Split(bufio.ScanLines)

	cards := ListOfCards{}
	for lines.Scan() {
		card := NewCard(lines.Text())
		cards = append(cards, card)
		fmt.Printf("%+v\n", card)
	}

	return cards
}

func PlayCards(cards ListOfCards) {
	for i, card := range cards {
		wins := card.Wins()
		fmt.Printf("card %d has %d wins\n", card.Id, wins)
		for c := 0; c < card.Count; c++ {
			for j := 1; j <= wins; j++ {
				cards[i+j].Count += 1
			}
		}
	}
}

func TotalCards(filename string) int {
	total := 0

	cards := ReadCards(filename)
	PlayCards(cards)
	for _, card := range cards {
		fmt.Printf("card %+v\n", card)
		total += card.Count
	}

	return total
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
	total := TotalCards("input.txt")
	fmt.Printf("total: %d\n", total)
}
