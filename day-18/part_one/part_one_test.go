package main

import "testing"

const INPUT string = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
const EXPECT int = 22

func Test(t *testing.T) {
	output := partOne(INPUT, Vec2{6, 6}, 12)
	if output != EXPECT {
		t.Fatalf("Expected: %d Got: %d", EXPECT, output)
	}
}
