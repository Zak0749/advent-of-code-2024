package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", partOne(string(input), Vec2{70, 70}, 1024))
}

type Vec2 struct {
	i, j int
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.i + v2.i, v.j + v2.j}
}

func (v Vec2) neighbors() []Vec2 {
	return []Vec2{
		v.Add(Vec2{0, 1}),
		v.Add(Vec2{0, -1}),
		v.Add(Vec2{1, 0}),
		v.Add(Vec2{-1, 0}),
	}
}

type Grid map[Vec2]int

func makeMap(input string, size Vec2, n int) Grid {
	out := Grid{}

	for i := range size.i + 1 {
		for j := range size.j + 1 {
			out[Vec2{i, j}] = size.i * size.j
		}
	}

	out[Vec2{0, 0}] = 0

	re := regexp.MustCompile(`[\d]+`)

	for _, line := range strings.Split(input, "\n")[:n] {
		matches := re.FindAllString(line, -1)

		i, err := strconv.Atoi(matches[0])
		if err != nil {
			panic(err)
		}

		j, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		out[Vec2{i, j}] = -1
	}

	return out
}

func traverseGrid(current Vec2, grid Grid) {
	for _, neighbor := range current.neighbors() {
		space, isset := grid[neighbor]
		if isset && space != -1 && grid[current]+1 < space {
			grid[neighbor] = grid[current] + 1
			traverseGrid(neighbor, grid)
		}
	}
}


func partOne(input string, size Vec2, n int) int {
	grid := makeMap(input, size, n)

	traverseGrid(Vec2{0, 0}, grid)

	return grid[size]
}
