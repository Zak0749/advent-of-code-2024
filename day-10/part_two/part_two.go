package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part Two: ", partTwo(string(input)))
}

type Point struct {
	x, y int
}

func (point Point) add(other Point) Point {
	return Point{point.x + other.x, point.y + other.y}
}

type TopographicMap map[Point]int

func parseInput(input string) TopographicMap {
	topographicMap := map[Point]int{}

	for y, line := range strings.Split(input, "\n") {
		for x, item := range line {
			height, err := strconv.Atoi(string(item))
			if err != nil {
				panic("invalid input")
			}

			topographicMap[Point{x, y}] = height
		}
	}

	return topographicMap
}

type Path struct {
	point Point
	next  []*Path
}

func pathFromPoint(current Point, last []Point, topographicMap TopographicMap) *Path {
	directions := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	if topographicMap[current] == 9 {
		return &Path{current, []*Path{}}
	}

	path := &Path{current, []*Path{}}

	for _, direction := range directions {
		nextPoint := current.add(direction)
		height, isSet := topographicMap[nextPoint]

		if isSet && height == topographicMap[current]+1 {
			path.next = append(path.next, pathFromPoint(nextPoint, append(last, nextPoint), topographicMap))
		}
	}

	return path
}

func countNines(path Path, topographicMap TopographicMap) int {
	if len(path.next) == 0 {
		if topographicMap[path.point] == 9 {
			return 1
		} else {
			return 0
		}
	}

	count := 0
	for _, next := range path.next {
		count += countNines(*next, topographicMap)
	}

	return count
}

func partTwo(input string) int {
	topographicMap := parseInput(input)

	scores := 0

	for point, height := range topographicMap {
		if height == 0 {
			path := pathFromPoint(point, []Point{}, topographicMap)

			scores += countNines(*path, topographicMap)
		}
	}

	return scores
}
