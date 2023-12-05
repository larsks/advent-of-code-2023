package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var (
	numbers map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
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
		digits := []int{}
		line := lines.Text()
		for pos := range line {
			if unicode.IsDigit(rune(line[pos])) {
				digits = append(digits, int(line[pos]-'0'))
			} else {
				for k, v := range numbers {
					if pos+len(k) <= len(line) && line[pos:pos+len(k)] == k {
						digits = append(digits, v)
					}
				}
			}
		}
		res := digits[0]*10 + digits[len(digits)-1]
		fmt.Printf("%s -> %v -> %d\n", line, digits, res)
		answer += res
	}

	return answer
}

func main() {
	answer := SumDigits("input.txt")
	fmt.Printf("sum: %d\n", answer)
}
