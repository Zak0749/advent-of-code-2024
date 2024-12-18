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

	fmt.Println("Part Two:", partTwo(string(input), Vec2{70, 70}))
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

func makeMap(walls []Vec2, size Vec2) Grid {
	grid := Grid{}

	max := size.i * size.j

	for i := range size.i + 1 {
		for j := range size.j + 1 {
			grid[Vec2{i, j}] = max
		}
	}

	for _, wall := range walls {
		grid[wall] = -1
	}

	grid[Vec2{0, 0}] = 0

	return grid
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

func getWalls(input string) []Vec2 {
	walls := []Vec2{}

	re := regexp.MustCompile(`[\d]+`)

	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllString(line, -1)

		i, err := strconv.Atoi(matches[0])
		if err != nil {
			panic(err)
		}

		j, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		walls = append(walls, Vec2{i, j})
	}

	return walls
}

func partTwo(input string, size Vec2) string {
	walls := getWalls(input)
	for i, wall := range walls {
		grid := makeMap(walls[:(i+1)], size)

		traverseGrid(Vec2{0, 0}, grid)

		if grid[size] == size.i*size.j {
			return fmt.Sprintf("%d,%d", wall.i, wall.j)
		}
	}

	panic("All have solutions")
}
