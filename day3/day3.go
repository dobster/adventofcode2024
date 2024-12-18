package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// process("sample3.txt")
	process("input.txt")
}

func process(filename string) {
	input := getInput(filename)
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

// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
// 161 (2*4 + 5*5 + 11*8 + 8*5)

// What do you get if you add up all of the results of the multiplications?
// func part1(input []string) (result int) {
// 	for _, line := range input {
// 		for i := 0; i < len(line); i++ {
// 			if strings.HasPrefix(line[i:], "mul(") {
// 				fmt.Println(line[i:i+12])
// 				i += 4
// 				var num1, num2 string
// 				for unicode.IsDigit(rune(line[i])) {
// 					num1 += string(line[i])
// 					i++
// 				}
// 				// fmt.Println("num1", num1)
// 				if line[i] == ',' {
// 					i++
// 					for unicode.IsDigit(rune(line[i])) {
// 						num2 += string(line[i])
// 						i++
// 					}
// 					// fmt.Println("num2", num2)
// 					if line[i] == ')' {
// 						result += mustInt(num1) * mustInt(num2)
// 						fmt.Printf("%s x %s\n", num1, num2)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

func part2(input []string) (result int) {
	doEnabled := true
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if strings.HasPrefix(line[i:], "don't()") {
				i += 6
				doEnabled = false
				continue
			}
			if strings.HasPrefix(line[i:], "do()") {
				i += 3
				doEnabled = true
				continue
			}
			if strings.HasPrefix(line[i:], "mul(") {
				j := i + 4
				var num1, num2 string
				for unicode.IsDigit(rune(line[j])) && len(num1) <= 2 {
					num1 += string(line[j])
					j++
				}
				if line[j] == ',' {
					j++
					for unicode.IsDigit(rune(line[j])) && len(num2) <= 2 {
						num2 += string(line[j])
						j++
					}
					if line[j] == ')' {
						if doEnabled {
							result += mustInt(num1) * mustInt(num2)
						}
						i = j
						continue
					}
				}
			}
		}
	}
	return result
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
