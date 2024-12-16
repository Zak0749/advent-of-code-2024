package main

import (
	"fmt"
	"math"
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

func cloneMap[T comparable, V any](mp map[T]V) map[T]V {
	clone := map[T]V{}
	for k, v := range mp {
		clone[k] = v
	}
	return clone
}

type Tile struct {
	value int
	object Object
	isEmpty bool
}

type Object rune

const (
	WALL  Object = '#'
	EMPTY Object = '.'
	START Object = 'S'
	END   Object = 'E'
	HIT   Object = 'X'
)

type Point struct {
	x, y int
}

func (v Point) add(o Point) Point {
	return Point{v.x + o.x, v.y + o.y}
}

func (v Point) equal(o Point) bool {
	return v.x == o.x && v.y == o.y
}

func (v Point) directions() []Point {
	return []Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
}

type Maze map[Point]Tile

func parseMaze(input string) (Maze, Point, Point) {
	maze := Maze{}
	start := Point{0, 0}
	end := Point{0, 0}

	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == '.' {
				maze[Point{x, y}] = Tile { math.MaxInt32, 0, true }
			} else {
				maze[Point{x, y}] = Tile { -1, Object(char), false }
			}

			if char == 'S' {
				start = Point{x, y}
			}

			if char == 'E' {
				end = Point{x, y}
			}
		}
	}

	return maze, start, end
}

type Path struct {
	score 	 int
	point     Point
	direction Point
	previous  *Path
	next      []*Path
}

func constructPath(last *Path, maze Maze) []int {
	scores := []int{}
	for _, direction := range last.point.directions() {
		scoreGain := 1

		if !direction.equal(last.direction) {
			scoreGain += 1000
		}

		point := last.point.add(direction)
		if maze[point].object == START || maze[point].object == WALL {
			continue
		}

		if maze[point].isEmpty && maze[point].value < last.score + scoreGain {
			continue
		}

		if maze[point].object == END {
			last.next = append(last.next, &Path{last.score + scoreGain, point, direction, last, nil})
			scores = append(scores, last.score + scoreGain)
		}

		maze[point] = Tile{
			last.score + scoreGain,
			0,
			true,
		}

		newPath := &Path{last.score + scoreGain, point, direction, last, nil}

		// printMaze(maze)

		scores = append(scores, constructPath(newPath, maze)...)

		last.next = append(last.next, newPath)
	}

	return scores
}

func printMaze(maze Maze) {
	y:=0
	for {
		x := 0
		for {
			if _, exists := maze[Point{x, y}]; !exists {
				break
			}

			val := maze[Point{x, y}]

			if val.isEmpty {
				fmt.Printf("%5d", val.value)
				} else {
				fmt.Printf("%5s", string(val.object))
			}
			x++
		}

		if x == 0 {
			break
		}

		fmt.Println()

		y++
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

}


func partOne(input string) int {
	maze, start, end := parseMaze(input)

	fmt.Println(end)

	printMaze(maze)

	startingPath := Path{0, start, Point{1, 0}, nil, nil}

	scores := constructPath(&startingPath, maze)

	fmt.Println(scores)

	printMaze(maze)



	return maze[end].value
}
