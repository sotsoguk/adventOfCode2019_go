package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func runCode(origCode []int, noun int, verb int) []int {

	code := make([]int, len(origCode))
	copy(code, origCode)
	code[1] = noun
	code[2] = verb
	ptr := 0
	running := true
	for running {
		switch code[ptr] {
		case 1:
			code[code[ptr+3]] = code[code[ptr+2]] + code[code[ptr+1]]
			ptr += 4
		case 2:
			code[code[ptr+3]] = code[code[ptr+2]] * code[code[ptr+1]]
			ptr += 4
		case 99:
			running = false
		}
	}
	return code
}

func main() {
	fmt.Println(os.Getwd())
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input02_2019.txt")
	lines := readAOC.ReadInput("2019/inputs/input02_2019.txt")
	var (
		solution1, solution2 int64
	)
	const (
		goal = 19690720
	)

	// prepare input
	ll := strings.Split(lines[0], ",")
	code := make([]int, len(ll))
	for i := range ll {
		code[i], _ = strconv.Atoi(ll[i])
	}

	// Part 1
	codePart1 := runCode(code, 12, 2)
	solution1 = int64(codePart1[0])

	// Part 2
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			memCpy := runCode(code, n, v)
			if memCpy[0] == goal {
				fmt.Println("Part2: Found the solution", n, v)
				solution2 = int64(100*n + v)
				break
			}
		}
	}
	fmt.Printf("AoC 2019 - Day 02\n-----------------\nLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v",
		len(lines), solution1, solution2)
}
