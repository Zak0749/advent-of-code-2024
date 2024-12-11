package main

// tis be trash
// should have just used a hash map 

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part Two: ", partTwo(string(input)))
}

func parseInput(input string) Stone {
	stones := []int{}

	trimmed := strings.Trim(input, "\n ")

	for _, s := range strings.Split(trimmed, " ") {
		stone, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		stones = append(stones, stone)
	}

	firstStone := Stone{
		stones[0],
		nil,
	}

	currentStone := &firstStone

	for _, stone := range stones[1:] {
		currentStone.next = &Stone{
			stone,
			nil,
		}

		currentStone = currentStone.next
	}

	return firstStone
}

type Stone struct {
	num  int
	next *Stone
}

func int_to_digits(n int) []int {
	digits := []int{}

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	slices.Reverse(digits)

	return digits
}

func digits_to_int(digits []int) int {
	n := 0

	for i, d := range digits {
		n += d * int(math.Pow(10, float64(len(digits)-i-1)))
	}

	return n
}

func (s *Stone) length() int {
	n := 0
	for s != nil {
		n++
		s = s.next
	}

	return n
}

func (s *Stone) blink(head *Stone, n int, out chan<- int, wg *sync.WaitGroup, chain_count int) {
	if n == 0 {
		out <- head.length()
		wg.Done()
		return
	}

	if s == nil {
		go head.blink(head, n-1, out, wg, 0)
		return
	}

	if chain_count > 1000 {
		wg.Add(1)
		go s.next.blink(s.next, n, out, wg, 0)
		s.next = nil
	}

	if s.num == 0 {
		s.num = 1

		s.next.blink(head, n, out, wg, chain_count+1)

		return
	}

	digits := int_to_digits(s.num)

	if len(digits)%2 == 0 {
		mid := len(digits) / 2

		endRef := s.next

		s.next = &Stone{
			num:  digits_to_int(digits[mid:]),
			next: endRef,
		}

		s.num = digits_to_int(digits[:mid])

		s.next.next.blink(head, n, out, wg, chain_count+1)

		return
	}

	s.num *= 2024
	s.next.blink(head, n, out, wg, chain_count+1)
}

func partTwo(input string) int {
	stones := parseInput(input)

	wg := &sync.WaitGroup{}
	out := make(chan int)

	wg.Add(1)

	stones.blink(&stones, 75, out, wg, 0)

	go func() {
		wg.Wait()
		close(out)
	}()

	count := 0
	for n := range out {
		count += n
	}

	return count
}
