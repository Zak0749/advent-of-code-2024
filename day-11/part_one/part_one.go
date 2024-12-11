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

type Stone int

func (s Stone) next() []Stone {
	if s == 0 {
		return []Stone{1}
	}

	str := strconv.Itoa(int(s))

	if len(str)%2 == 0 {
		mid := len(str) / 2

		first, err := strconv.Atoi(str[:mid])

		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(str[mid:])
		if err != nil {
			panic(err)
		}

		return []Stone{Stone(first), Stone(second)}
	} else {
		return []Stone{s * 2024}
	}
}

func parseInput(input string) []Stone {
	var stones []Stone

	trimmed := strings.Trim(input, "\n ")

	for _, s := range strings.Split(trimmed, " ") {
		stone, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		stones = append(stones, Stone(stone))
	}

	return stones
}

func partOne(input string) int {
	stones := parseInput(input)

	for range 25 {
		var newStones []Stone

		for _, s := range stones {
			newStones = append(newStones, s.next()...)
		}

		stones = newStones
	}

	return len(stones)
}
