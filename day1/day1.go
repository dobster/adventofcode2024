package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	process("sample.txt")
	process("input.txt")
}

func readData(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := string(b)
	lines := splitLines(s)
	return lines
}

func process(filename string) {
	lines := readData(filename)
	fmt.Printf("%20s: part1: %10d\n", filename, part1(lines))
	fmt.Printf("%20s: part2: %10d\n", filename, part2(lines))
	fmt.Println()
}

func part1(lines []string) (result int) {
	var l []int
	var m []int

	for _, a := range lines {
		comps := strings.Split(a, "   ")
		l = append(l, mustInt(comps[0]))
		m = append(m, mustInt(comps[1]))
	}

	slices.Sort(l)
	slices.Sort(m)

	for i := 0; i < len(l); i++ {
		result += abs(m[i] - l[i])
	}
	return
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func part2(lines []string) (result int) {
	var l []int
	var m []int

	lmap := make(map[int]int)
	mmap := make(map[int]int)

	for _, a := range lines {
		comps := strings.Split(a, "   ")
		l = append(l, mustInt(comps[0]))
		lmap[mustInt(comps[0])]++
		m = append(m, mustInt(comps[1]))
		mmap[mustInt(comps[1])]++
	}

	slices.Sort(l)
	slices.Sort(m)

	for i := 0; i < len(l); i++ {
		result += l[i] * mmap[l[i]]
	}
	return
}

func linesOfInts(lines []string) (a []int) {
	for _, line := range lines {
		a = append(a, mustInt(line))
	}
	return
}

func linesOfArrayOfInts(lines []string) (a [][]int) {
	for _, line := range lines {
		var b []int
		for _, word := range splitWords(line) {
			b = append(b, mustInt(word))
		}
		a = append(a, b)
	}
	return
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}

func splitWords(s string) []string {
	return strings.Split(s, " ")
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func mustFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func mustBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
