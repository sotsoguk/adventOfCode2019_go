package main

import (
	"fmt"
	"strconv"

	// "strconv"
	// "strings"

	// "github.com/adventOfCode2019_go/utils/mathUtils"
	"regexp"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type instr int

const (
	turnOn instr = iota
	turnOff
	toggle
)

type coord struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

var regs = regexp.MustCompile(`^(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)$`)

func create2DsliceBool(xSize, ySize uint) [][]bool {
	grid := make([][]bool, ySize)
	allGrid := make([]bool, xSize*ySize)
	for i := range grid {
		grid[i], allGrid = allGrid[:xSize], allGrid[xSize:]
	}
	return grid
}
func create2DsliceInt(xSize, ySize uint) [][]int {
	grid := make([][]int, ySize)
	allGrid := make([]int, xSize*ySize)
	for i := range grid {
		grid[i], allGrid = allGrid[:xSize], allGrid[xSize:]
	}
	return grid
}
func countLightsOn(grid [][]bool) (numLights int64) {
	for _, row := range grid {
		for _, e := range row {
			if e {
				numLights++
			}
		}
	}
	return
}
func countLightsIntensity(grid [][]int) (intensity int64) {
	for _, row := range grid {
		for _, e := range row {
			intensity += int64(e)
		}
	}
	return
}
func parseInstruction(s string) (instr, coord) {
	m := regs.FindStringSubmatch(s)
	var (
		inst instr
		c    coord
	)
	switch m[1] {
	case "turn on":
		inst = turnOn
	case "turn off":
		inst = turnOff
	case "toggle":
		inst = toggle
	}
	c.x0, _ = strconv.Atoi(m[2])
	c.y0, _ = strconv.Atoi(m[3])
	c.x1, _ = strconv.Atoi(m[4])
	c.y1, _ = strconv.Atoi(m[5])

	return inst, c
}

func execInstruction(grid [][]bool, inst instr, c coord) {
	for x := c.x0; x <= c.x1; x++ {
		for y := c.y0; y <= c.y1; y++ {
			switch inst {
			case turnOn:
				grid[x][y] = true
			case turnOff:
				grid[x][y] = false
			case toggle:
				grid[x][y] = !grid[x][y]
			}
		}
	}
}
func execInstruction2(grid [][]int, inst instr, c coord) {
	for x := c.x0; x <= c.x1; x++ {
		for y := c.y0; y <= c.y1; y++ {
			switch inst {
			case turnOn:
				grid[x][y] += 1
			case turnOff:
				grid[x][y] = int(mathUtils.Max(int64(grid[x][y]-1), 0))
			case toggle:
				grid[x][y] += 2
			}
		}
	}
}
func main() {
	lines := readAOC.ReadInput("2015/inputs/input06_2015.txt")
	var (
		solution1, solution2 int64
	)
	fmt.Println(len(lines))
	lights1 := create2DsliceBool(1000, 1000)
	lights2 := create2DsliceInt(1000, 1000)

	for _, line := range lines {
		i, c := parseInstruction(line)
		execInstruction(lights1, i, c)
		execInstruction2(lights2, i, c)
	}
	solution1 = countLightsOn(lights1)
	solution2 = countLightsIntensity(lights2)
	fmt.Printf("AoC 2015 - Day 06\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
