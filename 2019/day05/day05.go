package main

import (
	"fmt"
	"strconv"
	"strings"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func makeDigits(n int) []int {
	digits := make([]int, 5)
	for i := 0; i < 5; i++ {
		digits[4-i] = int(n % 10)
		// digits = append(digits, int(n%10))
		n /= 10
	}
	return digits
}
func runCode(origCode []int, input int) []int {

	code := make([]int, len(origCode))
	copy(code, origCode)
	result := make([]int, 0)
	var opcode, m1, m2 int
	ptr := 0
	running := true
	for running {
		op := makeDigits(code[ptr])
		m1, m2 = op[2], op[1]
		p1 := 0
		p2 := 0
		opcode = 10*op[3] + op[4]
		if opcode == 1 || opcode == 2 || (opcode >= 4 && opcode <= 8) {
			if m1 == 0 {
				p1 = code[code[ptr+1]]
			} else {
				p1 = code[ptr+1]
			}
		}
		if opcode == 1 || opcode == 2 || (opcode > 4 && opcode <= 8) {
			if m2 == 0 {
				p2 = code[code[ptr+2]]
			} else {
				p2 = code[ptr+2]
			}

		}
		switch opcode {
		case 1:

			code[code[ptr+3]] = p1 + p2
			ptr += 4
		case 2:

			code[code[ptr+3]] = p1 * p2
			ptr += 4
		case 3:
			code[code[ptr+1]] = input
			ptr += 2
		case 4:

			fmt.Print(p1, ",")
			result = append(result, p1)
			ptr += 2
		case 5:
			if p1 != 0 {
				ptr = p2
			} else {
				ptr += 3
			}
		case 6:
			if p1 == 0 {
				ptr = p2
			} else {
				ptr += 3
			}
		case 7:
			if p1 < p2 {
				code[code[ptr+3]] = 1
			} else {
				code[code[ptr+3]] = 0
			}
			ptr += 4
		case 8:
			if p1 == p2 {
				code[code[ptr+3]] = 1
			} else {
				code[code[ptr+3]] = 0
			}
			ptr += 4
		case 99:
			running = false
		default:
			fmt.Println("ERRROR:", opcode)
			break
		}

	}
	fmt.Println()
	return result
}

func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input02_2019.txt")
	// fmt.Println(os.Getwd())

	const (
		year = 2019
		day  = 5
		goal = 19690720
	)

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)

	var (
		solution1, solution2 int64
	)

	// prepare input
	// code8 := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	ll := strings.Split(lines[0], ",")
	// ll := strings.Split(code8, ",")

	code := make([]int, len(ll))
	for i := range ll {
		code[i], _ = strconv.Atoi(ll[i])
	}

	part1 := runCode(code, 1)
	part2 := runCode(code, 5)
	for _, i := range part1 {
		if i != 0 {
			solution1 = int64(i)
			break
		}
	}
	for _, i := range part2 {
		if i != 0 {
			solution2 = int64(i)
			break
		}
	}
	fmt.Println(part1, part2)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v",
		header, len(lines), solution1, solution2)
}
