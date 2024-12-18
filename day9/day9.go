package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	process("sample.txt")
	process("input.txt")
}

func process(filename string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := []byte(string(b))

	// fmt.Printf("%s part1: %d\n", filename, part1(s))
	fmt.Printf("%s part2: %d\n", filename, part2(s))
}

func part1(line []byte) (result int) {
	var disk []int

	id := 0
	for i := 0; i < len(line); i += 2 {
		filelen := mustInt(string(line[i]))
		for j := 0; j < filelen; j++ {
			disk = append(disk, id)
		}
		id++
		if i+1 < len(line) {
			freelen := mustInt(string(line[i+1]))
			for j := 0; j < freelen; j++ {
				disk = append(disk, -1)
			}
		}
	}

	// dump(disk)

	for i, j := 0, len(disk)-1; i < len(disk) && j >= 0 && i < j; i++ {
		if disk[i] == -1 {
			disk[i] = disk[j]
			disk[j] = -1
			for disk[j] == -1 {
				j--
			}
		}
	}

	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			checksum += i * disk[i]
		}
	}

	return checksum
}

type block struct {
	id   int
	size int
	pos  int
}

func part2(line []byte) (result int) {
	files := []block{}
	free := []block{}

	pos := 0
	id := 0
	for i := 0; i < len(line); i += 2 {
		filelen := mustInt(string(line[i]))
		files = append(files, block{id: id, size: filelen, pos: pos})
		pos += filelen
		id++
		if i+1 < len(line) {
			freelen := mustInt(string(line[i+1]))
			free = append(free, block{size: freelen, pos: pos})
			pos += freelen
		}
	}

	// dump2(files)
	// dump2(free)

	for i := len(files) - 1; i >= 0; i-- {
		for j := 0; j < len(free); j++ {
			if free[j].pos < files[i].pos && free[j].size >= files[i].size {
				// fmt.Printf("moving %d (len %d) from %d to %d\n", files[i].id, files[i].size, files[i].pos, free[j].pos)
				files[i].pos = free[j].pos
				free[j].size -= files[i].size
				free[j].pos += files[i].size
			}
		}
		// dump2(files)
		// dump2(free)
	}

	checksum := 0
	for i := 0; i < len(files); i++ {
		for j := 0; j < files[i].size; j++ {
			checksum += (files[i].pos + j) * files[i].id
		}
	}

	return checksum
}

func dump(disk []int) {
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%s", strconv.Itoa(disk[i]))
		}
	}
	fmt.Println()
}

func dump2(files []block) {
	disk := [100]string{}
	for i := 0; i < 100; i++ {
		disk[i] = "."
	}
	for _, file := range files {
		for i := 0; i < file.size; i++ {
			disk[file.pos+i] = strconv.Itoa(file.id)
		}
	}
	for i := 0; i < len(disk); i++ {
		fmt.Printf("%s", disk[i])
	}
	fmt.Println()
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
