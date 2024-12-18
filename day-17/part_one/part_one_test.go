package main

import (
	"testing"
)

const INPUT string = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
const EXPECT string = "4,6,3,5,6,3,5,2,1,0"

func Test(t *testing.T) {
	output := partOne(INPUT)
	if output != EXPECT {
		t.Fatalf("Expected: %s Got: %s", EXPECT, output)
	}
}

func Test1(t *testing.T) {
	program := Instruction{opcode: Opcode(2), combo: Combo(6)}
	state := State{a: 0, b: 0, c: 9, instructionNum: 0, output: []string{}}
	program.opcode.preform(program.combo, &state)
	if state.b != 1 {
		t.Fatalf("Expected: 1 Got: %d", state.b)
	}
}



func Test2(t *testing.T) {
	program := []Instruction{{Opcode(5), Combo(0)}, {Opcode(5), Combo(1)}, {Opcode(5), Combo(4)}}
	state := State{a: 10, b: 0, c: 0, instructionNum: 0, output: []string{}}

	for _, p := range program {
		p.opcode.preform(p.combo, &state)
	}

	if !equalSlices(state.output, []string{"0", "1", "2"}) {
		t.Fatalf("Expected: [0 1 2] Got: %v", state.output)
	}
}

func Test3(t *testing.T) {
	program := []Instruction{{Opcode(0), Combo(1)}, {Opcode(5), Combo(4)}, {Opcode(3), Combo(0)}}
	state := State{a: 2024, b: 0, c: 0, instructionNum: 0, output: []string{}}

	for {
		if state.instructionNum >= len(program) {
			break
		}

		instruction := program[state.instructionNum]

		instruction.opcode.preform(instruction.combo, &state)
	}

	// for _, p := range program {
	// 	state = p.opcode.preform(p.combo, state)
	// 	fmt.Printf("%#v\n", state)
	// }

	if !equalSlices(state.output, []string{"4", "2", "5", "6", "7", "7", "7", "7", "3", "1", "0"}) {
		t.Fatalf("Expected: [4 2 5 6 7 7 7 7 3 1 0] Got: %v", state.output)
	}

	if state.a != 0 {
		t.Fatalf("Expected: 0 Got: %d", state.a)
	}
}

func Test4(t *testing.T) {
	program := []Instruction{{Opcode(1), Combo(7)}}
	state := State{a: 0, b: 29, c: 0, instructionNum: 0, output: []string{}}

	for {
		if state.instructionNum >= len(program) {
			break
		}

		instruction := program[state.instructionNum]

		instruction.opcode.preform(instruction.combo, &state)
	}

	if state.b != 26 {
		t.Fatalf("Expected: 26 Got: %d", state.b)
	}
}

func Test5(t *testing.T) {
	program := []Instruction{{Opcode(4), Combo(0)}}
	state := State{a: 0, b: 2024, c: 43690, instructionNum: 0, output: []string{}}

	for _, p := range program {
		p.opcode.preform(p.combo, &state)
	}

	if state.b != 44354 {
		t.Fatalf("Expected: 44354 Got: %d", state.b)
	}
}
