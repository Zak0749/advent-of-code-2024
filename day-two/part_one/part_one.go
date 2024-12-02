package main

import (
	"fmt"
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

func is_positive(n int) bool {
	return n > 0
}

func partOne(input string) int {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		nums := strings.Split(line, " ")

		diffs := []int{}

		for i, v := range nums[:len(nums)-1] {
			current, err := strconv.Atoi(v)
			if err != nil {
				panic(strings.Trim(v, "\n "))
			}

			next, err := strconv.Atoi(nums[i+1])
			if err != nil {
				panic(nums[i+1])
			}

			diff := current - next
			diffs = append(diffs, diff)
		}

		positiveRun := is_positive(diffs[0])

		passed := true

		for _, v := range diffs {
			if positiveRun != is_positive(v) || v > 3 || v < -3 || v == 0 {
				passed = false
			}
		}
		if passed {
			count++
		}
	}
	return count
}
