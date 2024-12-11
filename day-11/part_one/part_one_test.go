package main

import "testing"

const INPUT string = `125 17`
const EXPECT int = 55312

func Test(t *testing.T) {
	output := partOne(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}
