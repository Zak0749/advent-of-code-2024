package main

import "testing"

const INPUT string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
const EXPECT int = 14

func Test(t *testing.T) {
	output := partOne(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}

func Test2(t *testing.T) {
	output := partOne(`..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........`)
	if output != 2 {
		t.Fatalf("Expected: %d Got: %d", 2, output)
	}
}

func Test3(t *testing.T) {
	output := partOne(`..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`)
	if output != 4 {
		t.Fatalf("Expected: %d Got: %d", 2, output)
	}
}

func Test4(t *testing.T) {
	output := partOne(`..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`)
	if output != 4 {
		t.Fatalf("Expected: %d Got: %d", 2, output)
	}
}
