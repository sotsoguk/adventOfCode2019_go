package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/imageutils"
	"github.com/adventOfCode2019_go/utils/intcode"
)

func getColor(c complex64, grid map[complex64]int) int {
	if grid[c] == 1 {
		return 1
	} else {
		return 0
	}
}

func setColor(c complex64, grid map[complex64]int, color int) {
	grid[c] = color
}
func turnRight(d *complex64) {
	*d *= complex(0, -1)
}
func turnLeft(d *complex64) {
	*d *= complex(0, 1)
}
func doStep(c *complex64, d *complex64) {
	*c += *d
}
func partX(part2 bool, code []int64) (int64, [][]int) {

	grid := make(map[complex64]int)
	var currPos complex64 = complex(0, 0)
	var dir complex64 = complex(0, 1)
	var robot intcode.VM
	if part2 {
		setColor(currPos, grid, 1)
	}
	robot.LoadCode(code)
	robot.Reset()
	for robot.Mode != 99 {

		robot.LoadInput(int64(getColor(currPos, grid)))
		robot.RunCode()
		colorO := int(robot.Output[len(robot.Output)-2])
		cDir := robot.Output[len(robot.Output)-1]
		setColor(currPos, grid, colorO)
		if cDir == 1 {
			turnRight(&dir)
		} else {
			turnLeft(&dir)
		}
		doStep(&currPos, &dir)

	}
	return int64(len(grid)), imageutils.ConvertMapToGrid(grid)
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 11
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	ll := strings.Split(lines[0], ",")
	code := make([]int64, len(ll))
	for i := range ll {
		code[i], _ = strconv.ParseInt(ll[i], 10, 64)
	}

	solution1, grid1 := partX(false, code)
	solution2, grid2 := partX(true, code)
	if output {
		imageutils.PrintGrid(grid1)
	}
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)
	imageutils.PrintGrid(grid2)
	//imageutils.RenderGrid("day11_01.png", grid1, 10)
}
