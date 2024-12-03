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

func partOne(input string) int {
	r, _ := regexp.Compile(`mul\(([\d]+,[\d]+)\)`)
	r2, _ := regexp.Compile(`[\d]+,[\d]+`)

	matches := r.FindAllString(input, -1)

	sum := 0

	for _, match := range matches {
		numLine := r2.FindString(match)

		arr := strings.Split(numLine, ",")

		num1, _ := strconv.Atoi(arr[0])
		num2, _ := strconv.Atoi(arr[1])

		fmt.Println(num1, num2)

		sum += num1 * num2
	}

	return sum
}
