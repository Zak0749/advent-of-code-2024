package main

import (
	"fmt"
	"math"
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

type Line struct {
	TestValue int
	Values    []int
}

type Operation int

const (
	add Operation = 0
	mul Operation = 1
	con Operation = 2
)

func (op Operation) operate(a int, b int) int {
	if op == add {
		return a + b
	} else if op == mul {
		return a * b
	} else {
		v, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		return v
	}
}

func parseInput(input string) []Line {
	split := strings.Split(input, "\n")
	result := []Line{}
	for _, line := range split {
		parts := strings.Split(line, ":")

		testValue, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("bad err")
		}

		values := []int{}

		for _, v := range strings.Split(strings.Trim(parts[1], "\n "), " ") {
			i, err := strconv.Atoi(strings.Trim(v, "\n "))
			if err != nil {
				panic("a bad e val")
			}

			values = append(values, i)
		}

		result = append(result, Line{
			testValue,
			values,
		})
	}

	return result
}

func OpArray(in int, l int) []Operation {
	arr := strconv.FormatInt(int64(in), 3)

	result := make([]Operation, l)
	// result := []Operation{}

	offset := l - len(arr)

	for i, d := range arr {
		if d == '0' {
			result[i + offset] = add
		} else if d == '1' {
			result[i + offset] = mul
		} else if d == '2' {
			result[i + offset] = con
		} else {
			panic("invalid digit")
		}
	}

	return result
}

func PowI(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func partTwo(input string) int {
	lines := parseInput(input)

	sum := 0
	

	for _, line := range lines {
		add := true
		l := PowI(3, len(line.Values)-1)
		for i := range l {
			opArray := OpArray(i, len(line.Values)-1)

			calc := line.Values[0]

			for i, op := range opArray {
				calc = op.operate(calc, line.Values[i+1])
			}

			if calc == line.TestValue && add {
				sum += calc
				add = false
			}
		}

	}

	return sum
}
