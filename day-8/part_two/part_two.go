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

	fmt.Println("Part Two: ", partTwo(string(input)))
}

type Vec struct {
	i, j int
}

func (v Vec) to(other Vec) Vec {
	return Vec{other.i - v.i, other.j - v.j}
}

func (v Vec) scalar(scalar int) Vec {
	return Vec{v.i * scalar, v.j * scalar}
}

func (v Vec) add(other Vec) Vec {
	return Vec{v.i + other.i, v.j + other.j}
}

func (v Vec) eq(other Vec) bool {
	return v.i == other.i && v.j == other.j
}

func parseInput(input string) (map[rune]map[Vec]bool, map[Vec]bool) {
	frequencyMaps := map[rune]map[Vec]bool{}

	for i, line := range strings.Split(input, "\n") {
		for j, item := range line {
			if item != '.' {
				_, isSet := frequencyMaps[item]

				if !isSet {
					frequencyMaps[item] = map[Vec]bool{}
				}

				frequencyMaps[item][Vec{i, j}] = true
			}
		}
	}

	antinodeMap := map[Vec]bool{}

	for i, line := range strings.Split(input, "\n") {
		for j := range strings.Trim(line, " ") {
			antinodeMap[Vec{i, j}] = false
		}
	}

	return frequencyMaps, antinodeMap
}

func ModI(a, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}

func simplifyWithCommonDiff(v Vec) Vec {
	result := Vec{v.i, v.j}
	for _, commonFactor := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23} {
		if ModI(result.i, commonFactor) == 0 && ModI(result.j, commonFactor) == 0 {
			result.i /= commonFactor
			result.j /= commonFactor
		}
	}
	return result
}

func getMultiples(v Vec) []Vec {
	result := []Vec{}

	for m := range 100 {

		result = append(result, v.scalar(m - 50))
	}

	return result
}

func printBoolMap(mp map[Vec]bool) {
	i := 0
	for {
		setOne := false
		j := 0
		for {
			v, isSet := mp[Vec{i, j}]

			if isSet {
				setOne = true
				if v {
					fmt.Print("X")
				} else {
					fmt.Print(".")
				}
			} else {
				break
			}

			j++
		}
		fmt.Println()

		if !setOne {
			break
		}
		i++
	}
}

func partTwo(input string) int {
	frequencyMaps, antinodeMap := parseInput(input)

	for _, nodes := range frequencyMaps {
		for pointA := range nodes {
			for pointB := range nodes {
				diff := simplifyWithCommonDiff(pointA.to(pointB))

				multiples := getMultiples(diff)

				for _, multipleA := range multiples {
					for _, multipleB := range multiples {
						point := pointA.add(multipleA)

						if point.eq(pointB.add(multipleB)) {
							_, isSet := antinodeMap[point]

							if isSet {
								antinodeMap[point] = true
							}
						}
					}
				}
			}
		}
	}

	printBoolMap(antinodeMap)

	count := 0
	for _, v := range antinodeMap {
		if v {
			count++
		}
	}
	return count
}
