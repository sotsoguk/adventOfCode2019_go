package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func runCode(code []int, noun int, verb int) []int {

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
	originalCode := make([]int, len(ll))
	for i := range ll {
		originalCode[i], _ = strconv.Atoi(ll[i])
	}
	// code := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	// code := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	// Part 1
	codePart1 := make([]int, len(originalCode))
	copy(codePart1, originalCode)
	codePart1 = runCode(codePart1, 12, 2)
	solution1 = int64(codePart1[0])

	// Part 2
	for n := 1; n < 100; n++ {
		for v := 1; v < 100; v++ {
			cc := make([]int, len(originalCode))
			copy(cc, originalCode)
			memCpy := runCode(cc, n, v)
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
