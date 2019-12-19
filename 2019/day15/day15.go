package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/intcode"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type Vec2i struct {
	x int
	y int
}

func (v Vec2i) add(w Vec2i) Vec2i {
	return Vec2i{v.x + w.x, v.y + w.y}
}

type dGrid struct {
	data                   [][]int
	xMin, xMax, yMin, yMax int
	width, height          int
}

func newdGrid(w, h, x, y int) dGrid {
	var g dGrid
	g.data = make([][]int, h)
	for i := 0; i < h; i++ {
		g.data[i] = make([]int, w)
	}
	g.xMax, g.xMin = x, x
	g.yMax, g.yMin = y, y
	g.width = w
	g.height = h
	return g
}

// type Command int64

const (
	CUp    int64 = 1
	CDown  int64 = 2
	CLeft  int64 = 3
	CRight int64 = 4
)

var reverseCommand = map[int64]int64{CUp: CDown, CDown: CUp, CLeft: CRight, CRight: CLeft}
var (
	Up    = Vec2i{0, -1}
	Down  = Vec2i{0, 1}
	Left  = Vec2i{-1, 0}
	Right = Vec2i{1, 0}
)

func dgrid2grid(g dGrid) [][]int {
	h, w := g.yMax-g.yMin+1, g.xMax-g.xMin+1
	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			grid[y][x] = g.data[y+g.yMin][x+g.xMin]
		}
	}
	return grid
}
func PrintGrid17(input [][]int) {
	rows, cols := len(input), len(input[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			pixel := input[y][x]
			switch pixel {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print(".")
			case 2:
				fmt.Print("O")
				// fmt.Print("!!!", y, x)
			case 3:
				fmt.Print("#")
			case 4:
				fmt.Print("D")
				// fmt.Print("!!!", y, x)
			case 5:
				fmt.Print("*")
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
func PrintGridNumbers(input [][]int) {
	rows, cols := len(input), len(input[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			pixel := input[y][x]
			if pixel < 10000 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func (g dGrid) at(v Vec2i) int {
	if v.x >= 0 && v.x < g.width && v.y >= 0 && v.y < g.height {
		return g.data[v.y][v.x]
	} else {
		return -1
	}
}
func (g *dGrid) set(v Vec2i, d int) {
	if v.x >= 0 && v.x < g.width && v.y >= 0 && v.y < g.height {
		(*g).data[v.y][v.x] = d
		g.xMin = mathUtils.Min32(g.xMin, v.x)
		g.xMax = mathUtils.Max32(g.xMax, v.x)
		g.yMin = mathUtils.Min32(g.yMin, v.y)
		g.yMax = mathUtils.Max32(g.yMax, v.y)
	}
}
func PrintDGrid(g dGrid) {
	rows, cols := g.yMax-g.yMin+1, g.xMax-g.xMin+1
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			pixel := g.data[y+g.yMin][x+g.xMin]
			switch pixel {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print(".")
			case 2:
				fmt.Print("O")
			case 3:
				fmt.Print("#")
			case 4:
				fmt.Print("D")
			case 5:
				fmt.Print("*")
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

type cmdList []string
type Step struct {
	fromPos Vec2i
	command int64
}

func NewDeque() *Deque {
	return &Deque{}
}

type Deque struct {
	Items []interface{}
}

func (s *Deque) Push(item interface{}) {
	temp := []interface{}{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) Inject(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Deque) Pop() interface{} {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) Eject() interface{} {
	i := len(s.Items) - 1
	defer func() {
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
	}()
	return s.Items[i]
}

func (s *Deque) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

func part1(code []int64) {
	var robot intcode.VM
	robot.LoadCode(code)
	robot.Reset()
	startPos := Vec2i{100, 100}
	pos := Vec2i{100, 100}
	grid := newdGrid(200, 200, pos.x, pos.y)
	// grid.set(pos, 4)
	steps := 0
	path := make([]Step, 0)
	for { //steps <= 100 {
		foundWay := false
		// fmt.Println("Hello")
		// turn Up
		if !foundWay && grid.at(pos.add(Up)) == 0 {
			robot.LoadInput(CUp)
			robot.RunCode()
			out := robot.Output[0]
			// fmt.Println("U", out)
			robot.ClearOuput()
			if out == 0 {
				grid.set(pos.add(Up), 3)
			} else {
				grid.set(pos.add(Up), int(out))
				path = append(path, Step{pos, CUp})
				pos = pos.add(Up)
				foundWay = true
			}
		}
		if !foundWay && grid.at(pos.add(Right)) == 0 {
			robot.LoadInput(CRight)
			robot.RunCode()
			out := robot.Output[0]
			// fmt.Println("R", out)
			robot.ClearOuput()
			if out == 0 {
				grid.set(pos.add(Right), 3)
			} else {
				grid.set(pos.add(Right), int(out))
				path = append(path, Step{pos, CRight})
				pos = pos.add(Right)
				foundWay = true
			}
		}
		if !foundWay && grid.at(pos.add(Down)) == 0 {
			robot.LoadInput(CDown)
			robot.RunCode()
			out := robot.Output[0]
			// fmt.Println("D", out)
			robot.ClearOuput()
			if out == 0 {
				grid.set(pos.add(Down), 3)
			} else {
				grid.set(pos.add(Down), int(out))
				path = append(path, Step{pos, CDown})
				pos = pos.add(Down)
				foundWay = true
			}
		}
		if !foundWay && grid.at(pos.add(Left)) == 0 {
			robot.LoadInput(CLeft)
			robot.RunCode()
			out := robot.Output[0]
			// fmt.Println("L", out)
			robot.ClearOuput()
			if out == 0 {
				grid.set(pos.add(Left), 3)
			} else {
				grid.set(pos.add(Left), int(out))
				path = append(path, Step{pos, CLeft})
				pos = pos.add(Left)
				foundWay = true
			}
		}
		if !foundWay && len(path) == 0 {
			break
		}
		if !foundWay {
			// reached dead end return
			lastStep := path[len(path)-1]
			path = path[:len(path)-1]
			pos = lastStep.fromPos
			robot.LoadInput(reverseCommand[lastStep.command])
			robot.RunCode()
			robot.ClearOuput()
		}
		// PrintDGrid(grid)
		steps++
		// fmt.Println(pos, path)
		// grid.set(pos, 4)
		// PrintDGrid(grid)

	}
	// robot.RunCode()
	grid.set(startPos, 4)
	PrintDGrid(grid)
	// fmt.Println(robot.Mode)
	// fmt.Println(grid.xMin, grid.xMax, grid.yMin, grid.yMax)
	nGrid := dgrid2grid(grid)
	fmt.Println()
	PrintGrid17(nGrid)
	distances := make([][]int, len(nGrid))
	for i := 0; i < len(nGrid); i++ {
		distances[i] = make([]int, len(nGrid[0]))
	}

	for y := 0; y < len(distances); y++ {
		for x := 0; x < len(distances[0]); x++ {
			distances[y][x] = 10000
		}

	}
	// PrintGridNumbers(distances)
	// find shortes path

	deq := NewDeque()
	start := Vec2i{21, 21}
	// goal := Vec2i{39, 1}
	cpos := start
	answer := -1
	distances[cpos.y][cpos.x] = 0
	deq.Push(cpos)
	for !deq.IsEmpty() {
		u := deq.Pop().(Vec2i)
		if nGrid[u.y][u.x] == 2 {
			answer = distances[u.y][u.x]
			break
		}
		//Walk up
		// UP
		v := u
		steps := 0
		for {
			steps++
			v = v.add(Up)
			if nGrid[v.y][v.x] == 3 || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}
		//Right
		v = u
		steps = 0
		for {
			steps++
			v = v.add(Right)
			if nGrid[v.y][v.x] == 3 || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}
		//Down
		v = u
		steps = 0
		for {
			steps++
			v = v.add(Down)
			if nGrid[v.y][v.x] == 3 || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}
		//Left
		v = u
		steps = 0
		for {
			steps++
			v = v.add(Left)
			if nGrid[v.y][v.x] == 3 || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}

	}
	fmt.Println(answer)

	//Part 2
	part2(nGrid, Vec2i{39, 1})

	// show path

}

func part2(grid [][]int, start Vec2i) {
	distances := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		distances[i] = make([]int, len(grid[0]))
	}
	for y := 0; y < len(distances); y++ {
		for x := 0; x < len(distances[0]); x++ {
			distances[y][x] = 10000
		}

	}
	PrintGrid17(grid)
	h, w := len(grid), len(grid[0])
	deq := NewDeque()
	//start := Vec2i{21, 21}
	// goal := Vec2i{39, 1}
	cpos := start
	canWalk := func(pos Vec2i) bool {
		return pos.x >= 0 && pos.y >= 0 && pos.x < w && pos.y < h && grid[pos.y][pos.x] != 3
	}
	distances[cpos.y][cpos.x] = 0
	deq.Push(cpos)
	d := 0
	rnds := 0
	for !deq.IsEmpty() { //&& rnds < 10 {
		rnds++
		u := deq.Pop().(Vec2i)

		v := u

		for {

			v = v.add(Up)
			if !canWalk(v) || distances[v.y][v.x] < 10000 {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + 1
				deq.Inject(v)
				d = mathUtils.Max32(d, distances[u.y][u.x]+1)
				break
			}
		}
		//Right
		v = u

		for {

			v = v.add(Right)
			if !canWalk(v) || distances[v.y][v.x] < 10000 { //<= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + 1
				deq.Inject(v)
				d = mathUtils.Max32(d, distances[u.y][u.x]+1)
				break
			}
		}
		//Down
		v = u

		for {

			v = v.add(Down)
			if !canWalk(v) || distances[v.y][v.x] < 10000 {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + 1
				deq.Inject(v)
				d = mathUtils.Max32(d, distances[u.y][u.x]+1)
				break
			}
		}
		//Left
		v = u

		for {

			v = v.add(Left)
			if !canWalk(v) || distances[v.y][v.x] < 10000 {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + 1
				deq.Inject(v)
				d = mathUtils.Max32(d, distances[u.y][u.x]+1)
				break
			}
		}
		// fmt.Print("\033[H\033[2J")
		grid2 := make([][]int, len(grid))
		for i := 0; i < len(grid); i++ {
			grid2[i] = make([]int, len(grid[0]))
		}
		for y := 0; y < len(distances); y++ {
			for x := 0; x < len(distances[0]); x++ {
				if distances[y][x] != 10000 {
					grid2[y][x] = 5
				} else {
					grid2[y][x] = grid[y][x]
				}
			}
		}
		// fmt.Println("\n After ", d, " minute\n")
		// PrintGrid17(grid2)
	}

	fmt.Println(d)
	// PrintGridNumbers(distances)
	// maxVal := 0
	for y := 0; y < len(distances); y++ {
		for x := 0; x < len(distances[0]); x++ {
			if distances[y][x] != 10000 {
				grid[y][x] = 5
			}
		}
	}
	// PrintGrid17(grid)
	// fmt.Println(maxVal)

}
func findPath(grid [][]int) cmdList {
	var (
		Up    = Vec2i{0, -1}
		Down  = Vec2i{0, 1}
		Left  = Vec2i{-1, 0}
		Right = Vec2i{1, 0}
	)

	var (
		turnLeft  = map[Vec2i]Vec2i{Up: Left, Right: Up, Down: Right, Left: Down}
		turnRight = map[Vec2i]Vec2i{Up: Right, Right: Down, Down: Left, Left: Up}
	)
	// find Robot
	var pos, dir Vec2i
	h, w := len(grid), len(grid[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			switch grid[y][x] {
			case 11:
				pos = Vec2i{x, y}
				dir = Up
			case 12:
				pos = Vec2i{x, y}
				dir = Right
			case 13:
				pos = Vec2i{x, y}
				dir = Down
			case 14:
				pos = Vec2i{x, y}
				dir = Left
			}
		}
	}
	fmt.Println(pos, dir)

	// compute Path
	var path cmdList
	isWalkable := func(pos Vec2i) bool {
		return pos.x >= 0 && pos.x < w && pos.y < h && pos.y >= 0 && grid[pos.y][pos.x] == 2
	}
	for {
		currLength := 0
		for isWalkable(pos.add(dir)) {
			currLength++
			pos = pos.add(dir)
		}
		if currLength != 0 {
			path = append(path, strconv.Itoa(currLength))
		}
		if nextDir := turnLeft[dir]; isWalkable(pos.add(nextDir)) {
			dir = nextDir
			path = append(path, "L")
		} else if nextDir := turnRight[dir]; isWalkable(pos.add(nextDir)) {
			dir = nextDir
			path = append(path, "R")
		} else {
			break
		}
	}
	return path

}

func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 15
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	// filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	filePath := fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)
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
	// var robot intcode.VM
	// robot.LoadCode(code)
	// robot.Reset()
	// robot.RunCode()
	part1(code)
	// testGrid := newdGrid(100, 100, 30, 30)
	// for i := 0; i < 10; i++ {
	// 	testGrid.set(Vec2i{10 + i, 20}, 3)
	// 	testGrid.set(Vec2i{10 + i, 30}, 3)
	// 	testGrid.set(Vec2i{10, 20 + i}, 3)
	// }
	// PrintDGrid(testGrid)
	// fmt.Println(robot.Output)
	// asciiBytes := make([]byte, len(robot.Output))
	// for i := range robot.Output {
	// 	asciiBytes[i] = byte(robot.Output[i])
	// }
	// fmt.Println(string(asciiBytes))
	//outputcamera := camera2grid(&robot.Output, true)
	// fmt.Println(findPath(outputcamera))
	// // findIntersections(outputcamera)
	// PrintGrid17(outputcamera)

	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)

}
