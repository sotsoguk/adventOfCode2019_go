package main

import (
	"fmt"
	"strconv"
	"strings"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
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
func runCode(origCode []int, input []int) []int {

	code := make([]int, len(origCode))
	copy(code, origCode)
	result := make([]int, 0)
	var opcode, m1, m2 int
	ptr := 0
	inPtr := 0
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
			code[code[ptr+1]] = input[inPtr]
			ptr += 2
			inPtr++
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

func runCode2(code []int, input []int, inPtr int, ptr int) ([]int, []int, int, int, bool, int) {

	// code := make([]int, len(origCode))
	// copy(code, origCode)
	// result := make([]int, 0)
	var opcode, m1, m2 int
	var result int
	var finished bool
	// ptr := 0
	// inPtr := 0
	running := true
	// finished := false
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
			code[code[ptr+1]] = input[inPtr]
			ptr += 2
			inPtr++
		case 4:

			fmt.Print(p1, ",")
			// result = append(result, p1)
			result = p1
			ptr += 2
			running = false
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
			finished = true
		default:
			fmt.Println("ERRROR:", opcode)
			break
		}

	}
	// fmt.Println()
	return code, input, inPtr, ptr, finished, result
}
func runAmplifier(origCode []int, phases []int) int {
	codeA := make([]int, len(origCode))
	codeB := make([]int, len(origCode))
	codeC := make([]int, len(origCode))
	codeD := make([]int, len(origCode))
	codeE := make([]int, len(origCode))
	copy(codeA, origCode)
	copy(codeB, origCode)
	copy(codeC, origCode)
	copy(codeD, origCode)
	copy(codeE, origCode)
	// run Amplifier
	inA := []int{phases[0], 0}
	outA := runCode(codeA, inA)
	inB := []int{phases[1], outA[0]}
	outB := runCode(codeB, inB)
	inC := []int{phases[2], outB[0]}
	outC := runCode(codeC, inC)
	inD := []int{phases[3], outC[0]}
	outD := runCode(codeD, inD)
	inE := []int{phases[4], outD[0]}
	outE := runCode(codeE, inE)
	return outE[0]
}

func runAmplifier2(origCode []int, phases []int) int {
	codeA := make([]int, len(origCode))
	codeB := make([]int, len(origCode))
	codeC := make([]int, len(origCode))
	codeD := make([]int, len(origCode))
	codeE := make([]int, len(origCode))
	copy(codeA, origCode)
	copy(codeB, origCode)
	copy(codeC, origCode)
	copy(codeD, origCode)
	copy(codeE, origCode)
	var (
		ptrA, ptrB, ptrC, ptrD, ptrE           int
		inPtrA, inPtrB, inPtrC, inPtrD, inPtrE int
	)
	inA := []int{phases[0], 0}
	inB := []int{phases[1]}
	inC := []int{phases[2]}
	inD := []int{phases[3]}
	inE := []int{phases[4]}
	// var fA, fB, fC, fD, fE bool
	var fE bool
	var rA, rB, rC, rD, rE int
	var oldRE int
	for !fE {
		oldRE = rE
		codeA, inA, inPtrA, ptrA, _, rA = runCode2(codeA, inA, inPtrA, ptrA)
		inB = append(inB, rA)
		codeB, inB, inPtrB, ptrB, _, rB = runCode2(codeB, inB, inPtrB, ptrB)
		inC = append(inC, rB)
		codeC, inC, inPtrC, ptrC, _, rC = runCode2(codeC, inC, inPtrC, ptrC)
		inD = append(inD, rC)
		codeD, inD, inPtrD, ptrD, _, rD = runCode2(codeD, inD, inPtrD, ptrD)
		inE = append(inE, rD)
		codeE, inE, inPtrE, ptrE, fE, rE = runCode2(codeE, inE, inPtrE, ptrE)
		inA = append(inA, rE)
		// fmt.Println(rA, rB, rC, rD, rE, fA, fE)
	}
	fmt.Println()
	return oldRE
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input02_2019.txt")
	// fmt.Println(os.Getwd())

	const (
		year = 2019
		day  = 7
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
	// ll := strings.Split(lines[0], ",")
	// testCode1 := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	// testCode1 := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	// testCode1 := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	// testCode1 := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
	// testCode1 := "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
	ll := strings.Split(lines[0], ",")
	// ll := strings.Split(testCode1, ",")
	code := make([]int, len(ll))
	for i := range ll {
		code[i], _ = strconv.Atoi(ll[i])
	}

	phases := []int{5, 6, 7, 8, 9}
	// fmt.Println("MAX", runAmplifier2(code, phases))
	phPerm := permutations(phases)
	// // fmt.Println(phPerm)
	maxThrust := 0
	for _, perm := range phPerm {
		currThrust := runAmplifier2(code, perm)
		maxThrust = mathUtils.Max32(currThrust, maxThrust)
	}
	fmt.Println("MAX:", maxThrust)
	//fmt.Println(runAmplifier(code, phases))
	// part1 := runCode(code, 1)
	// part2 := runCode(code, 5)

	// for _, i := range part1 {
	// 	if i != 0 {
	// 		solution1 = int64(i)
	// 		break
	// 	}
	// }
	// for _, i := range part2 {
	// 	if i != 0 {
	// 		solution2 = int64(i)
	// 		break
	// 	}
	// }
	// fmt.Println(part1, part2)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v",
		header, len(lines), solution1, solution2)
}
