package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	process("sample.txt")
	process("input.txt")
}

func process(filename string) {
	input := getInput(filename)
	fmt.Printf("%s part1: %d\n", filename, part1(input))
	fmt.Printf("%s part2: %d\n", filename, part2(input))
}

func getInput(filename string) [][]int {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var input [][]int

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		var nums []int
		for _, str := range strings.Fields(line) {
			nums = append(nums, mustInt(str))
		}
		input = append(input, nums)
	}

	return input
}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.

// How many reports are safe?

// 16 min
func part1(input [][]int) (result int) {
	for _, line := range input {
		if line[1] > line[0] {
			searching := true
			// increasing
			for i := 1; i < len(line) && searching; i++ {
				diff := line[i] - line[i-1]
				if diff < 1 || diff > 3 {
					searching = false
					break
				}
			}
			if searching {
				result++
			}
		} else if line[1] < line[0] {
			searching := true
			// decreasing
			for i := 1; i < len(line) && searching; i++ {
				diff := line[i-1] - line[i]
				if diff < 1 || diff > 3 {
					searching = false
					break
				}
			}
			if searching {
				result++
			}
		}
	}
	return result
}

// 15 min
func part2(input [][]int) (result int) {
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			a := make([]int, len(line))
			copy(a, line)
			a = slices.Delete(a, i, i+1)
			if isSafe(a) {
				result++
				break
			}
		}
	}
	return result
}

func isSafe(a []int) bool {
	// fmt.Println(a)
	if a[1] > a[0] {
		searching := true
		// increasing
		for i := 1; i < len(a) && searching; i++ {
			diff := a[i] - a[i-1]
			if diff < 1 || diff > 3 {
				searching = false
				break
			}
		}
		if searching {
			return true
		}
	} else if a[1] < a[0] {
		searching := true
		// decreasing
		for i := 1; i < len(a) && searching; i++ {
			diff := a[i-1] - a[i]
			if diff < 1 || diff > 3 {
				searching = false
				break
			}
		}
		if searching {
			return true
		}
	}
	return false
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
