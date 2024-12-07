package main

import "testing"

const INPUT string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
const EXPECT int = 18

func Test(t *testing.T) {
	output := partOne(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}

func Test2(t *testing.T) {
	output := partOne(`..X...
.SAMX.
.A..A.
XMAS.S
.X....`)
	if output != 4 {
		t.Fatalf("Expected: %d Got: %d", 4, output)
	}
}
