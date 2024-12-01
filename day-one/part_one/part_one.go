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

func partOne(input string) int {
	listOne := []int{}
	listTwo := []int{}

	for _, line := range strings.Split(input, "\n") {
		nums := strings.Split(line, "   ")

		firstNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic("num one not valid int")
		}

		secondNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic("num two not valid int")
		}

		listOne = append(listOne, firstNum)
		listTwo = append(listTwo, secondNum)
	}

	slices.Sort(listOne)
	slices.Sort(listTwo)

	totalDifference := 0

	for i := range listOne {
		difference := listTwo[i] - listOne[i]

		if difference < 0 {
			totalDifference += -difference
		} else {
			totalDifference += difference
		}
	}

	return totalDifference
}
