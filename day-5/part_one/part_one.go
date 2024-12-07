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

func parse_ordering(ordering_str string) map[int][]int {
	ordering := map[int][]int{}

	for _, line := range strings.Split(ordering_str, "\n") {
		info := strings.Split(line, "|")

		key, err := strconv.Atoi(info[0])
		if err != nil {
			panic(err)
		}

		value, err := strconv.Atoi(info[1])
		if err != nil {
			panic(err)
		}

		_, isset := ordering[key]

		if isset {
			ordering[key] = append(ordering[key], value)
		} else {
			ordering[key] = []int{value}
		}

	}

	return ordering
}

func parse_updates(update_str string) [][]int {
	updates := [][]int{}

	for _, line := range strings.Split(update_str, "\n") {
		update := []int{}

		for _, num := range strings.Split(line, ",") {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			update = append(update, val)
		}

		updates = append(updates, update)
	}

	return updates
}

func partOne(input string) int {
	sections := strings.Split(input, "\n\n")
	ordering_str, update_str := sections[0], sections[1]

	ordering := parse_ordering(ordering_str)
	updates := parse_updates(update_str)

	middle_nums := []int{}

	for _, update := range updates {
		fmt.Println("update: ", update)
		hasError := false

		for i, num := range update {
			ordering_num, is_set := ordering[num]

			fmt.Printf("num: %d, ordering_num: %d, is_set: %t\n", num, ordering_num, is_set)

			if is_set {
				for _, num2 := range update[:i] {
					if slices.Contains(ordering_num, num2) {
						fmt.Printf("containg %d error \n", num2)
						hasError = true
					}
				}
			}
		}

		if !hasError {
			half := len(update) / 2
			middle_nums = append(middle_nums, update[half])
		}
	}

	count := 0

	for _, num := range middle_nums {
		count += num
	}

	return count
}
