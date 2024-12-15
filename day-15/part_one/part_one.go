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

type Vec struct {
	i, j int
}

func (v Vec) add(o Vec) Vec {
	return Vec{v.i + o.i, v.j + o.j}
}

func (v Vec) sub(o Vec) Vec {
	return Vec{v.i - o.i, v.j - o.j}
}

func (v Vec) equals(o Vec) bool {
	return v.i == o.i && v.j == o.j
}

type Location rune

const (
	NX		Location = 0
	WALL  Location = '#'
	OPEN  Location = '.'
	BOX   Location = 'O'
	ROBOT Location = '@'
)

func toLocation(char rune) Location {
	switch char {
	case '#':
		return WALL
	case '.':
		return OPEN
	case 'O':
		return BOX
	case '@':
		return ROBOT
	}

	panic("Invalid Location")
}

func constructMap(input string) (map[Vec]Location, Vec) {
	out := map[Vec]Location{}
	robot := Vec{}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.TrimSpace(line) {
			out[Vec{x, y}] = toLocation(char)

			if char == '@' {
				robot = Vec{x, y}
			}
		}
	}
	return out, robot
}

func parseDirections(input string) []Vec {
	out := []Vec{}

	for _, char := range strings.TrimSpace(input) {
		switch char {
		case 'v':
			out = append(out, Vec{0, 1})
		case '^':
			out = append(out, Vec{0, -1})
		case '>':
			out = append(out, Vec{1, 0})
		case '<':
			out = append(out, Vec{-1, 0})
		}
	}

	return out
}

func printWarehouse(warehouse map[Vec]Location) {
	j := 0
	for {
		i := 0

		for {
			if loc, ok := warehouse[Vec{i, j}]; ok {
				fmt.Print(string(loc))
			} else {
				break
			}

			i++
		}

		if i == 0 {
			break
		}

		fmt.Println()

		j++
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func parseInput(input string) (map[Vec]Location, Vec, []Vec) {
	parts := strings.Split(input, "\n\n")
	warehouse, robot := constructMap(parts[0])
	instructions := parseDirections(parts[1])

	return warehouse, robot, instructions
}

func makeMove(warehouse map[Vec]Location, robot Vec, direction Vec) (map[Vec]Location, Vec) {
	newRobot := robot.add(direction)

	if warehouse[newRobot] == BOX {
		moveTo := Vec{newRobot.i, newRobot.j}

		for warehouse[moveTo] == BOX {
			moveTo = moveTo.add(direction)
		}

		if warehouse[moveTo] == OPEN {
			swap := Vec{moveTo.i, moveTo.j}
			for !swap.equals(robot) {
				warehouse[swap] = warehouse[swap.sub(direction)]
				swap = swap.sub(direction)
			}

			warehouse[robot] = OPEN

			return warehouse, newRobot
		}
	}

	if warehouse[newRobot] == OPEN {
		warehouse[robot] = OPEN
		warehouse[newRobot] = ROBOT
		return warehouse, newRobot
	}

	return warehouse, robot
}

func partOne(input string) int {
	warehouse, robot, instructions := parseInput(input)

	fmt.Println("Initial Warehouse")
	printWarehouse(warehouse)


	for _, instruction := range instructions {
		warehouse, robot = makeMove(warehouse, robot, instruction)
	}

	count := 0 

	for pos, loc := range warehouse {
		if loc == BOX {
			count += 100 * pos.j + pos.i
		}
	}

	return count
}
