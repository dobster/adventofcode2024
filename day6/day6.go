package main

import (
	"fmt"
	"os"
	"strings"
)

type puzzle struct {
	m [][]byte
}

type pos struct {
	x, y int
}

type key struct {
	xy  pos
	dir pos
}

func (p puzzle) isInBounds(g pos) bool {
	return g.x >= 0 && g.x < len(p.m) && g.y >= 0 && g.y < len(p.m[0])
}

func rotate(g pos) pos {
	switch {
	case g.x == 0 && g.y == -1:
		return pos{1, 0}
	case g.x == 1 && g.y == 0:
		return pos{0, 1}
	case g.x == 0 && g.y == 1:
		return pos{-1, 0}
	case g.x == -1 && g.y == 0:
		return pos{0, -1}
	default:
		panic("really confused now")
	}
}

func (p puzzle) startpos() pos {
	for row := 0; row < len(p.m); row++ {
		for col := 0; col < len(p.m[0]); col++ {
			if p.m[row][col] == '^' {
				return pos{col, row}
			}
		}
	}
	panic("could not find start")
}

func (p puzzle) dump(g pos) {
	for i := 0; i < len(p.m); i++ {
		for j := 0; j < len(p.m[0]); j++ {
			if g.x == j && g.y == i {
				fmt.Printf("%c", 'G')
			} else {
				fmt.Printf("%c", p.m[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	process("sample.txt")
	process("input.txt")
}

func process(filename string) {
	p := parseInput(readInput(filename))
	fmt.Printf("%s: part1: %d\n", filename, p.part1())
	fmt.Printf("%s: part2: %d\n", filename, p.part2(filename))
}

func (p puzzle) part1() (result int) {
	return len(p.findLocs())
}

func (p puzzle) part2(filename string) (result int) {
	locs := p.findLocs()
	start := p.startpos()

	for _, loc := range locs {
		if loc.x != start.x || loc.y != start.y {

			m2 := parseInput(readInput(filename))

			m2.m[loc.y][loc.x] = '#'

			if m2.isLoop() {
				result++
			}

		}
	}

	return
}

func (p puzzle) isLoop() bool {
	g := p.startpos()
	visited := make(map[key]bool)
	dir := pos{0, -1}
	for p.isInBounds(g) {
		visited[key{xy: g, dir: dir}] = true

		next := pos{g.x + dir.x, g.y + dir.y}
		for p.isInBounds(next) && p.m[next.y][next.x] == '#' {
			dir = rotate(dir)
			next = pos{g.x + dir.x, g.y + dir.y}
		}

		g = pos{g.x + dir.x, g.y + dir.y}
		if visited[key{xy: g, dir: dir}] {
			return true
		}
	}
	return false
}

func (p puzzle) findLocs() (locs []pos) {
	g := p.startpos()
	visited := make(map[pos]bool)
	dir := pos{0, -1}
	for p.isInBounds(g) {
		visited[g] = true

		next := pos{g.x + dir.x, g.y + dir.y}
		for p.isInBounds(next) && p.m[next.y][next.x] == '#' {
			dir = rotate(dir)
			next = pos{g.x + dir.x, g.y + dir.y}
		}

		g = pos{g.x + dir.x, g.y + dir.y}
	}

	for k, _ := range visited {
		locs = append(locs, k)
	}

	return
}

func readInput(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func parseInput(inp []string) *puzzle {
	var p puzzle
	for _, line := range inp {
		p.m = append(p.m, []byte(line))
	}
	return &p
}
