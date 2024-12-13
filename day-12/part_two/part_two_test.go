package main

import "testing"

const INPUT string = ``
const EXPECT int = 0

func Test1(t *testing.T) {
	output := partTwo(`AAAA
BBCD
BBCC
EEEC`)
	if output != 80 {
		t.Fatalf("Expected: %d Got: %d", 80, output)
	}
}

func Test2(t *testing.T) {
	output := partTwo(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`)
	if output != 436 {
		t.Fatalf("Expected: %d Got: %d", 436, output)
	}
}

func Test3(t *testing.T) {
	output := partTwo(`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`)
	if output != 236 {
		t.Fatalf("Expected: %d Got: %d", 236, output)
	}
}


func Test4(t *testing.T) {
	output := partTwo(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`)
	if output != 368 {
		t.Fatalf("Expected: %d Got: %d", 368, output)
	}
}

