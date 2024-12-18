package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", partOne(string(input)))
}

type Opcode int

const (
	ADV Opcode = 0
	BXL Opcode = 1
	BST Opcode = 2
	JNZ Opcode = 3
	BXC Opcode = 4
	OUT Opcode = 5
	BDV Opcode = 6
	CDV Opcode = 7
)

func (o Opcode) preform(combo Combo, state *State) {
	switch o {
	case ADV:
		n := combo.value(state)
		state.a = int(math.Trunc(float64(state.a) / math.Pow(2, float64(n))))
		state.instructionNum += 1
	case BXL:
		state.b = int(combo) ^ state.b
		state.instructionNum += 1
	case BST:
		n := combo.value(state)
		state.b = n % 8
		state.instructionNum += 1
	case JNZ:
		fmt.Println(state.a == 0, state.a, combo, int(combo)/2)
		if state.a == 0 {
			state.instructionNum += 1
		} else {
			state.instructionNum = (int(combo) / 2)
		}
	case BXC:
		state.b = state.c ^ state.b
		state.instructionNum += 1
	case OUT:
		n := combo.value(state)
		state.output = append(state.output, strconv.Itoa(n%8))
		state.instructionNum += 1
	case BDV:
		n := combo.value(state)
		state.b = int(math.Trunc(float64(state.a) / math.Pow(2, float64(n))))
		state.instructionNum += 1
	case CDV:
		n := combo.value(state)
		state.c = int(math.Trunc(float64(state.a) / math.Pow(2, float64(n))))
		state.instructionNum += 1
	}
}

type Combo int

const (
	ZERO  Combo = 0
	ONE   Combo = 1
	TWO   Combo = 2
	THREE Combo = 3
	A     Combo = 4
	B     Combo = 5
	C     Combo = 6
)

func (c Combo) value(s *State) int {
	switch c {
	case ZERO:
		return 0
	case ONE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case A:
		return s.a
	case B:
		return s.b
	case C:
		return s.c
	default:
		panic("Invalid combo")
	}
}

type State struct {
	a              int
	b              int
	c              int
	instructionNum int
	output         []string
}

type Instruction struct {
	opcode Opcode
	combo  Combo
}

func parseInput(input string) (State, []Instruction) {
	re := regexp.MustCompile(`[\d]+`)
	split := strings.Split(input, "\n\n")

	registers := []int{}
	for _, reg := range re.FindAllString(split[0], -1) {
		digit, _ := strconv.Atoi(reg)
		registers = append(registers, digit)
	}

	state := State{
		registers[0],
		registers[1],
		registers[2],
		0,
		[]string{},
	}

	instructions := []Instruction{}

	for ins := range slices.Chunk(re.FindAllString(split[1], -1), 2) {
		op, _ := strconv.Atoi(ins[0])
		com, _ := strconv.Atoi(ins[1])
		instructions = append(instructions, Instruction{Opcode(op), Combo(com)})
	}

	return state, instructions
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func partOne(input string) string {
	state, instructions := parseInput(input)

	for {
		if state.instructionNum >= len(instructions) {
			break
		}

		fmt.Printf("%#v\n", state)

		instruction := instructions[state.instructionNum]

		instruction.opcode.preform(instruction.combo, &state)
	}

	return strings.Join(state.output, ",")
}
