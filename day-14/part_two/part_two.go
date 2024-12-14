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

	partTwo(string(input), Vec{101, 103})
}

func partTwo(input string, size Vec) int {
	robots := getRobots(input)

	times := 10000

	max := 0.

	for t := range times {
		sum := Vec{0, 0}
		for i, robot := range robots {
			newPos := robot.pos.Add(robot.vel).Mod(size)
			robots[i].pos = newPos
			sum = sum.Add(newPos)
		}

		dx := float64(sum.i) / float64(len(robots))
		dy := float64(sum.j) / float64(len(robots))

		avg := dx*dx + dy*dy

		if avg > max {
			max = avg
			fmt.Println(t+1, avg)
			printRobotsOnGrid(robots, size)
		}
	}

	fmt.Println(max)

	return 0
}

type Vec struct {
	i, j int
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.i + v2.i, v.j + v2.j}
}

func (v Vec) Mul(n int) Vec {
	return Vec{v.i * n, v.j * n}
}

func (v Vec) Div(n int) Vec {
	return Vec{v.i * n, v.j * n}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (v Vec) Mod(v2 Vec) Vec {
	return Vec{mod(v.i, v2.i), mod(v.j, v2.j)}
}

type Robot struct {
	pos Vec
	vel Vec
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
			show := false
			for _, robot := range robots {
				if robot.pos.i == i && robot.pos.j == j {
					show = true
				}
			}

			if !show {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
}
