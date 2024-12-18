package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, filename := range []string{"sample.txt", "input.txt"} {
		inp := getInput(filename)
		fmt.Printf("%s part1: %d\n", filename, part1(inp))
		fmt.Printf("%s part2: %d\n", filename, part2(inp))
	}
}

func part1(inp [][]byte) (result int) {
	m := make(map[byte][][2]int)
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[y]); x++ {
			freq := inp[y][x]
			if freq != '.' {
				m[freq] = append(m[freq], [2]int{y, x})
			}
		}
	}

	// fmt.Println(m)

	locs := make(map[[2]int]int)
	for _, positions := range m {
		for _, pos1 := range positions {
			for _, pos2 := range positions {
				diffx, diffy := pos1[1]-pos2[1], pos1[0]-pos2[0]

				if diffx == 0 && diffy == 0 {
					continue
				}

				antinode := [2]int{pos1[0] + diffy, pos1[1] + diffx}
				// fmt.Printf("pairs [%v] and [%v] gives antinode [%v]\n", pos1, pos2, antinode)
				if isValidAntinode(antinode, inp) {
					locs[antinode]++
				}

				antinode2 := [2]int{pos2[0] - diffy, pos2[1] - diffx}
				if isValidAntinode(antinode2, inp) {
					locs[antinode2]++
				}

			}
		}
	}
	result += len(locs)

	return
}

func dump(inp [][]byte, locs map[[2]int]int) {
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[0]); x++ {
			cnt, ok := locs[[2]int{y, x}]
			if ok {
				fmt.Printf("%d", cnt)
			} else {
				fmt.Printf("%c", inp[y][x])
			}
		}
		fmt.Println()
	}
}

func isValidAntinode(pos [2]int, inp [][]byte) bool {
	return pos[0] >= 0 && pos[0] < len(inp) && pos[1] >= 0 && pos[1] < len(inp[0])
}

func part2(inp [][]byte) (result int) {
	m := make(map[byte][][2]int)
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[y]); x++ {
			freq := inp[y][x]
			if freq != '.' {
				m[freq] = append(m[freq], [2]int{y, x})
			}
		}
	}

	locs := make(map[[2]int]int)
	for _, positions := range m {
		for _, pos1 := range positions {
			for _, pos2 := range positions {
				diffx, diffy := pos1[1]-pos2[1], pos1[0]-pos2[0]

				if diffx == 0 && diffy == 0 {
					continue
				}

				antinode1 := [2]int{pos1[0], pos1[1]}
				for isValidAntinode(antinode1, inp) {
					locs[antinode1]++
					antinode1 = [2]int{antinode1[0] + diffy, antinode1[1] + diffx}
				}

				antinode2 := [2]int{pos2[0], pos2[1]}
				for isValidAntinode(antinode2, inp) {
					locs[antinode2]++
					antinode2 = [2]int{antinode2[0] - diffy, antinode2[1] - diffx}
				}
			}
		}
	}

	return len(locs)
}

func getInput(filename string) (inp [][]byte) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(b), "\n") {
		inp = append(inp, []byte(line))
	}
	return
}
