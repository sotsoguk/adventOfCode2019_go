package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/intcode"
)

func PrintGrid19(input [][]int) {
	rows, cols := len(input), len(input[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			pixel := input[y][x]
			switch pixel {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")

			}
		}
		fmt.Println()
	}
}

func part1(code []int64) int64 {

	var solution1 int64
	const (
		rows = 50
		cols = 50
	)
	var robot intcode.VM
	robot.LoadCode(code)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			robot.Reset()
			robot.LoadInputs([]int64{int64(x), int64(y)})
			robot.RunCode()
			if robot.Output[0] == 1 {
				solution1++
			}

		}
	}
	return solution1
}
func part2(code []int64) int64 {
	// 1. calculate the slopes and compute a (float) solution.
	// Upper line: y = cx + d
	// Lower line: y = ax + b
	// Realized: b,d = 0 .....

	// 2. Use the calculated solution as a starting point for looking for the correct spot

	var xs, ys, curr, prev int
	var a, c float64

	var robot intcode.VM
	robot.LoadCode(code)
	xs = 1000
	foundSlopes := false
	for !foundSlopes {
		prev = curr
		robot.Reset()
		robot.LoadInputs([]int64{int64(xs), int64(ys)})
		robot.RunCode()
		curr = int(robot.Output[0])
		if curr == 1 && prev == 0 {
			c = float64(ys) / 1000
		}
		if curr == 0 && prev == 1 {
			a = float64(ys) / 1000
			foundSlopes = true
		}
		ys++
	}

	//calc analytic solution
	xCt := (99*c + 99) / (a - c)
	yC := int(c * (xCt + 99))
	xC := int(xCt)

	// margin for int / float errors
	xs, ys = xC-10, yC-10

	for y := ys; y < ys+1000; y++ {
		for x := xs; x < xs+1000; x++ {
			robot.Reset()
			robot.LoadInputs([]int64{int64(x), int64(y)})
			robot.RunCode()
			out := int(robot.Output[0])

			if out == 1 {
				//trCorner
				robot.Reset()
				robot.LoadInputs([]int64{int64(x + 99), int64(y)})
				robot.RunCode()
				tr := int(robot.Output[0])
				if tr != 1 {
					continue
				}
				//leftlower
				robot.Reset()
				robot.LoadInputs([]int64{int64(x), int64(y + 99)})
				robot.RunCode()
				ll := int(robot.Output[0])
				if ll == 1 {
					return int64(x*10000 + y)
				}
			}
		}
	}
	return 0
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()
	const (
		year   = 2019
		day    = 19
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	// filePath := fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	ll := strings.Split(lines[0], ",")
	code := make([]int64, len(ll))
	for i := range ll {
		code[i], _ = strconv.ParseInt(ll[i], 10, 64)
	}

	solution1 = part1(code)
	solution2 = part2(code)
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)

}
