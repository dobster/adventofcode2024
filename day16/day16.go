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
	fmt.Printf("%s part1: %d\n", filename, part1(inp))
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

func dump(m map[pos]string, h *PathHeap) {
	v := make(map[pos]string)
	for _, p := range *h {
		v[p.p] = "X"
	}

	for y := 0; y < 15; y++ {
		for x := 0; x < 15; x++ {
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

type key struct {
	p pos
	d dir
}

func part1(m map[pos]string) int {
	r := findReindeer(m)
	d := dir{1, 0} // East

	explored := make(map[key]int)

	h := &PathHeap{Path{r, d, 0}}
	heap.Init(h)

	for {
		dump(m, h)

		path := heap.Pop(h).(Path)
		// fmt.Printf("considering %v\n", path)

		fwdPos := pos{path.p.x + path.d.dx, path.p.y + path.d.dy}

		if m[fwdPos] == "E" {
			return path.cost + 1
		}

		if m[fwdPos] == "." {
			fwdPath := Path{p: fwdPos, d: d, cost: path.cost + 1}
			val, ok := explored[key{fwdPath.p, fwdPath.d}]
			if !ok || val > fwdPath.cost {
				heap.Push(h, fwdPath)
			}
		}

		rightDir := rightTurn(path.d)
		rightPos := pos{path.p.x + rightDir.dx, path.p.y + rightDir.dy}
		if m[rightPos] == "." {
			rightPath := Path{rightPos, rightDir, path.cost + 1000 + 1}
			val, ok := explored[key{rightPath.p, rightPath.d}]
			if !ok || val > rightPath.cost {
				heap.Push(h, rightPath)
			}
		}

		leftDir := leftTurn(path.d)
		leftPos := pos{path.p.x + leftDir.dx, path.p.y + leftDir.dy}
		if m[leftPos] == "." {
			leftPath := Path{leftPos, leftDir, path.cost + 1000 + 1}
			val, ok := explored[key{leftPath.p, leftPath.d}]
			if !ok || val > leftPath.cost {
				heap.Push(h, leftPath)
			}
		}

		explored[key{path.p, path.d}] = path.cost
	}
}

func main() {
	process("sample.txt")
}
