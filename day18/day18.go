package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type pos struct {
	x, y int
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func readInput(s string) []pos {
	inp := []pos{}
	for _, line := range strings.Split(s, "\n") {
		comps := strings.Split(line, ",")
		inp = append(inp, pos{mustInt(comps[0]), mustInt(comps[1])})
	}
	return inp
}

func dump(grid map[pos]int, size int, exit pos) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			p := pos{x, y}
			if p == exit {
				fmt.Printf("E")
				continue
			}
			i, ok := grid[p]
			if !ok {
				panic("not found")
			}
			if i == -1 {
				fmt.Printf("#")
				continue
			}
			if i == 0 {
				fmt.Printf(".")
				continue
			}
			fmt.Printf("O")
		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(s string, size int, bytes int) int {
	incoming := readInput(s)

	grid := make(map[pos]int)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			grid[pos{x, y}] = 0
		}
	}

	for i := 0; i < bytes; i++ {
		grid[incoming[i]] = -1
	}

	exit := pos{size - 1, size - 1}
	start := pos{0, 0}

	lastCells := []pos{start}
	steps := 0
	for len(lastCells) > 0 {
		steps++
		nextCells := []pos{}
		for _, cell := range lastCells {
			for _, dir := range []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				p := pos{cell.x + dir.x, cell.y + dir.y}
				i, ok := grid[p]
				if ok && i == 0 {
					nextCells = append(nextCells, p)
					grid[p] = steps
					if p == exit {
						return steps
					}
				}
			}
		}
		lastCells = nextCells
	}
	panic("no way through!")
}

func part2(s string, size int) pos {
	incoming := readInput(s)

bytesloop:
	for bytes := 1; ; bytes++ {

		grid := make(map[pos]int)
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				grid[pos{x, y}] = 0
			}
		}

		for i := 0; i < bytes; i++ {
			grid[incoming[i]] = -1
		}

		exit := pos{size - 1, size - 1}
		start := pos{0, 0}

		lastCells := []pos{start}
		steps := 0
		for len(lastCells) > 0 {
			steps++
			nextCells := []pos{}
			for _, cell := range lastCells {
				for _, dir := range []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					p := pos{cell.x + dir.x, cell.y + dir.y}
					i, ok := grid[p]
					if ok && i == 0 {
						nextCells = append(nextCells, p)
						grid[p] = steps
						if p == exit {
							continue bytesloop
						}
					}
				}
			}
			lastCells = nextCells
		}
		return incoming[bytes-1]
	}
}

func main() {
	// fmt.Printf("part1: sample: %d\n", part1(sample, 7, 12))
	fmt.Printf("part2: sample: %v\n", part2(sample, 7))
	// fmt.Printf("part1: input: %d\n", part1(input, 71, 1024))
	fmt.Printf("part2: input: %v\n", part2(input, 71))
}
