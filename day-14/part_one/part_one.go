package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Vec struct {
	i, j int
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.i + v2.i, v.j + v2.j}
}

func (v Vec) Mul(n int) Vec {
	return Vec{v.i * n, v.j * n}
}

// Werid as go does diff to python so impl python now
func mod(a, b int) int {
	return (a % b + b) % b
}

func (v Vec) Mod(v2 Vec) Vec {
	return Vec{mod(v.i, v2.i), mod(v.j, v2.j)}
}

type Robot struct {
	pos Vec
	vel Vec
}

func product[T int | float64](values []T) T {
	var sum T = 1
	for _, v := range values {
		sum *= v
	}
	return sum
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", partOne(string(input), Vec{101, 103}))
}

func getRobots(input string) []Robot {
	robots := []Robot{}
	re := regexp.MustCompile(`(-|)\d+`)

	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllString(line, -1)
		nums := []int{}

		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		pos := Vec{nums[0], nums[1]}
		vel := Vec{nums[2], nums[3]}
		robots = append(robots, Robot{pos, vel})
	}

	return robots
}

func printRobotsOnGrid(robots []Robot, size Vec) {
	for j := range size.j {
		for i := range size.i {
			n := 0
			for _, robot := range robots {
				if robot.pos.i == i && robot.pos.j == j {
					n++
				}
			}

			if n == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(n)
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func partOne(input string, size Vec) int {
	fmt.Println(-10 % 3)
	robots := getRobots(input)

	times := 100

	half := Vec{-size.i / 2, -size.j / 2}

	for range times {
		for i, robot := range robots {
			newPos := robot.pos.Add(robot.vel).Mod(size)
			robots[i].pos = newPos
		}
		printRobotsOnGrid(robots, size)
	}

	quadrants := []int{0, 0, 0, 0}

	for _, robot := range robots {
		quad := robot.pos.Add(half)

		if quad.i < 0 && quad.j < 0 {
			quadrants[0]++
		} else if quad.i < 0 && quad.j > 0 {
			quadrants[1]++
		} else if quad.i > 0 && quad.j < 0 {
			quadrants[2]++
		} else if quad.i > 0 && quad.j > 0 {
			quadrants[3]++
		}
	}

	fmt.Println(quadrants)

	return product(quadrants)
}
