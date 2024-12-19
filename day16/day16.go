package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

type dir struct {
	dx, dy int
}

func readInput(filename string) map[pos]string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	m := make(map[pos]string)

	for y, line := range lines {
		for x, ch := range line {
			m[pos{x, y}] = string(ch)
		}
	}

	return m
}

func process(filename string) {
	inp := readInput(filename)
	part1result := part1(inp)
	fmt.Printf("%s part1: %d\n", filename, part1result)
	fmt.Printf("%s part2: %d\n", filename, part2(inp, part1result))
}

func findReindeer(m map[pos]string) pos {
	for k, v := range m {
		if v == "S" {
			return k
		}
	}
	panic("no reindeer found!")
}

type Path struct {
	p    pos
	d    dir
	cost int
	path []pos
}

type PathHeap []Path

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PathHeap) Push(p any) {
	*h = append(*h, p.(Path))
}

func (h PathHeap) Len() int {
	return len(h)
}

func (h PathHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h PathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h PathHeap) Top() Path {
	return h[0]
}

func rightTurn(d dir) dir {
	switch d {
	case dir{1, 0}: // E
		return dir{0, 1} // S
	case dir{0, 1}: // S
		return dir{-1, 0} // W
	case dir{-1, 0}: // W
		return dir{0, -1} // N
	case dir{0, -1}: // N
		return dir{1, 0} // E
	default:
		panic("so confused")
	}
}

func leftTurn(d dir) dir {
	switch d {
	case dir{1, 0}: // E
		return dir{0, -1} // N
	case dir{0, -1}: // N
		return dir{-1, 0} // W
	case dir{-1, 0}: // W
		return dir{0, 1} // S
	case dir{0, 1}: // S
		return dir{1, 0} // E
	default:
		panic("so confused")
	}
}

func direction(d dir) string {
	switch d {
	case dir{1, 0}:
		return ">"
	case dir{0, -1}:
		return "^"
	case dir{-1, 0}:
		return "<"
	case dir{0, 1}:
		return "v"
	default:
		return "?"
	}
}

func dump(m map[pos]string, h *PathHeap) {
	v := make(map[pos]string)
	for _, p := range *h {
		v[p.p] = direction(p.d)
	}

	for y := 0; y < 142; y++ {
		for x := 0; x < 142; x++ {
			if z, ok := v[pos{x, y}]; ok {
				fmt.Printf(z)
			} else if z, ok := m[pos{x, y}]; ok {
				fmt.Printf("%s", z)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func dumpBestPath(m map[pos]string, bestPathTiles []pos) {
	for y := 0; y < 15; y++ {
		for x := 0; x < 15; x++ {
			if contains(bestPathTiles, pos{x, y}) {
				fmt.Printf("O")
			} else if z, ok := m[pos{x, y}]; ok {
				fmt.Printf("%s", z)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(m map[pos]string) int {
	initp := findReindeer(m)
	initd := dir{1, 0} // East

	explored := make(map[pos]int)

	h := &PathHeap{Path{initp, initd, 0, []pos{}}}
	heap.Init(h)

	for {
		// dump(m, h)

		path := heap.Pop(h).(Path)
		// fmt.Printf("exploring %v %s\n", path.p, direction(path.d))

		fwdPos := pos{path.p.x + path.d.dx, path.p.y + path.d.dy}

		if m[fwdPos] == "E" {
			return path.cost + 1
		}

		if m[fwdPos] == "." && explored[fwdPos] == 0 {
			fwdPath := Path{fwdPos, path.d, path.cost + 1, []pos{}}
			heap.Push(h, fwdPath)
		}

		rightDir := rightTurn(path.d)
		rightPos := pos{path.p.x + rightDir.dx, path.p.y + rightDir.dy}
		if m[rightPos] == "." && explored[rightPos] == 0 {
			rightPath := Path{rightPos, rightDir, path.cost + 1000 + 1, []pos{}}
			heap.Push(h, rightPath)
		}

		leftDir := leftTurn(path.d)
		leftPos := pos{path.p.x + leftDir.dx, path.p.y + leftDir.dy}
		if m[leftPos] == "." && explored[leftPos] == 0 {
			leftPath := Path{leftPos, leftDir, path.cost + 1000 + 1, []pos{}}
			heap.Push(h, leftPath)
		}

		explored[path.p] = path.cost
	}
}

func contains(paths []pos, p pos) bool {
	for _, path := range paths {
		if path.x == p.x && path.y == p.y {
			return true
		}
	}
	return false
}

func newPath(path []pos, p pos) []pos {
	newp := make([]pos, len(path)+1)
	copy(newp, path)
	newp[len(path)] = p
	return newp
}

type key struct {
	p pos
	d dir
}

func part2(m map[pos]string, bestCost int) int {
	initp := findReindeer(m)
	initd := dir{1, 0} // East

	explored := make(map[key]int)

	h := &PathHeap{Path{initp, initd, 0, []pos{initp}}}
	heap.Init(h)

	bestTiles := []pos{}

	for len(*h) > 0 {
		path := heap.Pop(h).(Path)

		if path.cost > bestCost {
			break
		}

		fwdPos := pos{path.p.x + path.d.dx, path.p.y + path.d.dy}

		if m[fwdPos] == "E" {
			if path.cost+1 == bestCost {
				for _, tile := range path.path {
					bestTiles = append(bestTiles, tile)
				}
				bestTiles = append(bestTiles, fwdPos)
			} 
			continue
		}

		fwdPath := Path{fwdPos, path.d, path.cost + 1, newPath(path.path, fwdPos)}
		kFwd := key{fwdPath.p, fwdPath.d}
		if m[fwdPos] == "." && (explored[kFwd] == 0 || explored[kFwd] == fwdPath.cost) {
			heap.Push(h, fwdPath)
		}

		rightDir := rightTurn(path.d)
		rightPos := pos{path.p.x + rightDir.dx, path.p.y + rightDir.dy}
		rightPath := Path{rightPos, rightDir, path.cost + 1000 + 1, newPath(path.path, rightPos)}
		kRight := key{rightPath.p, rightPath.d}
		if m[rightPos] == "." && (explored[kRight] == 0 || explored[kRight] == path.cost+1000+1) {
			heap.Push(h, rightPath)
		}

		leftDir := leftTurn(path.d)
		leftPos := pos{path.p.x + leftDir.dx, path.p.y + leftDir.dy}
		leftPath := Path{leftPos, leftDir, path.cost + 1000 + 1, newPath(path.path, leftPos)}
		kLeft := key{leftPath.p, leftPath.d}
		if m[leftPos] == "." && (explored[kLeft] == 0 || explored[kLeft] == path.cost+1000+1) {
			heap.Push(h, leftPath)
		}

		explored[key{path.p, path.d}] = path.cost
	}

	bestTileMap := make(map[pos]bool)
	for _, tile := range bestTiles {
		bestTileMap[tile] = true
	}

	return len(bestTileMap)
}

func main() {
	process("sample.txt")
	process("sample2.txt")
	process("input.txt")
}
