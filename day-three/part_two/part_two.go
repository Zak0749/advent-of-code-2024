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

	fmt.Println("Part Two: ", partTwo(string(input)))
}

func partTwo(input string) int {
	r, _ := regexp.Compile(`mul\(([\d]+,[\d]+)\)|do\(\)|don't\(\)`)
	r2, _ := regexp.Compile(`[\d]+,[\d]+`)

	actions := r.FindAllString(input, -1)

	sum := 0

	doInstructions := true

	for _, action := range actions {
		fmt.Println(action)

		if action == "do()" {
			doInstructions = true
		} else if action == "don't()" {
			doInstructions = false
		} else if doInstructions {
			numLine := r2.FindString(action)

			arr := strings.Split(numLine, ",")

			num1, _ := strconv.Atoi(arr[0])
			num2, _ := strconv.Atoi(arr[1])

			sum += num1 * num2
		}
	}

	return sum
}
