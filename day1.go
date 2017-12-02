package main

import (
	"strings"

	"github.com/duncanfinney/advent-of-code-2017/lib"
)

func main() {
	input := lib.ReadInput("day1")

	println("part1=", part1(input))
	println("part2=", part2(input))
}

func part1(input string) int {
	return addItemsThatAreNItemsAhead(input, 1)
}

func part2(input string) int {
	return addItemsThatAreNItemsAhead(input, len(input)/2)
}

func addItemsThatAreNItemsAhead(input string, n int) int {
	split := strings.Split(input, "")

	total := 0
	for idx, char := range split {
		nextDigit := (idx + n) % len(split)
		if char == split[nextDigit] {
			total += lib.Atoi(char)
		}
	}
	return total
}
