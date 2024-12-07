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

	fmt.Println("Part Two: ", partTwo(string(input)))
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

func getErrors(update []int, ordering map[int][]int) int {
	for i, num := range update {
		ordering_num, is_set := ordering[num]

		if is_set {
			for _, num2 := range update[:i] {
				if slices.Contains(ordering_num, num2) {
					return i
				}
			}
		}
	}

	return -1
}

func fixErrors(update []int, ordering map[int][]int) []int {
	errors := getErrors(update, ordering)
	if errors == -1 {
		return update
	}

	newUpdate := slices.Clone(update)

	val := newUpdate[errors]
	newUpdate = append(newUpdate[:(errors)], newUpdate[errors+1:]...)

	newUpdate = append([]int{val}, newUpdate...)

	return fixErrors(newUpdate, ordering)
}

func partTwo(input string) int {
	sections := strings.Split(input, "\n\n")
	ordering_str, update_str := sections[0], sections[1]

	ordering := parse_ordering(ordering_str)
	updates := parse_updates(update_str)

	middle_nums := []int{}

	for _, update := range updates {

		if getErrors(update, ordering) != -1 {
			newUpdate := fixErrors(update, ordering)
			fmt.Println("newUpdate: ", newUpdate)

			half := len(newUpdate) / 2
			middle_nums = append(middle_nums, newUpdate[half])
		}
	}

	count := 0

	for _, num := range middle_nums {
		count += num
	}

	return count
}
