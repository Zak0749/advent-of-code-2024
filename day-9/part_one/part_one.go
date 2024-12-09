package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", partOne(string(input)))
}

func parseFile(input string) []int {
	out := []int{}

	for i, c := range input {
		if i%2 == 0 {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			for range num {
				out = append(out, i/2)
			}
		} else {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			for range num {
				out = append(out, -1)
			}
		}
	}

	return out
}


func partOne(input string) int {
	data := parseFile(input)

	moved := make([]int, len(data))

	for i := range moved {
		moved[i] = -1
	}

	emptiedTo := len(data)

	for i, n := range data {
		if n == -1 {
			c := 1
			for _, m := range slices.Backward(data[i:(emptiedTo)]) {
				if m != -1 {
					emptiedTo -= c
					moved[i] = m
					break
				}
				c++
			}			
		} else {
			moved[i] = n
		}

		if emptiedTo - 1 == i {
			break
		}
	}


	count := 0

	for i, n := range moved {
		if n != -1 {
			count += i * n
		}
	}


	return count
}
