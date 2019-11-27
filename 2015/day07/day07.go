package main

import (
	"fmt"

	// "strconv"
	// "strings"

	// "github.com/adventOfCode2019_go/utils/mathUtils"

	readAOC "github.com/adventOfCode2019_go/utils"
)

type instr int

const (
	AND instr = iota
	OR
	NOT
	LSHIFT
	RSHIFT
)

type coord struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

var wiresValues = make(map[string]int)

func main() {
	lines := readAOC.ReadInput("2015/inputs/input07_2015.txt")
	var (
		solution1, solution2 int64
	)
	fmt.Printf("Length of input (in lines): %v\n", len(lines))

	fmt.Printf("AoC 2015 - Day 07\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
