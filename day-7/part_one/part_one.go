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

	fmt.Println("Part One:", partOne(string(input)))
}

type Line struct {
	TestValue int
	Values    []int
}

type Operation int

const (
	add Operation = 0
	mul Operation = 1
)

func (op Operation) operate(a int, b int) int {
	if op == add {
		return a + b
	} else {
		return a * b
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

func OpArray(in int, op_num int) []Operation {
	boolArray := make([]Operation, op_num)

	// Iterate through each bit position
	for i := 0; i < op_num; i++ {
		// Check if the bit at position `i` is set
		if (in & (1 << i)) != 0 {
			boolArray[op_num-1-i] = 1
		} else {
			boolArray[op_num-1-i] = 0
		}
	}

	return boolArray
}

func PowI(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func partOne(input string) int {
	lines := parseInput(input)

	sum := 0

	for _, line := range lines {
		l := PowI(2, len(line.Values)-1) - 1
		for i := range l + 1 {
			opArray := OpArray(i, len(line.Values) - 1)
			
			calc := line.Values[0]
			
			for i, op := range opArray {
				calc = op.operate(calc, line.Values[i + 1])
			}

			if calc == line.TestValue {
				sum += calc
				break
			}
		}

	}

	return sum
}
