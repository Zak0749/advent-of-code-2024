package main

import (
	"fmt"
	"math"
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

	fmt.Println("Part Two: ", partTwo(string(input)))
}

func parseInput(input string) []int {
	stones := []int{}

	trimmed := strings.Trim(input, "\n ")

	for _, s := range strings.Split(trimmed, " ") {
		stone, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		stones = append(stones, stone)
	}

	return stones
}

func int_to_digits(n int) []int {
	digits := []int{}

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	slices.Reverse(digits)

	return digits
}

func digits_to_int(digits []int) int {
	n := 0

	for i, d := range digits {
		n += d * int(math.Pow(10, float64(len(digits)-i-1)))
	}

	return n
}

func blink(n int) []int {
	if n == 0 {
		return []int{1}
	}

	digits := int_to_digits(n)

	if len(digits)%2 == 0 {
		mid := len(digits) / 2

		first := digits_to_int(digits[:mid])
		second := digits_to_int(digits[mid:])

		return []int{first, second}
	}

	res := []int{n * 2024}

	return res
}

type Previous struct {
	stone int
	n     int
}

type PreviousMap map[Previous]int

func treeLength(stone int, n int, p PreviousMap) int {
	if n == 0 {
		return 1
	}

	if value, ok := p[Previous{stone, n}]; ok {
		return value
	}

	blinks := blink(stone)


	count := 0

	for _, b := range blinks {
		num := treeLength(b, n-1, p)
		count += num
	}

	p[Previous{stone, n}] = count

	return count
}

func partTwo(input string) int {
	stones := parseInput(input)

	previous := PreviousMap{}

	total := 0

	for _, stone := range stones {
		total += treeLength(stone, 75, previous)
	}

	return total
}
