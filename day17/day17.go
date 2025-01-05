package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type puzzle struct {
	a, b, c    int
	program    []int
	programStr string
}

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

//go:embed input.txt
var input string

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func mustInts(s string) []int {
	var ints []int
	for _, s := range strings.Split(s, ",") {
		ints = append(ints, mustInt(s))
	}
	return ints
}

func intsToString(ints []int) string {
	var s []string
	for _, i := range ints {
		s = append(s, strconv.Itoa(i))
	}
	return strings.Join(s, ",")
}

func readInput(s string) *puzzle {
	lines := strings.Split(s, "\n")
	a := mustInt(strings.TrimPrefix(lines[0], "Register A: "))
	b := mustInt(strings.TrimPrefix(lines[1], "Register B: "))
	c := mustInt(strings.TrimPrefix(lines[2], "Register C: "))
	programStr := strings.TrimPrefix(lines[4], "Program: ")
	program := mustInts(programStr)
	return &puzzle{a, b, c, program, programStr}
}

func (p puzzle) comboOperand(operand int, a, b, c int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	}
	panic("unknown operand")
}

func (p puzzle) part1() string {
	a, b, c := p.a, p.b, p.c
	instrPtr := 0
	output := []int{}
	for instrPtr < len(p.program) {
		operand := p.program[instrPtr+1]
		switch p.program[instrPtr] {
		case adv:
			numerator := a
			denominator := int(math.Pow(2, float64(p.comboOperand(operand, a, b, c))))
			a = numerator / denominator
		case bxl:
			b = b ^ operand
		case bst:
			b = p.comboOperand(operand, a, b, c) % 8
		case jnz:
			if a != 0 {
				instrPtr = operand
				continue
			}
		case bxc:
			b = b ^ c
		case out:
			output = append(output, p.comboOperand(operand, a, b, c)%8)
		case bdv:
			numerator := a
			denominator := int(math.Pow(2, float64(p.comboOperand(operand, a, b, c))))
			b = numerator / denominator
		case cdv:
			numerator := a
			denominator := int(math.Pow(2, float64(p.comboOperand(operand, a, b, c))))
			c = numerator / denominator
		}
		instrPtr += 2
	}
	return intsToString(output)
}

func (p puzzle) part1Direct() string {
	a, b, c := p.a, p.b, p.c

	output := []int{}

	for {
		b = a % 8
		b = b ^ 5
		c = a / int(math.Pow(2, float64(b)))
		b = b ^ 6
		a = a >> 3
		b = b ^ c
		output = append(output, b%8)
		if a == 0 {
			break
		}
	}

	return intsToString(output)
}

// func (p puzzle) part2(aa int) string {
// 	result := []string{}
// 	a, b, c := aa, 0, 0
// 	for i := 0; ; i++ {
// 		b = a % 8
// 		b = b ^ 5
// 		c = a >> b
// 		b = b ^ 6
// 		a = a >> 3
// 		b = b ^ c
// 		output := b % 8
// 		// if output != p.program[i] {
// 		// 	fmt.Println()
// 		// 	continue aaloop
// 		// }
// 		//	fmt.Printf("%d, ", output)
// 		result = append(result, strconv.Itoa(output))
// 		if a == 0 {
// 			break
// 		}
// 	}
// 	return strings.Join(result, ",")
// }

// func findPossibleAs() []big.Int {
// 	poss := []big.Int{}
// 	mult := big.NewInt(8)
// 	for _, numerator := range []int64{1, 2, 3, 4, 5, 6, 7} {
// 		a := big.NewInt(numerator)
// 		for i := 0; i < 15; i++ {
// 			a.Mul(a, mult)
// 		}
// 		poss = append(poss, *a)
// 	}
// 	return poss
// }

// func runProgram(a *big.Int) []int {
// 	k5 := big.NewInt(5)
// 	k6 := big.NewInt(6)
// 	k8 := big.NewInt(8)

// 	b, c := big.NewInt(0), big.NewInt(0)

// 	output := []int{}
// 	i := 0
// 	for a.String() != "0" {
// 		i++

// 		b.Mod(a, k8)
// 		b.Xor(b, k5)
// 		denominator := int64(math.Pow(2, float64(b.Int64())))
// 		c.Div(a, big.NewInt(denominator))
// 		b.Xor(b, k6)
// 		a.Div(a, k8)
// 		b.Xor(b, c)
// 		b.Mod(b, k8)
// 		output = append(output, int(b.Int64()))
// 	}
// 	return output
// }
//

// func pow(x, y int) int {
// 	return int(math.Pow(float64(x), float64(y)))
// }

// func whatAs(b, c int) (alist []int) {
// 	for a := c * pow(2, b); a < c*pow(2, b+1); a++ {
// 		alist = append(alist, a)
// 	}
// 	return
// }

// func whatBs(newB int) (blist []int) {
// 	for b := 0; b < 8; b++ {
// 		for c := 0; c < 8; c++ {
// 			ans := b ^ c
// 			if ans == newB {
// 				bb := b ^ 6
// 				alist := whatAs(bb, c)
// 				bbb := bb ^ 5
// 				for _, a := range alist {
// 					if a%8 == bbb {
// 						blist = append(blist, bbb)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return
// }

func main() {
	// maxA := int64(math.Pow(8,15)-1)
	// minA := int64(math.Pow(8,14) + math.Pow(8,13))

	// for i := minA; i <= maxA; i++ {
	// 	if i % minA == 0 {
	// 		fmt.Printf("%d\n", i)
	// 	}
	// }

	// fmt.Println("finished")

	// var zstr string
	// for a := int64(105553116266496); zstr != "2,4,1,5,7,5,1,6,0,3,4,1,5,5,3,0"; a++ {
	// 	z := runProgram(big.NewInt(a))
	// 	zstrs := []string{}
	// 	for i := 0; i < len(z); i++ {
	// 		zstrs = append(zstrs, strconv.Itoa(z[i]))
	// 	}
	// 	zstr = strings.Join(zstrs, ",")
	// }
	// fmt.Println(zstr)

	// // pexample := readInput(example)
	// // pexample2 := readInput(example2)
	// pinput := readInput(input)

	// // fmt.Printf("example: part1: %s\n", pexample.part1())
	// // fmt.Printf("input: part1: %s\n", pinput.part1())
	// // fmt.Printf("example2: part2: %d\n", pexample2.part2())
	// fmt.Printf("input: part1: %s\n", readInput(input).part1Direct())
	// fmt.Printf("input: part2: %d\n", readInput(input).part2())

	// m := make(map[int]int)

	// for a := 0; a <= 7; a++ {
	// 	b := a ^ 5
	// 	c := a / int(math.Pow(2.0, float64(b)))
	// 	b = b ^ 6
	// 	b = b ^ c
	// 	b = b % 8

	// 	fmt.Printf("a=%d: %d <--> %s\n", a, b, readInput(input).part2(a))

	// 	strconv.Ato
	// 	// m[a] = b
	// }

	// fmt.Println(m)

	// for k, v := range m {
	// }
	//
	// for i := 0; i <= 7; i++ {
	// 	blist := whatBs(i)
	// 	fmt.Printf("to make %d you need %v\n", i, blist)
	// }
	//

	// k := 0b10000
	// b := 0b10101010101
	// c := b / k
	// fmt.Printf("%0b / %0b = %0b\n", b, k, c)
	//
	// 001,

	// a := 0b1000001 // 0b1001
	// fmt.Println(a)
	// fmt.Printf("input: part2(a): %s\n", readInput(input).part2(a))

	inp := readInput(input)

	possibleAs := []int{}

	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7; j++ {
			for k := 0; k <= 7; k++ {
				a := i<<6 + j<<3 + k
				res := inp.part2(a)
				if res[len(res)-1] == inp.program[len(inp.program)-1] {
					possibleAs = append(possibleAs, a)
				}
			}
		}
	}

	// fmt.Println(possibleAs)

	for rindex := 2; rindex <= 14; rindex++ {
		possibleAs = inp.solve(possibleAs, rindex)
	}
	//for count := 0; count < 16-3; count++ {
	// for _, p := range possibleAs {
	// 	for l := 0; l <= 7; l++ {
	// 		a := p<<3 + l
	// 		res := inp.part2(a)
	// 		if res[len(res)-2] == inp.program[len(inp.program)-2] {
	// 			fmt.Printf("%d %v\n", a, res[len(res)-2:])
	// 			nextPossibleAs = append(nextPossibleAs, a)
	// 		}
	// 	}
	// }

	// fmt.Println(nextPossibleAs)

	soln := math.MaxInt
loop:
	for _, a := range possibleAs {
		res := inp.part2(a)
		for i := 0; i < len(res); i++ {
			if res[i] != inp.program[i] {
				continue loop
			}
		}
		if a < soln {
			soln = a
		}
	}

	fmt.Printf("soln=%d res=%v\n", soln, inp.part2(soln))

}

func (pz puzzle) solve(possibleAs []int, rindex int) []int {
	nextPossibleAs := []int{}
	for _, p := range possibleAs {
		for l := 0; l <= 7; l++ {
			a := p<<3 + l
			res := pz.part2(a)
			if res[len(res)-rindex] == pz.program[len(pz.program)-rindex] {
				nextPossibleAs = append(nextPossibleAs, a)
			}
		}
	}
	return nextPossibleAs
}

func (p puzzle) part2(aa int) []int {
	result := []int{}
	a, b, c := aa, 0, 0
	for i := 0; ; i++ {
		b = a % 8
		b = b ^ 5
		c = a >> b
		b = b ^ 6
		a = a >> 3
		b = b ^ c
		output := b % 8
		// if output != p.program[i] {
		// 	fmt.Println()
		// 	continue aaloop
		// }
		//	fmt.Printf("%d, ", output)
		result = append(result, output)
		if a == 0 {
			break
		}
	}
	return result
}
