package main

import "testing"

const INPUT string = `2333133121414131402`
const EXPECT int = 1928

func Test(t *testing.T) {
	output := partOne(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}
