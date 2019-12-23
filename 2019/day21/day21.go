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

type Vec2i struct {
	x int
	y int
}

func (v Vec2i) add(w Vec2i) Vec2i {
	return Vec2i{v.x + w.x, v.y + w.y}
}
func getColor(c complex64, grid map[complex64]int) int {
	if grid[c] == 1 {
		return 1
	} else {
		return 0
	}
}
func PrintGrid17(input [][]int) {
	rows, cols := len(input), len(input[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			pixel := input[y][x]
			switch pixel {
			case 1:
				fmt.Print(".")
			case 2:
				fmt.Print("#")
			case 3:
				fmt.Print("O")
			case 11:
				fmt.Print("^")
			case 12:
				fmt.Print(">")
			case 13:
				fmt.Print("v")
			case 14:
				fmt.Print("<")
			}
		}
		fmt.Println()
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
func camera2grid(output *[]int64, skipFinalCR bool) [][]int {
	// compute dimensions
	conversionMap := map[int64]int{46: 1, 35: 2, 94: 11}
	var width, height int
	for i := 0; i < len(*output); i++ {
		if (*output)[i] == 10 {
			width = i
			break
		}
	}
	for i := 0; i < len(*output); i++ {
		if (*output)[i] == 10 {
			height++
		}
	}
	if skipFinalCR {
		height--
	}

	// convert
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	var row, col int
	for i := 0; i < len(*output); i++ {
		if (*output)[i] == 10 {
			row++
			col = 0
			continue
		} else {
			grid[row][col] = conversionMap[(*output)[i]]
			col++
		}
	}
	// fmt.Println(width, height)
	return grid
}

func byte2int(input []byte) []int64 {
	output := make([]int64, len(input))
	for i := range input {
		output[i] = int64(input[i])
	}
	return output
}
func cmds2Byte(cmds []string, part2 bool) []byte {
	bytes := make([]byte, 0)
	for _, c := range cmds {
		bytes = append(bytes, []byte(c)...)
		bytes = append(bytes, 10)
	}
	if part2 {
		bytes = append(bytes, []byte("RUN")...)
		bytes = append(bytes, 10)
	} else {
		bytes = append(bytes, []byte("WALK")...)
		bytes = append(bytes, 10)
	}
	return bytes
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 21
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

	//solution1, grid1 := partX(false, code)
	//solution2, grid2 := partX(true, code)
	//if output {
	// imageutils.PrintGrid(grid1)
	//}
	var robot intcode.VM
	cmd1 := "NOT D J"
	cmd2 := "WALK"
	c1 := []byte(cmd1)
	c1 = append(c1, 10)
	c1 = append(c1, []byte(cmd2)...)
	c1 = append(c1, 10)
	// cs := []string{"OR A J", "AND C J", "NOT J J", "AND D J"}
	cs2 := []string{"NOT H J", "OR C J", "AND B J", "AND A J", "NOT J J", "AND D J"}
	// cby := cmds2Byte(cs, false)
	cby2 := cmds2Byte(cs2, true)
	robot.LoadCode(code)
	robot.Reset()
	robot.LoadInputs(byte2int(cby2))
	robot.RunCode()
	bout := make([]byte, len(robot.Output))
	for i := range robot.Output {
		bout[i] = byte(robot.Output[i])
	}
	fmt.Println(string(bout))
	fmt.Println(robot.Output)
	// outputcamera := camera2grid(&robot.Output, true)
	// PrintGrid17(output)
	// fmt.Println(robot.Output)
	// asciiBytes := make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	// outputcamera := camera2grid(&robot.Output, true)
	// fmt.Println(findPath(outputcamera))
	// findIntersections(outputcamera)
	// PrintGrid17(outputcamera)

	// a := "R,12,L,10,L,4,L,6"
	// b := "L,6,R,12,L,6"
	// c := "L,10,L,10,L,4,L,6"
	// prog := "B,A,B,A,B,C,A,C,B,C"
	// progBytes := []byte(prog)
	// progBytes = append(progBytes, 10)
	// aBytes := []byte(a)
	// bBytes := []byte(b)
	// cBytes := []byte(c)
	// aBytes = append(aBytes, 10)
	// bBytes = append(bBytes, 10)
	// cBytes = append(cBytes, 10)

	// fmt.Println(aBytes)
	// fmt.Println(bBytes)
	// fmt.Println(cBytes)
	// fmt.Println(progBytes)
	// // part2
	// code2 := make([]int64, len(code))

	// copy(code2, code)
	// code2[0] = 2
	// robot.LoadCode(code2)
	// robot.Reset()
	// robot.RunCode()
	// asciiBytes := make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	// fmt.Println(robot.Mode)
	// robot.ClearOuput()
	// robot.LoadInputs(byte2int(progBytes))
	// fmt.Println("Loaded inputs")
	// robot.RunCode()
	// asciiBytes = make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	// // Load function A
	// robot.LoadInputs(byte2int(aBytes))
	// robot.ClearOuput()
	// robot.RunCode()
	// asciiBytes = make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	// fmt.Println(robot.Mode)

	// // Load b
	// robot.LoadInputs(byte2int(bBytes))
	// robot.ClearOuput()
	// robot.RunCode()
	// asciiBytes = make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	// fmt.Println(robot.Mode)

	// // load C
	// robot.LoadInputs(byte2int(cBytes))
	// robot.ClearOuput()
	// robot.RunCode()
	// asciiBytes = make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	// fmt.Println(robot.Mode)
	// //no feed
	// robot.LoadInputs([]int64{110, 10})
	// robot.ClearOuput()
	// robot.RunCode()
	// asciiBytes = make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	// fmt.Println(robot.Mode)
	// fmt.Println(robot.Output)
	// fmt.Println(byte2int(aBytes))
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)
	// imageutils.PrintGrid(grid2)
	//imageutils.RenderGrid("day11_01.png", grid1, 10)
}
