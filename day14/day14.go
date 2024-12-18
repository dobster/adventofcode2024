package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type dir pos

type robot struct {
	p pos
	v dir
}

func readInput(filename string) (robots []robot) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		pv := strings.Split(line, " ")
		pp := strings.Split(pv[0], "=")
		ppp := strings.Split(pp[1], ",")
		p := pos{mustInt(ppp[0]), mustInt(ppp[1])}

		vv := strings.Split(pv[1], "=")
		vvv := strings.Split(vv[1], ",")
		v := dir{mustInt(vvv[0]), mustInt(vvv[1])}

		robots = append(robots, robot{p, v})
	}
	// fmt.Println(robots)
	return
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func process(filename string, wide, tall int, secs int) {
	inp := readInput(filename)
	fmt.Printf("%d robots\n", len(inp))
	// fmt.Printf("%s part1: %d\n", filename, part1(inp, wide, tall, secs))
	fmt.Printf("%s part2: %d\n", filename, part2(inp, wide, tall))
}

func part1(robots []robot, wide, tall int, secs int) int {
	for i := 0; i < secs; i++ {
		for r := 0; r < len(robots); r++ {
			robots[r].p.x += robots[r].v.x
			robots[r].p.y += robots[r].v.y

			if robots[r].p.x >= wide {
				robots[r].p.x -= wide
			}
			if robots[r].p.x < 0 {
				robots[r].p.x += wide
			}
			if robots[r].p.y >= tall {
				robots[r].p.y -= tall
			}
			if robots[r].p.y < 0 {
				robots[r].p.y += tall
			}
		}
	}

	var quadrant [4]int
	for _, robot := range robots {
		// TOP LEFT
		if robot.p.x < wide/2 && robot.p.y < tall/2 {
			quadrant[0]++
		}
		// TOP RIGHT
		if robot.p.x > wide/2 && robot.p.y < tall/2 {
			quadrant[1]++
		}
		// BOTTOM LEFT
		if robot.p.x < wide/2 && robot.p.y > tall/2 {
			quadrant[2]++
		}
		// BOTTOM RIGHT
		if robot.p.x > wide/2 && robot.p.y > tall/2 {
			quadrant[3]++
		}
	}

	return quadrant[0] * quadrant[1] * quadrant[2] * quadrant[3]
}

func part2(robots []robot, wide, tall int) int {
	// cycles := 0
	// // cycles := 10 // 57 + 46 + 55 + 48 + 53 + 50 + 51 + 52 + 49 + 54 + 47 + 56 + 45 + 58 + 43 + 60 + 41 + 62 + 39 + 64 + 37 + 66 + 35 + 68 + 33 + 70 + 31 + 72 + 29 + 74 + 27 + 76 + 25 + 78 + 23 + 80 + 21 + 82 + 19 + 84 + 17 + 86 + 15 + 88 + 13 + 90 + 11 + 92 + 7 + 94 + 5 + 96 + 3 + 98

	// for i := 0; i < 12; i++ {
	// 	moveRobots(robots, wide, tall)
	// 	cycles++
	// }

	// fmt.Println(cycles)
	// dump(robots, wide, tall)

	// cycle1, cycle2 := 57, 46

	// for i := 0; i < 100; i++ {
	// 	for j := 0; j < cycle1; j++ {
	// 		moveRobots(robots, wide, tall)
	// 		cycles++
	// 	}

	// 	fmt.Println(cycles)
	// 	dump(robots, wide, tall)

	// 	cycle1 -= 2

	// 	for j := 0; j < cycle2; j++ {
	// 		moveRobots(robots, wide, tall)
	// 		cycles++
	// 	}

	// 	fmt.Println(cycles)
	// 	dump(robots, wide, tall)

	// 	cycle2 += 2

	// }

	// return cycles

	cycles := 0
	for i := 0; i < 2999; i++ {
		moveRobots(robots, wide, tall)
		cycles++
	}

	for i := 0; i < 2000; i++ {
		for j := 0; j < 103; j++ {
			moveRobots(robots, wide, tall)
			cycles++
		}
		fmt.Println(cycles)
		dump(robots, wide, tall)
	}

	return cycles
}

func moveRobots(robots []robot, wide, tall int) {
	for r := 0; r < len(robots); r++ {
		robots[r].p.x += robots[r].v.x
		robots[r].p.y += robots[r].v.y

		if robots[r].p.x >= wide {
			robots[r].p.x -= wide
		}
		if robots[r].p.x < 0 {
			robots[r].p.x += wide
		}
		if robots[r].p.y >= tall {
			robots[r].p.y -= tall
		}
		if robots[r].p.y < 0 {
			robots[r].p.y += tall
		}
	}

}

func findTree(robots map[pos]int, first pos) bool {
	visited := make(map[pos]bool)
	visited[first] = true
	findFriends(robots, first, visited)
	return len(visited) > 50
}

func findFriends(robots map[pos]int, p pos, visited map[pos]bool) {
	if !visited[p] && robots[p] > 0 {
		visited[p] = true
		findFriends(robots, pos{p.x, p.y + 1}, visited)
		findFriends(robots, pos{p.x + 1, p.y}, visited)
		findFriends(robots, pos{p.x - 1, p.y}, visited)
		findFriends(robots, pos{p.x, p.y - 1}, visited)
	}
}

func checkVert(robots map[pos]int, wide, tall int) bool {
	last := 0
	for row := 0; row < 10; row++ {
		nrobots := 0
		for col := 0; col < wide; col++ {
			if robots[pos{row, col}] > 0 {
				nrobots++
			}
		}
		if nrobots != 0 && nrobots != last+1 {
			return false
		}
		last = nrobots
	}
	return true
}

func checkHoriz(robots map[pos]int, wide, tall int) bool {
	for col := 0; col < 10; col++ {
		nrobots := 0
		for row := 0; row < tall; row++ {
			if robots[pos{row, col}] > 0 {
				nrobots++
			}
		}
		if nrobots != col+1 {
			return false
		}
	}
	return true
}

func dump(robots []robot, wide, tall int) {
	m := make(map[pos]int)

	for _, robot := range robots {
		m[robot.p]++
	}

	for row := 0; row < tall; row++ {
		for col := 0; col < wide; col++ {
			if m[pos{row, col}] > 0 {
				fmt.Printf("%d", m[pos{row, col}])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// process("sample.txt", 11, 7, 100)
	process("input.txt", 101, 103, 100)
}
