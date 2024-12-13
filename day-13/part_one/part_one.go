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

	fmt.Println("Part One:", partOne(string(input)))
}

type Vector struct {
	i, j int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.i + v2.i, v.j + v2.j}
}

func (v Vector) Multiply(scalar int) Vector {
	return Vector{v.i * scalar, v.j * scalar}
}

type Machine struct {
	prize Vector
	a     Vector
	b     Vector
}

func getMachines(input string) []Machine {
	re := regexp.MustCompile(`[\d]+`)
	machines := []Machine{}

	for _, line := range strings.Split(input, "\n\n") {
		matches := re.FindAllString(line, -1)
		nums := []int{}

		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		machines = append(machines, Machine{
			a:     Vector{nums[0], nums[1]},
			b:     Vector{nums[2], nums[3]},
			prize: Vector{nums[4], nums[5]},
		})
	}

	return machines
}

func solveMachine(machine Machine) (int, int, bool) {
	a := float64(machine.a.i)
	b := float64(machine.a.j)
	c := float64(machine.b.i)
	d := float64(machine.b.j)
	p := float64(machine.prize.i)
	q := float64(machine.prize.j)

	det := a * d - b * c

	if det == 0 {
		return 0, 0, true
	}

	x := (d * p - c * q) / det
	y := (a * q - b * p) / det

	if machine.a.i * int(x) + machine.b.i * int(y) == machine.prize.i && machine.a.j * int(x) + machine.b.j * int(y) == machine.prize.j {
		return int(x), int(y), true
	}

	return 0, 0, false
}

func partOne(input string) int {
	machines := getMachines(input)

	count := 0


	for machine := range machines {
		x, y, ok := solveMachine(machines[machine])
		fmt.Println(x, y, ok)
		if ok {
			count += 3 * x + 1 * y
		}
	}

	return count
}
