package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	process("sample.txt")
	process("input.txt")
}

func process(filename string) {
	input := getInput(filename)
	fmt.Printf("%s part1: %d\n", filename, part1(input))
	fmt.Printf("%s part2: %d\n", filename, part2(input))
}

func getInput(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	return lines
}

func getData(input []string) (a [][]byte) {
	for _, line := range input {
		a = append(a, []byte(line))
	}
	return
}

func check2(a []string, i, j [4]int) bool {
	for k := 0; k < 4; k++ {
		if i[k] < 0 || i[k] >= len(a) {
			return false
		}
		if j[k] < 0 || j[k] >= len(a[0]) {
			return false
		}
	}

	s := fmt.Sprintf("%c%c%c%c", a[i[0]][j[0]], a[i[1]][j[1]], a[i[2]][j[2]], a[i[3]][j[3]])
	if s == "XMAS" || s == "SAMX" {
		return true
	}

	return false
}

func check(a []string, i, j int) (count int) {
	// horizontal
	if check2(a, [4]int{i, i, i, i}, [4]int{j, j + 1, j + 2, j + 3}) {
		// fmt.Printf("%d %d HORIZONTAL\n", i, j)
		count++
	}

	// diagonal right
	if check2(a, [4]int{i, i + 1, i + 2, i + 3}, [4]int{j, j + 1, j + 2, j + 3}) {
		// fmt.Printf("%d %d DIAGONAL RIGHT\n", i, j)
		count++
	}

	// diagonal left
	if check2(a, [4]int{i, i + 1, i + 2, i + 3}, [4]int{j, j - 1, j - 2, j - 3}) {
		// fmt.Printf("%d %d DIAGONAL LEFT\n", i, j)
		count++
	}

	// vertical
	if check2(a, [4]int{i, i + 1, i + 2, i + 3}, [4]int{j, j, j, j}) {
		// fmt.Printf("%d %d VERTICAL\n", i, j)
		count++
	}

	return
}

func part1(a []string) (result int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			result += check(a, i, j)
		}
	}

	return
}

func checkP2(a []string, i, j int) (result int) {
	lr := fmt.Sprintf("%c%c%c", a[i-1][j-1], a[i][j], a[i+1][j+1])
	rl := fmt.Sprintf("%c%c%c", a[i-1][j+1], a[i][j], a[i+1][j-1])

	if (lr == "MAS" || lr == "SAM") && (rl == "MAS" || rl == "SAM") {
		result++
	}

	return
}

func part2(a []string) (result int) {
	for i := 1; i < len(a)-1; i++ {
		for j := 1; j < len(a[0])-1; j++ {
			result += checkP2(a, i, j)
		}
	}

	return
}
