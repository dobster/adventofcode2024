package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	process("sample1.txt", 1)
	process("sample2.txt", 6)
	process("sample2.txt", 25)
	process("input.txt", 25)
	process("input.txt", 75)
}

func process(filename string, blinks int) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var nums []int
	var nums2 []int

	for _, fld := range strings.Fields(string(b)) {
		nums = append(nums, mustInt(fld))
		nums2 = append(nums2, mustInt(fld))
	}

	// fmt.Printf("%s %d part1: %d\n", filename, blinks, part1(nums, blinks))
	fmt.Printf("%s %d part2: %d\n", filename, blinks, part2(nums2, blinks))
}

func part1(nums []int, blinks int) int {
	//fmt.Println(nums)

	for i := 0; i < blinks; i++ {
		for j := 0; j < len(nums); j++ {
			s := strconv.Itoa(nums[j])
			switch {
			case s == "0":
				nums[j] = 1
			case len(s)%2 == 0:
				half := len(s) / 2
				left := s[:half]
				right := s[half:]
				nums[j] = mustInt(left)
				if j == len(nums)-1 {
					nums = append(nums, mustInt(right))
				} else {
					nums = slices.Insert(nums, j+1, mustInt(right))
				}
				j++
			default:
				nums[j] *= 2024
			}
		}
		// fmt.Println(nums)
	}
	return len(nums)
}

type cache struct {
	mem map[[2]int]int
}

func part2(nums []int, blinks int) int {
	c := cache{make(map[[2]int]int)}
	score := 0
	for i := 0; i < len(nums); i++ {
		j := c.numStones(nums[i], blinks)
		// fmt.Printf("%d -> %d\n", nums[i], j)
		score += j
	}
	return score
}

func (c *cache) numStones(i int, blinks int) int {
	if blinks == 0 {
		return 1
	}
	s := strconv.Itoa(i)
	switch {
	case i == 0:
		return c.numStones(1, blinks-1)

	case len(s)%2 == 0:
		half := len(s) / 2

		left := mustInt(s[:half])
		sleft, ok := c.mem[[2]int{left, blinks - 1}]
		if !ok {
			sleft = c.numStones(left, blinks-1)
			c.mem[[2]int{left, blinks - 1}] = sleft
		}

		right := mustInt(s[half:])
		sright, ok := c.mem[[2]int{right, blinks - 1}]
		if !ok {
			sright = c.numStones(right, blinks-1)
			c.mem[[2]int{right, blinks - 1}] = sright
		}

		return sleft + sright

	default:
		return c.numStones(i*2024, blinks-1)
	}
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
