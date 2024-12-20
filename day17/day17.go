package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type puzzle struct {
	a, b, c    int
	program    []int
	programStr string
}

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

//go:embed input.txt
var input string

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func mustInts(s string) []int {
	var ints []int
	for _, s := range strings.Split(s, ",") {
		ints = append(ints, mustInt(s))
	}
	return ints
}

func intsToString(ints []int) string {
	var s []string
	for _, i := range ints {
		s = append(s, strconv.Itoa(i))
	}
	return strings.Join(s, ",")
}

func readInput(s string) *puzzle {
	lines := strings.Split(s, "\n")
	a := mustInt(strings.TrimPrefix(lines[0], "Register A: "))
	b := mustInt(strings.TrimPrefix(lines[1], "Register B: "))
	c := mustInt(strings.TrimPrefix(lines[2], "Register C: "))
	programStr := strings.TrimPrefix(lines[4], "Program: ")
	program := mustInts(programStr)
	return &puzzle{a, b, c, program, programStr}
}

func (p puzzle) comboOperand(operand int, a, b, c int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	}
	panic("unknown operand")
}

func (p puzzle) part1() string {
	a, b, c := p.a, p.b, p.c
	instrPtr := 0
	output := []int{}
	for instrPtr < len(p.program) {
		operand := p.program[instrPtr+1]
		switch p.program[instrPtr] {
		case adv:
			numerator := a
			denominator := int(math.Pow(2, float64(p.comboOperand(operand, a, b, c))))
			a = numerator / denominator
		case bxl:
			b = b ^ operand
		case bst:
			b = p.comboOperand(operand, a, b, c) % 8
		case jnz:
			if a != 0 {
				instrPtr = operand
				continue
			}
		case bxc:
			b = b ^ c
		case out:
			output = append(output, p.comboOperand(operand, a, b, c)%8)
		case bdv:
			numerator := a
			denominator := int(math.Pow(2, float64(p.comboOperand(operand, a, b, c))))
			b = numerator / denominator
		case cdv:
			numerator := a
			denominator := int(math.Pow(2, float64(p.comboOperand(operand, a, b, c))))
			c = numerator / denominator
		}
		instrPtr += 2
	}
	return intsToString(output)
}

func (p puzzle) part2() int {
	fmt.Printf("want length of %d\n", len(p.programStr))
	for i,j := 3, int64(8); i <= 31; i += 2 {
		fmt.Printf("%d=%d\n", i, j)
		j = j * j
	}
	lengths := make(map[int]int)
	for a := 2147483648; a < math.MaxInt; a++ {
		p.a = a
		output := p.part1()
		if _, ok := lengths[len(output)]; !ok {
			lengths[len(output)] = a
			fmt.Println(lengths)
		}
		if output == p.programStr {
			return a
		}
	}
	panic("no more numbers")
}

func main() {
	// pexample := readInput(example)
	// pexample2 := readInput(example2)
	pinput := readInput(input)

	// fmt.Printf("example: part1: %s\n", pexample.part1())
	// fmt.Printf("input: part1: %s\n", pinput.part1())
	// fmt.Printf("example2: part2: %d\n", pexample2.part2())
	fmt.Printf("input: part2: %d\n", pinput.part2())

}
