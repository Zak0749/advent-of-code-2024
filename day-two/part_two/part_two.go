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

	fmt.Println("Part Two: ", partTwo(string(input)))
}

func is_positive(n int) bool {
	return n > 0
}

func to_int_arr(s []string) []int {
	ints := []int{}
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(strings.Trim(v, "\n "))
		}
		ints = append(ints, i)
	}
	return ints
}

func find_sign(arr []int) bool {
	posNum := 0
	negNum := 0

	for _, v := range arr {
		if v > 0 {
			posNum++
		} else if v < 0 {
			negNum++
		}
	}

	return posNum > negNum
}

func get_diffs(nums []int) []int {
	diffs := []int{}

	for i, v := range nums[:len(nums)-1] {
		diff := v - nums[i+1]
		diffs = append(diffs, diff)
	}

	return diffs
}

func is_valid(diffs []int) bool {
	positiveRun := find_sign(diffs)
	passed := true

	for _, v := range diffs {
		if positiveRun != is_positive(v) || v > 3 || v < -3 || v == 0 {
			passed = false
		}
	}

	return passed
}

func partTwo(input string) int {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		nums := to_int_arr(strings.Split(line, " "))

		diffs := get_diffs(nums)

		isValid := is_valid(diffs)

		if isValid {
			count++
		} else {
			for i := range nums {
				newNums := make([]int, len(nums))
				copy(newNums, nums)
				newNums = append(newNums[:i], newNums[i+1:]...)
				newDiffs := get_diffs(newNums)
				newIsValid := is_valid(newDiffs)
				if newIsValid {
					count++
					break
				}
			}
		}
	}
	return count
}
