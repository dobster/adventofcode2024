package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// want 36

func main() {
	process("sample1.txt")
	process("input.txt")
}

func process(filename string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s part1: %d\n", filename, part1(b))
	fmt.Printf("%s part2: %d\n", filename, part2(b))
}

type pos struct {
	r, c int
}

func isAdjacent(p1, p2 pos) bool {
	for _, d := range [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
		if p1.c+d[1] == p2.c && p1.r+d[0] == p2.r {
			return true
		}
	}
	return false
}

func part2(b []byte) int {
	heights := make(map[int][]pos)
	score := 0

	lines := strings.Split(string(b), "\n")
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			h := mustInt(string(lines[r][c]))
			heights[h] = append(heights[h], pos{r, c})
		}
	}

	for _, trailhead := range heights[0] {
		peaks := peaksReached(trailhead, 0, heights)
		score += len(peaks)
	}

	return score
}

func part1(b []byte) int {
	heights := make(map[int][]pos)
	score := 0

	lines := strings.Split(string(b), "\n")
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			h := mustInt(string(lines[r][c]))
			heights[h] = append(heights[h], pos{r, c})
		}
	}

	for _, trailhead := range heights[0] {
		peaks := peaksReached(trailhead, 0, heights)
		uniquePeaks := make(map[pos]bool)
		for _, peak := range peaks {
			uniquePeaks[peak] = true
		}
		// fmt.Printf("trailhead %v has %d score\n", trailhead, len(uniquePeaks))
		score += len(uniquePeaks)
	}

	return score
}

func peaksReached(p pos, h int, heights map[int][]pos) (peaks []pos) {
	// fmt.Printf("%v %d\n", p, h)
	if h == 9 {
		peaks = []pos{p}
	} else {
		for _, nexth := range heights[h+1] {
			if isAdjacent(nexth, p) {
				peaks = append(peaks, peaksReached(nexth, h+1, heights)...)
			}
		}
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
