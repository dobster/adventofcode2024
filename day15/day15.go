package main

import (
	"fmt"
	"os"
	"strings"
)

// The GPS coordinate of a box is equal to 100 times its distance from the top edge of the map plus its distance from the left edge of the map.
// (This process does not stop at wall tiles; measure all the way to the edges of the map.)

// So, the box shown below has a distance of 1 from the top edge of the map and 4 from the left edge of the map, resulting in a GPS coordinate of 100 * 1 + 4 = 104.
// #######
// #...O..
// #......

// The lanternfish would like to know the sum of all boxes' GPS coordinates after the robot finishes moving.
// In the larger example, the sum of all boxes' GPS coordinates is 10092. In the smaller example, the sum is 2028.

type pos struct {
	row, col int
}

type dir struct {
	xrow, xcol int
}

func readInput(filename string) (maze map[pos]string, instr []dir, robot pos) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")

	maze = make(map[pos]string)
	instr = []dir{}

	i := 0
	for ; strings.HasPrefix(lines[i], "#"); i++ {
		for j := 0; j < len(lines[i]); j++ {
			p := pos{i, j}
			s := string(lines[i][j])
			maze[p] = s
			if s == "@" {
				robot = p
			}
		}
	}

	i++

	for ; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			var d dir
			switch lines[i][j] {
			case '<':
				d = dir{0, -1}
			case '^':
				d = dir{-1, 0}
			case '>':
				d = dir{0, 1}
			case 'v':
				d = dir{1, 0}
			}
			instr = append(instr, d)
		}
	}

	return
}

func readInput2(filename string) (maze map[pos]string, instr []dir, robot pos) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")

	maze = make(map[pos]string)
	instr = []dir{}

	i := 0
	for ; strings.HasPrefix(lines[i], "#"); i++ {
		for j := 0; j < len(lines[i]); j++ {
			p1 := pos{i, j * 2}
			p2 := pos{i, j*2 + 1}
			s := string(lines[i][j])

			switch s {
			case "#":
				maze[p1] = "#"
				maze[p2] = "#"
			case "O":
				maze[p1] = "["
				maze[p2] = "]"
			case ".":
				maze[p1] = "."
				maze[p2] = "."
			case "@":
				maze[p1] = "@"
				maze[p2] = "."
				robot = p1
			}
		}
	}

	i++

	for ; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			var d dir
			switch lines[i][j] {
			case '<':
				d = dir{0, -1}
			case '^':
				d = dir{-1, 0}
			case '>':
				d = dir{0, 1}
			case 'v':
				d = dir{1, 0}
			}
			instr = append(instr, d)
		}
	}

	return
}

func dump(maze map[pos]string) {
	for row := 0; maze[pos{row, 0}] != ""; row++ {
		for col := 0; maze[pos{row, col}] != ""; col++ {
			fmt.Printf("%s", maze[pos{row, col}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(filename string) (result int) {
	maze, instr, robot := readInput(filename)

	for _, d := range instr {
		// fmt.Println(d)
		// dump(maze)
		p := robot

		p.row += d.xrow
		p.col += d.xcol

		switch maze[p] {
		case "#":
			continue
		case ".":
			maze[robot] = "."
			robot = p
			maze[robot] = "@"
			continue
		case "O":
			q := p
			for maze[q] == "O" {
				q.row, q.col = q.row+d.xrow, q.col+d.xcol
			}
			if maze[q] == "." {
				for q != robot {
					maze[q] = maze[pos{q.row - d.xrow, q.col - d.xcol}]
					q.row, q.col = q.row-d.xrow, q.col-d.xcol
				}
				maze[robot] = "."
				robot = p
				maze[robot] = "@"
			}
			continue
		}
	}

	for p, c := range maze {
		if c == "O" {
			result += 100*p.row + p.col
		}
	}

	return
}

func part2(filename string) (result int) {
	maze, instr, robot := readInput2(filename)

instrloop:
	for _, d := range instr {
		// fmt.Println(d)
		// dump(maze)
		p := robot

		p.row += d.xrow
		p.col += d.xcol

		switch maze[p] {
		case "#":
			// fmt.Println("found a brick wall!")
			continue instrloop
		case ".":
			maze[robot] = "."
			robot = p
			maze[robot] = "@"
			continue instrloop
		case "[", "]":
			q := []pos{robot, p}
			if d.xcol != 0 {
				// fmt.Println("left/right")
				for {
					p = pos{p.row, p.col + d.xcol}
					if maze[p] == "#" {
						continue instrloop
					}
					if maze[p] == "." {
						break
					}
					q = append(q, p)
				}
			} else {
				// fmt.Println("up/down")
				if maze[p] == "[" {
					q = append(q, pos{p.row, p.col + 1})
				}
				if maze[p] == "]" {
					q = append(q, pos{p.row, p.col - 1})
				}
				// fmt.Printf("checking q=%v\n", q)
				boxes := true
				for row := p.row; boxes; row += d.xrow {
					boxes = false
					new := []pos{}
					for i := 0; i < len(q); i++ {
						if q[i].row == row {
							// fmt.Printf("row %d: considering whats up/down from %v\n", row, q[i])
							switch maze[pos{row + d.xrow, q[i].col}] {
							case "#":
								continue instrloop
							case ".":
								continue
							case "[":
								new = append(new, pos{row + d.xrow, q[i].col})
								new = append(new, pos{row + d.xrow, q[i].col + 1})
								boxes = true
							case "]":
								new = append(new, pos{row + d.xrow, q[i].col})
								new = append(new, pos{row + d.xrow, q[i].col - 1})
								boxes = true
							}
						}
					}
					for i := 0; i < len(new); i++ {
						if !contains(q, new[i]) {
							q = append(q, new[i])
						}
					}
				}
			}

			// fmt.Printf("good to move, q=%v\n", q)

			for i := len(q) - 1; i >= 0; i-- {
				r := pos{q[i].row + d.xrow, q[i].col + d.xcol}
				maze[q[i]], maze[r] = maze[r], maze[q[i]]
			}

			robot = pos{robot.row + d.xrow, robot.col + d.xcol}
		}
	}
	dump(maze)

	for p, c := range maze {
		if c == "[" {
			result += 100*p.row + p.col
		}
	}

	return
}

func contains(q []pos, p pos) bool {
	for i := 0; i < len(q); i++ {
		if q[i].row == p.row && q[i].col == p.col {
			return true
		}
	}
	return false
}

func process(filename string) {
	// fmt.Printf("%s part1: %d\n", filename, part1(filename))
	fmt.Printf("%s part2: %d\n", filename, part2(filename))
}

func main() {
	// process("sample.txt")
	//process("sample2.txt")
	// process("sample3.txt")
	process("input.txt")
}
