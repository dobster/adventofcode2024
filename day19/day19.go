package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type puzzle struct {
	towels  []string
	designs []string
	m       map[string]bool
}

func readInput(input string) *puzzle {
	lines := strings.Split(input, "\n")
	towels := strings.Split(lines[0], ", ")
	designs := []string{}
	for i := 2; i < len(lines); i++ {
		designs = append(designs, lines[i])
	}
	return &puzzle{towels, designs, make(map[string]bool)}
}

func (p *puzzle) isPossible(design string) bool {
	fmt.Printf("isPossible(%s)\n", design)
	if found, ok := p.m[design]; ok {
		return found
	}
	for _, towel := range p.towels {
		if strings.HasPrefix(design, towel) {
			remainingDesign := strings.TrimPrefix(design, towel)
			if p.isPossible(remainingDesign) {
				p.m[remainingDesign] = true
				return true
			} else {
				p.m[remainingDesign] = false
			}
		}
	}
	// fmt.Println("not possible")
	return false
}

func (p *puzzle) part1() int {
	var possible int
	for _, towel := range p.towels {
		p.m[towel] = true
	}
	for _, design := range p.designs {
		if p.isPossible(design) {
			fmt.Printf("%s is possible\n", design)
			possible++
		} else {
			fmt.Printf("%s is not\n", design)
		}
	}
	return possible
}

func (p *puzzle) possibilities(design string) int {
	fmt.Printf("isPossible(%s)\n", design)
	if count, ok := p.m[design]; ok {
		return count
	}
	for _, towel := range p.towels {
		if strings.HasPrefix(design, towel) {
			remainingDesign := strings.TrimPrefix(design, towel)
			if p.isPossible(remainingDesign) {
				p.m[remainingDesign] = true
				return true
			} else {
				p.m[remainingDesign] = false
			}
		}
	}
	// fmt.Println("not possible")
	return false
}

func (p *puzzle) part2() int {
	var possible int
	for _, towel := range p.towels {
		p.m[towel] = true
	}
	for _, design := range p.designs {
		if p.isPossible(design) {
			fmt.Printf("%s is possible\n", design)
			possible++
		} else {
			fmt.Printf("%s is not\n", design)
		}
	}
	return possible
}

func main() {
	fmt.Printf("part2: sample: %d\n", readInput(sample).part2())
	fmt.Printf("part2: input: %d\n", readInput(input).part2())
}
