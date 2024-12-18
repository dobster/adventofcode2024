package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type pos struct {
	row, col int
}

type puzzle struct {
	plots   map[pos]string // pos -> crop
	crops   map[pos]int    // pos -> region
	regions map[int][]pos  // region -> []pos
}

func main() {
	process("sample.txt")
	process("sample2.txt")
	process("sample3.txt")
	process("input.txt")
}

func process(filename string) {
	inp1 := readInput(filename)
	fmt.Printf("%s part1: %d\n", filename, inp1.part1())

	inp2 := readInput(filename)
	fmt.Printf("%s part2: %d\n", filename, inp2.part2())
}

func readInput(filename string) (pz puzzle) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	pz.plots = make(map[pos]string)
	pz.regions = make(map[int][]pos)
	pz.crops = make(map[pos]int)

	for row, line := range strings.Split(string(b), "\n") {
		for col, ch := range line {
			pz.plots[pos{row, col}] = string(ch)
		}

	}

	return
}

func (pz *puzzle) part1() (total int) {
	region := 0
	for p, crop := range pz.plots {
		if _, ok := pz.crops[p]; !ok {
			region++
			pz.assignRegion(p, crop, region)
		}
	}

	for r := 1; r <= region; r++ {
		price := pz.perimeter(r) * pz.area(r)
		total += price
	}

	return
}

func (pz *puzzle) assignRegion(p pos, plant string, region int) {
	if pz.plots[p] == plant && pz.crops[p] == 0 {
		pz.crops[p] = region
		pz.regions[region] = append(pz.regions[region], p)

		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			pz.assignRegion(pos{p.row + dir[0], p.col + dir[1]}, plant, region)
		}
	}
}

func (pz *puzzle) perimeter(region int) (perim int) {
	maxrow, maxcol := 0, 0
	minrow, mincol := math.MaxInt, math.MaxInt

	cells := make(map[pos]bool)

	for _, pos := range pz.regions[region] {
		if pos.row < minrow {
			minrow = pos.row
		}
		if pos.row > maxrow {
			maxrow = pos.row
		}
		if pos.col < mincol {
			mincol = pos.col
		}
		if pos.col > maxcol {
			maxcol = pos.col
		}
		cells[pos] = true
	}

	for col := mincol - 1; col <= maxcol+1; col++ {
		in := false
		for row := minrow - 1; row <= maxrow+1; row++ {
			if cells[pos{row, col}] {
				if !in {
					perim++
					in = true
				}
			} else if in {
				perim++
				in = false
			}
		}
	}

	for row := minrow - 1; row <= maxrow+1; row++ {
		in := false
		for col := mincol - 1; col <= maxcol+1; col++ {
			if cells[pos{row, col}] {
				if !in {
					perim++
					in = true
				}
			} else if in {
				perim++
				in = false
			}
		}
	}

	return
}

func (pz *puzzle) area(region int) int {
	return len(pz.regions[region])
}

func (pz *puzzle) allSides(region int) (sides int) {
	maxrow, maxcol := 0, 0
	minrow, mincol := math.MaxInt, math.MaxInt

	cells := make(map[pos]bool)

	for _, pos := range pz.regions[region] {
		if pos.row < minrow {
			minrow = pos.row
		}
		if pos.row > maxrow {
			maxrow = pos.row
		}
		if pos.col < mincol {
			mincol = pos.col
		}
		if pos.col > maxcol {
			maxcol = pos.col
		}
		cells[pos] = true
	}

	for col := mincol - 1; col <= maxcol+1; col++ {
		for row := minrow - 1; row <= maxrow+1; row++ {
			if cells[pos{row, col}] {
				if !cells[pos{row - 1, col}] && (cells[pos{row - 1, col - 1}] || !cells[pos{row, col - 1}]) {
					sides++
				}
				if !cells[pos{row + 1, col}] && (cells[pos{row + 1, col - 1}] || !cells[pos{row, col - 1}]) {
					sides++
				}
			}
		}
	}

	for row := minrow - 1; row <= maxrow+1; row++ {
		for col := mincol - 1; col <= maxcol+1; col++ {
			if cells[pos{row, col}] {
				if (!cells[pos{row, col - 1}] && (cells[pos{row - 1, col - 1}] || !cells[pos{row - 1, col}])) {
					sides++
				}
				if (!cells[pos{row, col + 1}] && (cells[pos{row - 1, col + 1}] || !cells[pos{row - 1, col}])) {
					sides++
				}
			}
		}
	}

	return
}

func (pz *puzzle) part2() (total int) {
	region := 0
	for p, crop := range pz.plots {
		if _, ok := pz.crops[p]; !ok {
			region++
			pz.assignRegion(p, crop, region)
		}
	}

	for r := 1; r <= region; r++ {
		price := pz.area(r) * pz.allSides(r)

		// fmt.Printf("region %s area %d sides %d price %d\n", pz.plots[pz.regions[r][0]], pz.area(r), pz.allSides(r), price)

		total += price
	}

	return
}
