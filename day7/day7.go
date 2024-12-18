package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	testvalue int
	numbers   []int
}

func main() {
	process("sample.txt")
	process("input.txt")
}

func process(filename string) {
	inputs := readInput(filename)
	fmt.Printf("%s part1: %d\n", filename, part1(inputs))
	fmt.Printf("%s part2: %d\n", filename, part2(inputs))
}

func part1(inputs []input) (result int) {
	for _, inp := range inputs {
		if isPossible(inp.testvalue, inp.numbers) {
			result += inp.testvalue
		}
	}
	return
}

func part2(inputs []input) (result int) {
	for _, inp := range inputs {
		if isPossiblePart2(inp.testvalue, 0, inp.numbers) {
			result += inp.testvalue
			fmt.Println(inp.testvalue)
		}
	}
	return
}

func isPossiblePart2(testvalue int, inc int, vals []int) bool {
	if len(vals) == 0 {
		return testvalue == inc
	}

	if len(vals) == 1 {
		if testvalue == inc*vals[0] {
			return true
		}
		if testvalue == inc+vals[0] {
			return true
		}
		if testvalue == mustInt(strconv.Itoa(inc)+strconv.Itoa(vals[0])) {
			return true
		}
	}

	if isPossiblePart2(testvalue, inc*vals[0], vals[1:]) {
		return true
	}

	if isPossiblePart2(testvalue, inc+vals[0], vals[1:]) {
		return true
	}

	if isPossiblePart2(testvalue, mustInt(strconv.Itoa(inc)+strconv.Itoa(vals[0])), vals[1:]) {
		return true
	}

	return false
}

func isPossible(testvalue int, vals []int) bool {
	if len(vals) == 1 {
		if testvalue == vals[0] {
			return true
		}
		return false
	}

	multvals := []int{vals[0] * vals[1]}
	multvals = append(multvals, vals[2:]...)
	if isPossible(testvalue, multvals) {
		return true
	}

	plusvals := []int{vals[0] + vals[1]}
	plusvals = append(plusvals, vals[2:]...)
	if isPossible(testvalue, plusvals) {
		return true
	}

	return false
}

func readInput(filename string) (inputs []input) {
	lines := mustRead(filename)
	for _, line := range lines {
		flds := strings.Fields(line)
		testflds := strings.Split(flds[0], ":")
		testvalue := mustInt(testflds[0])
		var numbers []int
		for _, fld := range flds[1:] {
			numbers = append(numbers, mustInt(fld))
		}

		inp := input{testvalue, numbers}
		inputs = append(inputs, inp)
	}
	return
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func mustRead(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
