package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type puzzle struct {
	ord    map[int][]int
	prints [][]int
}

func main() {
	solve("sample.txt")
	solve("input.txt")
}

func solve(filename string) {
	inp := readInput(getInput(filename))
	fmt.Printf("%s part1: %d\n", filename, inp.part1())
	fmt.Printf("%s part2: %d\n", filename, inp.part2())
}

func (p puzzle) part1() (result int) {
	for _, list := range p.prints {
		if p.correct(list) {
			result += middle(list)
		}
	}
	return result
}

func (p puzzle) correct(list []int) bool {
	seen := make(map[int]bool)
	for _, pg := range list {
		for _, otherpg := range p.ord[pg] {
			if seen[otherpg] {
				return false
			}
		}
		seen[pg] = true
	}
	return true
}

func (p puzzle) reorder(idx int) {
	seen := make(map[int]int)
	for i := 0; i < len(p.prints[idx]); i++ {
		pg := p.prints[idx][i]
		for _, otherpg := range p.ord[pg] {
			x, ok := seen[otherpg]
			if ok {
				p.prints[idx][i], p.prints[idx][x] = p.prints[idx][x], p.prints[idx][i]
				return
			}
		}
		seen[pg] = i
	}
	return
}

func middle(list []int) int {
	i := len(list) / 2
	return list[i]
}

func (p puzzle) part2() (result int) {
	for i := 0; i < len(p.prints); i++ {
		if !p.correct(p.prints[i]) {
			for !p.correct(p.prints[i]) {
				p.reorder(i)
			}
			result += middle(p.prints[i])
		}
	}
	return result
}

func getInput(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func readInput(lines []string) (p puzzle) {
	p.ord = make(map[int][]int)
	i := 0
	for lines[i] != "" {
		flds := strings.Split(lines[i], "|")
		i1, i2 := mustInt(flds[0]), mustInt(flds[1])
		p.ord[i1] = append(p.ord[i1], i2)
		i++
	}
	i++
	for i < len(lines) {
		var list []int
		flds := strings.Split(lines[i], ",")
		for _, fld := range flds {
			list = append(list, mustInt(fld))
		}
		p.prints = append(p.prints, list)
		i++
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
