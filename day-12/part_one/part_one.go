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

	fmt.Println("Part One:", partOne(string(input)))
}

type Vec struct {
	x, y int
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.x + v2.x, v.y + v2.y}
}

func (v Vec) Equals(v2 Vec) bool {
	return v.x == v2.x && v.y == v2.y
}

func (v Vec) IsAdjacent(v2 Vec) bool {
	return v.x == v2.x && (v.y == v2.y+1 || v.y == v2.y-1) || v.y == v2.y && (v.x == v2.x+1 || v.x == v2.x-1)
}

func (v Vec) Adjacent() []Vec {
	return []Vec{
		v.Add(Vec{0, 1}),
		v.Add(Vec{0, -1}),
		v.Add(Vec{1, 0}),
		v.Add(Vec{-1, 0}),
	}
}

type PlantMap map[Vec]bool

type GardenMap map[rune]PlantMap

func getGardenMap(input string) GardenMap {
	gardenMap := GardenMap{}
	for y, row := range strings.Split(input, "\n") {
		for x, cell := range row {
			if _, ok := gardenMap[cell]; !ok {
				gardenMap[cell] = map[Vec]bool{}
			}

			gardenMap[cell][Vec{x, y}] = true
		}
	}
	return gardenMap
}

func branchConnect(current Vec, previous *[]Vec, gardenMap map[Vec]bool) map[Vec]bool {
	result := map[Vec]bool{current: true}

	*previous = append(*previous, current)

	for _, next := range current.Adjacent() {
		value, ok := gardenMap[next]
		if value && ok && !slices.Contains(*previous, next) {
			for k := range branchConnect(next, previous, gardenMap) {
				result[k] = true
			}
		}
	}

	return result
}

func findConnecting(gardenMap GardenMap) map[rune][]map[Vec]bool {
	connections := map[rune][]map[Vec]bool{}

	for c, garden := range gardenMap {
	mainLoop:
		for point := range garden {
			l, isset := connections[c]

			if !isset {
				prev := []Vec{}
				connections[c] = []map[Vec]bool{
					branchConnect(point, &prev, garden),
				}

				continue
			}

			for _, area := range l {
				if area[point] {
					continue mainLoop
				}
			}

			prev := []Vec{}

			connections[c] = append(connections[c], branchConnect(point, &prev, garden))
		}
	}

	return connections
}

func partOne(input string) int {
	gardenMap := getGardenMap(input)
	count := 0

	connections := findConnecting(gardenMap)

	for _, areas := range connections {
		for _, area := range areas {
			perimeter := []Vec{}

			for point := range area {
				for _, adj := range point.Adjacent() {
					if !area[adj] {
						perimeter = append(perimeter, adj)
					}
				}
			}

			count += len(perimeter) * len(area)
		}
	}

	return count
}
