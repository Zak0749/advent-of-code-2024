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

func partTwo(input string) int {
	locationIDs := []int{}
	frequency := map[int]int{}

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

		locationIDs = append(locationIDs, firstNum)

		frequency[secondNum] += 1
	}

	sum := 0

	for _, id := range locationIDs {
		sum += id * frequency[id]
	}

	return sum
}
