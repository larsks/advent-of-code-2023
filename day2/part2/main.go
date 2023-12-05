package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Bag struct {
		Red, Green, Blue int
	}

	Game struct {
		Id    int
		Pulls []*Bag
	}
)

func NewBag(red, green, blue int) *Bag {
	return &Bag{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func (bag *Bag) Allows(contents *Bag) bool {
	return bag.Red <= contents.Red && bag.Green <= contents.Green && bag.Blue <= contents.Blue
}

func (bag *Bag) String() string {
	return fmt.Sprintf("<red:%d, green:%d, blue:%d>", bag.Red, bag.Green, bag.Blue)
}

func (bag *Bag) Power() int {
	return bag.Red * bag.Green * bag.Blue
}

func (game *Game) Allows(contents *Bag) bool {
	for _, pull := range game.Pulls {
		if !pull.Allows(contents) {
			return false
		}
	}
	return true
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (game *Game) MinimumContents() *Bag {
	bag := Bag{}

	for _, pull := range game.Pulls {
		bag.Red = Max(bag.Red, pull.Red)
		bag.Green = Max(bag.Green, pull.Green)
		bag.Blue = Max(bag.Blue, pull.Blue)
	}

	return &bag
}

func ParseGame(line string) *Game {
	game := Game{}

	x := strings.Split(line, ": ")
	gameid, err := strconv.Atoi(strings.Split(x[0], " ")[1])
	if err != nil {
		panic(err)
	}
	game.Id = gameid

	pulls := strings.Split(x[1], "; ")
	for _, pull := range pulls {
		colors := strings.Split(pull, ", ")
		choices := &Bag{}
		for _, color := range colors {
			x := strings.Split(color, " ")
			count, err := strconv.Atoi(x[0])
			if err != nil {
				panic(err)
			}
			name := x[1]
			switch name {
			case "red":
				choices.Red = count
			case "green":
				choices.Green = count
			case "blue":
				choices.Blue = count
			}
		}
		game.Pulls = append(game.Pulls, choices)
	}
	return &game
}

func CheckGames(filename string, contents *Bag) int {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	lines := bufio.NewScanner(fd)
	lines.Split(bufio.ScanLines)
	final := 0

	for lines.Scan() {
		line := lines.Text()
		game := ParseGame(line)
		mins := game.MinimumContents()
		final += mins.Power()
	}

	return final
}

func main() {
	power := CheckGames("input.txt", NewBag(12, 13, 14))
	fmt.Printf("power: %d\n", power)
}
