package main

import "testing"

const INPUT string = `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`
const EXPECT int = 9

func Test(t *testing.T) {
	output := partTwo(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}

func Test2(t *testing.T) {
	output := partTwo(`M.S
.A.
M.S`)

	if output != 1 {
		t.Fatalf("Expected: %d Got: %d", 1, output)
	}
}