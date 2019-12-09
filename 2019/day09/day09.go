package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

func makeDigits(n int64) []int64 {
	digits := make([]int64, 5)
	for i := 0; i < 5; i++ {
		digits[4-i] = int64(n % 10)
		// digits = append(digits, int(n%10))
		n /= 10
	}
	return digits
}
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
func runCode(origCode []int64, input []int64) []int64 {
	reverseMemory := int64(2000)
	code := make([]int64, mathUtils.Max(int64(len(origCode)), reverseMemory))
	copy(code, origCode)
	result := make([]int64, 0)

	var ptr, inPtr, rPtr int64
	running := true
	printOutput := false

	for running {
		op := makeDigits(code[ptr])
		var offset3 int64
		m1, m2, m3 := op[2], op[1], op[0]
		if m3 == 2 {
			offset3 = rPtr
		}
		var opcode, p1, p2 int64
		opcode = 10*op[3] + op[4]
		if opcode == 1 || opcode == 2 || (opcode >= 4 && opcode <= 9) {
			if m1 == 0 {
				p1 = code[code[ptr+1]]
			} else if m1 == 1 {
				p1 = code[ptr+1]
			} else if m1 == 2 {
				p1 = code[code[ptr+1]+rPtr]
			}
		}
		if opcode == 1 || opcode == 2 || (opcode > 4 && opcode <= 8) {
			if m2 == 0 {
				p2 = code[code[ptr+2]]
			} else if m2 == 1 {
				p2 = code[ptr+2]
			} else if m2 == 2 {
				p2 = code[code[ptr+2]+rPtr]
			}

		}
		switch opcode {
		case 1:

			code[code[ptr+3]+offset3] = p1 + p2
			ptr += 4
		case 2:

			code[code[ptr+3]+offset3] = p1 * p2
			ptr += 4
		case 3:
			if m1 == 0 {
				code[code[ptr+1]+offset3] = input[inPtr]
			} else if m1 == 2 {
				code[code[ptr+1]+rPtr] = input[inPtr]
			}
			ptr += 2
			inPtr++
		case 4:

			if printOutput {
				fmt.Print(p1, ",")
			}
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
				code[code[ptr+3]+offset3] = 1
			} else {
				code[code[ptr+3]+offset3] = 0
			}
			ptr += 4
		case 8:
			if p1 == p2 {
				code[code[ptr+3]+offset3] = 1
			} else {
				code[code[ptr+3]+offset3] = 0
			}
			ptr += 4
		case 9:
			rPtr += p1
			ptr += 2
		case 99:
			running = false
		default:
			fmt.Println("Unknown Opcode:", opcode)
			running = false
			break
		}

	}
	if printOutput {
		fmt.Println()
	}
	return result
}

func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year = 2019
		day  = 9
	)

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)

	var (
		solution1, solution2 int64
	)

	ll := strings.Split(lines[0], ",")
	code := make([]int64, len(ll))
	for i := range ll {
		code[i], _ = strconv.ParseInt(ll[i], 10, 64)
	}

	solution1 = runCode(code, []int64{1})[0]
	solution2 = runCode(code, []int64{2})[0]
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v",
		header, len(lines), solution1, solution2, elapsed)
}
