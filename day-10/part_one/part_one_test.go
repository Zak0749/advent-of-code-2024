package main

import "testing"

const INPUT string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
const EXPECT int = 36

func Test(t *testing.T) {
	output := partOne(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}
