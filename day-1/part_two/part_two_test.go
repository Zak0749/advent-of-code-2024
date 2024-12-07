package main

import "testing"

const INPUT string = `3   4
4   3
2   5
1   3
3   9
3   3`
const EXPECT int = 31

func Test(t *testing.T) {
	output := partTwo(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}
