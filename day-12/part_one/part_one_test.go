package main

import "testing"


func Test1(t *testing.T) {
	output := partOne(`AAAA
BBCD
BBCC
EEEC`)
	if output != 140 {
		t.Fatalf("Expected: %d Got: %d", 140, output)
	}
}

func Test2(t *testing.T) {
	output := partOne(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`)
	if output != 772 {
		t.Fatalf("Expected: %d Got: %d", 772, output)
	}
}

func Test3(t *testing.T) {
	output := partOne(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)
	if output != 1930 {
		t.Fatalf("Expected: %d Got: %d", 1930, output)
	}
}
