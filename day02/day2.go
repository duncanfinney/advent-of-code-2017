package main

import (
	"math"
	"strings"

	"github.com/duncanfinney/advent-of-code-2017/lib"
)

func main() {
	input := lib.ReadInput("day2")

	println("part1=", part1(input))
	println("part2=", part2(input))
}

func part1(input string) int {
	checksum := 0

	for _, row := range strings.Split(input, "\n") {
		rowAsNums := rowAsNums(row)

		min := math.MaxInt32
		max := 0

		for _, cellVal := range rowAsNums {
			if cellVal > max {
				max = cellVal
			}
			if cellVal < min {
				min = cellVal
			}
		}
		checksum += (max - min)
	}

	return checksum
}

func rowAsNums(row string) []int {
	ret := []int{}
	items := strings.Fields(row)
	for _, cellStr := range items {
		cellAsNum := lib.Atoi(cellStr)
		ret = append(ret, cellAsNum)
	}
	return ret
}

func part2(input string) int {
	checksum := 0

loop:
	for _, rowStr := range strings.Split(input, "\n") {
		row := rowAsNums(rowStr)

		for idx1, n1 := range row {
			for idx2, n2 := range row {
				if n2%n1 == 0 && idx1 != idx2 {
					//found the one for this row
					checksum += (n2 / n1)
					continue loop
				}
			}

		}
	}
	return checksum
}
