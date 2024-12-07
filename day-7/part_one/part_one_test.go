package main

import "testing"

const INPUT string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
const EXPECT int = 3749

func Test(t *testing.T) {
	output := partOne(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}
