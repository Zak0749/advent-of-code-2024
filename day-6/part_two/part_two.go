package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part Two: ", partTwo(string(input)))
}

type Vec2 struct {
	I, J int
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.I + v2.I, v.J + v2.J}
}

type Grid map[Vec2]rune

type Guard struct {
	Pos Vec2
	Dir Vec2
}

func (g Guard) Move(grid Grid) Guard {
	nextPos := g.Pos.Add(g.Dir)
	for grid[nextPos] == '#' {
		g.Dir = Vec2{-g.Dir.J, g.Dir.I}
		nextPos = g.Pos.Add(g.Dir)
	}

	g.Pos = nextPos
	return g
}

func parseInput(input string) (Grid, Guard) {
	grid := make(map[Vec2]rune)
	var guardPos Vec2
	for j, line := range strings.Split(input, "\n") {
		for i, char := range line {
			grid[Vec2{i, j}] = char
			if char == '^' {
				guardPos = Vec2{i, j}
			}
		}
	}
	return grid, Guard{guardPos, Vec2{0, -1}}
}

func stuckInLoop(grid Grid, guard Guard) bool {
	previous := map[Vec2][]Vec2{}

	num := 0
	for {
		num++
		guard = guard.Move(grid)

		if grid[guard.Pos] == 0 {
			return false
		}

		prev, isSet := previous[guard.Pos]

		if isSet && slices.Contains(prev, guard.Dir) {
			num = 0

			grid[guard.Pos] = '^'

			grid[guard.Pos] = '.'

			return true
		}

		if !isSet {
			previous[guard.Pos] = []Vec2{}
		}

		previous[guard.Pos] = append(previous[guard.Pos], guard.Dir)
	}
}

func partTwo(input string) int {
	grid, guard := parseInput(input)

	count := 0

	for location, char := range grid {
		if char == '.' {
			grid[location] = '#'

			stuck := stuckInLoop(grid, guard)

			if stuck {
				count++
			}

			grid[location] = '.'
		}
	}

	return count
}
