package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part Two: ", partTwo(string(input)))
}

func intoGrid(input string) Grid {
	grid := map[[2]int]rune{}
	for i, row := range strings.Split(input, "\n") {
		for j, c := range row {
			grid[[2]int{j, i}] = c
		}
	}
	return grid
}

type Grid map[[2]int]rune

func (grid Grid) isValid(x int, y int, c rune) bool {
	val, isSet := grid[[2]int{x, y}]

	return isSet && val == c
}

func check(grid Grid, x int, y int, toCheck map[[2]int]rune) bool {
	for point, expect := range toCheck {
		if !grid.isValid(x+point[0], y+point[1], expect) {
			return false
		}
	}

	return true
}

func partTwo(input string) int {
	grid := intoGrid(input)
	count := 0

	checks := []map[[2]int]rune{
		{
			{-1, 1}: 'M',
			{1, -1}: 'S',
			{1, 1}:  'M',
			{-1, -1}: 'S',
		},
		{
			{-1, 1}: 'S',
			{1, -1}: 'M',
			{1, 1}:  'S',
			{-1, -1}: 'M',
		},
		{
			{-1, 1}: 'M',
			{1, -1}: 'S',
			{1, 1}:  'S',
			{-1, -1}: 'M',
		},
		{
			{-1, 1}: 'S',
			{1, -1}: 'M',
			{1, 1}:  'M',
			{-1, -1}: 'S',
		},
	}

	for p, c := range grid {
		x, y := p[0], p[1]

		if c != 'A' {
			continue
		}

		for _, c := range checks {
			if check(grid, x, y, c) {
				count++
				break
			}
		}
	}

	return count
}
