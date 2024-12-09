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

	fmt.Println("Part Two: ", partTwo(string(input)))
}

type DiskItem struct {
	id  int
	len int
}

func parseFile(input string) []DiskItem {
	out := []DiskItem{}

	for i, c := range input {
		if i%2 == 0 {
			l, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			out = append(out, DiskItem{
				id:  int(i / 2),
				len: l,
			})
		} else {
			l, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			out = append(out, DiskItem{
				id:  -1,
				len: l,
			})
		}
	}

	return out
}

func partTwo(input string) int {
	data := parseFile(input)
	moved := []DiskItem{}

	taken := []int{}

	for _, n := range data {
		if n.id != -1 && !slices.Contains(taken, n.id) {
			moved = append(moved, DiskItem{
				id:  n.id,
				len: n.len,
			})

			taken = append(taken, n.id)
		} else {
			lenToFill := n.len
			for _, m := range slices.Backward(data) {
				if m.id != -1 && !slices.Contains(taken, m.id) && m.len <= lenToFill {

					moved = append(moved, DiskItem{
						id:  m.id,
						len: m.len,
					})

					taken = append(taken, m.id)

					lenToFill -= m.len

					if lenToFill == 0 {
						break
					}
				}
			}

			if lenToFill > 0 {
				moved = append(moved, DiskItem{
					id:  -1,
					len: lenToFill,
				})
			}
		}
	}

	count := 0
	i := 0

	for _, n := range moved {
		if n.id != -1 {
			for range n.len {
				count += i * n.id
				i += 1
			}
		} else {
			i += n.len
		}
	}

	return count
}
