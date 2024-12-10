package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", partOne(string(input)))
}

type Point struct {
	x, y int
}

func (point Point) add(other Point) Point {
	return Point{point.x + other.x, point.y + other.y}
}

func (point Point) eq(other Point) bool {
	return point.x == other.x && point.y == other.y
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

func findRecursive(current Point, last []Point, topographicMap TopographicMap) []Point {
	fmt.Println(current, topographicMap[current])
	if topographicMap[current] == 9 {
		return []Point{current}
	}

	directions := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	results := []Point{}

	for _, direction := range directions {
		nextPoint := current.add(direction)
		height, isSet := topographicMap[nextPoint]

		fmt.Printf("nextPoint: %v, height: %v, isSet: %v  %t \n", 
    nextPoint, 
    height, 
    isSet, 
    isSet && height == topographicMap[current]+1 && !slices.ContainsFunc(last, func(e Point) bool { return e.eq(nextPoint) }))

		if isSet && height == topographicMap[current]+1 && !slices.ContainsFunc(last, func(e Point) bool { return e.eq(nextPoint) }) {
			results = append(results, findRecursive(
				nextPoint,
				append(last, current),
				topographicMap,
			)...)
		}
	}

	return results
}

func printMap(topographicMap TopographicMap) {
	y := 0
	for {
		x := 0
		for {
			value, isSet := topographicMap[Point{x, y}]

			if !isSet {
				break
			}

			fmt.Print(value)

			x++
		}
		fmt.Println()
		if x == 0 {
			break
		}
		y++
	}
}

func partOne(input string) int {
	topographicMap := parseInput(input)
	printMap(topographicMap)

  scores := 0


	for point, height := range topographicMap {
		if height == 0 {
	    peakMap := map[Point]bool{}

			for _, peak := range findRecursive(point, []Point{}, topographicMap) {
				peakMap[peak] = true
				fmt.Println(peak)
			}

      for _, peak := range peakMap {
        if peak {
          scores += 1
        }
      }
		}
	}

	return scores
}
