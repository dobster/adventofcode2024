package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	ax, ay         int
	bx, by         int
	prizex, prizey int
}

func main() {
	process("sample.txt")
	process("input.txt")
}

func process(filename string) {
	inp1 := readInput(filename)
	fmt.Printf("%s part1: %d\n", filename, part1(inp1))

	inp2 := readInput(filename)
	fmt.Printf("%s part2: %d\n", filename, part2(inp2))
}

func readInput(filename string) (g []game) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "Button A:") {
			buttonA := strings.Split(lines[i], ": ")[1]
			buttonAB := strings.Split(buttonA, ", ")
			ax := mustInt(strings.Split(buttonAB[0], "+")[1])
			ay := mustInt(strings.Split(buttonAB[1], "+")[1])

			i++
			buttonB := strings.Split(lines[i], ": ")[1]
			buttonBB := strings.Split(buttonB, ", ")
			bx := mustInt(strings.Split(buttonBB[0], "+")[1])
			by := mustInt(strings.Split(buttonBB[1], "+")[1])

			i++
			prize := strings.Split(lines[i], ": ")[1]
			prizeXY := strings.Split(prize, ", ")
			px := mustInt(strings.Split(prizeXY[0], "=")[1])
			py := mustInt(strings.Split(prizeXY[1], "=")[1])

			g = append(g, game{ax, ay, bx, by, px, py})
		}
	}

	return
}

func part1(g []game) (result int) {
	for _, gg := range g {
		// fmt.Println(gg.ax, gg.ay, gg.bx, gg.by, gg.prizex, gg.prizey)
		result += solve(gg)
	}

	return
}

func solve(g game) (tokens int) {
	bPresses := (g.prizex / g.bx)
	if bPresses*g.bx == g.prizex && bPresses*g.by == g.prizey {
		return bPresses
	}

	for ; bPresses >= 0; bPresses-- {
		aPresses := (g.prizex - bPresses*g.bx) / g.ax
		if (bPresses*g.bx+aPresses*g.ax) == g.prizex && (bPresses*g.by+aPresses*g.ay) == g.prizey {
			return bPresses + aPresses*3
		}
	}

	return 0
}

func part2(g []game) (result int) {
	for _, gm := range g {
		gm.prizex += 10000000000000
		gm.prizey += 10000000000000
		result += solve2(gm)
	}
	return
}

func solve2(g game) (tokens int) {
	if (g.ay*g.prizex-g.ax*g.prizey)%(g.ay*g.bx-g.ax*g.by) != 0 {
		return
	}
	bPresses := (g.ay*g.prizex - g.ax*g.prizey) / (g.ay*g.bx - g.ax*g.by)
	if (g.prizex-bPresses*g.bx)%g.ax != 0 {
		return
	}
	aPresses := (g.prizex - bPresses*g.bx) / g.ax
	tokens = aPresses*3 + bPresses
	// fmt.Printf("aPresses=%d bPresses=%d tokens=%d\n", aPresses, bPresses, tokens)
	return
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
