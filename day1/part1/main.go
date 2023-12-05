package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func SumDigits(filename string) int {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	lines := bufio.NewScanner(fd)
	lines.Split(bufio.ScanLines)
	answer := 0
	for lines.Scan() {
		digits := []rune{}
		line := lines.Text()
		for _, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, c)
			}
		}

		res := int(digits[0]-'0')*10 + int(digits[len(digits)-1]-'0')
		answer += res
		fmt.Printf("%s -> %v -> %d\n", line, string(digits), res)
	}

	return answer
}

func main() {
	answer := SumDigits("input.txt")
	fmt.Printf("sum: %d\n", answer)
}
