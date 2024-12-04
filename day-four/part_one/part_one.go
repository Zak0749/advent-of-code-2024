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

	fmt.Println("Part One:", partOne(string(input)))
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

func check(grid Grid, x int, y int, dx int, dy int) bool {
	for n, r := range []rune{'M', 'A', 'S'} {
		i := x+((n + 1)*dx)
		j := y+((n+1)*dy)
		if !grid.isValid(i, j, r) {
			return false
		}
	}

	return true
}

func partOne(input string) int {
	grid := intoGrid(input)
	count := 0

	for p, c := range grid {
		x, y := p[0], p[1]

		if c != 'X' {
			continue
		}

		for _, dy := range []int{-1, 0, 1} {
			for _, dx := range []int{-1, 0, 1} {
				if dx == 0 && dy == 0 {
					continue
				}

				if check(grid, x, y, dx, dy) {
					count++
				}
			}
		}
	}

	return count
}
