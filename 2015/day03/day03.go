package main

import (
	"fmt"
	// "strconv"
	// "strings"

	// "github.com/adventOfCode2019_go/utils/mathUtils"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func main() {
	lines := readAOC.ReadInput("2015/inputs/input03_2015.txt")
	var (
		currPos              complex128 = complex(0, 0) // set starting position
		currPosRobo          complex128 = complex(0, 0) // set starting position for RoboSanta
		solution1, solution2 int64
	)
	fmt.Println(len(lines))

	grid := make(map[complex128]int)

	grid[currPos] = 1

	for _, instr := range lines[0] {
		var dir complex128

		switch instr {
		case '<':
			dir = complex(-1, 0)
		case '>':
			dir = complex(1, 0)
		case 'v':
			dir = complex(0, -1)
		case '^':
			dir = complex(0, 1)
		}
		// update position
		currPos += dir
		grid[currPos]++

	}
	solution1 = int64(len(grid))

	// Part 2
	currPos = complex(0, 0)
	grid2 := make(map[complex128]int)
	for i, instr := range lines[0] {
		var dir complex128
		// compute complex number according to direction instruction
		switch instr {
		case '<':
			dir = complex(-1, 0)
		case '>':
			dir = complex(1, 0)
		case 'v':
			dir = complex(0, -1)
		case '^':
			dir = complex(0, 1)
		}
		// update position
		if i%2 == 0 {
			currPos += dir
			grid2[currPos]++
		} else {
			currPosRobo += dir
			grid2[currPosRobo]++
		}

	}
	solution2 = int64(len(grid2))

	// for k, v := range grid {
	// 	fmt.Println(k, v)
	// }
	// fmt.Print(len(grid))
	fmt.Printf("AoC 2015 - Day 03\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
