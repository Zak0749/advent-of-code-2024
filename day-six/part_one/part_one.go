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
	// Check if we can move forward
	nextPos := g.Pos.Add(g.Dir)

	if grid[nextPos] == '#' {
		right := Vec2{-g.Dir.J, g.Dir.I}
		g.Dir = right
		nextPos = g.Pos.Add(right)
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

func partOne(input string) int {
	grid, guard := parseInput(input)
	// guardStart := Vec2{guard.Pos.I, guard.Pos.J}

	guard = guard.Move(grid)

	count := 1

	for {
		guard = guard.Move(grid)

		if grid[guard.Pos] == 0 {
			break
		}

		grid[guard.Pos] = 'X'
		fmt.Println(guard.Pos, "next:", string(grid[guard.Pos.Add(guard.Dir)]))
	}

	for _, char := range grid {
		if char == 'X' || char == '^' {
			count++
		}
	}

	return count
}
