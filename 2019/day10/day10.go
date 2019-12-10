package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type Slope struct {
	dx       int
	dy       int
	negative bool
}
type Pos struct {
	x int
	y int
}
type Coord struct {
	x int
	y int
}

func (p Coord) distance(q Coord) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))

}
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

func (s *Slope) shorten() {
	gcds := mathUtils.Abs32(mathUtils.Gcd(s.dx, s.dy))
	s.dx /= gcds
	s.dy /= gcds
}
func makeGrid(input []string) [][]int {
	rows, cols := len(input), len(input)
	spaceMap := make([][]int, rows)
	for i := 0; i < rows; i++ {
		spaceMap[i] = make([]int, cols)
	}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if input[y][x] == 35 {
				spaceMap[y][x] = 1
			}
		}
	}
	return spaceMap
}

func printGrid(input [][]int) {
	rows, cols := len(input), len(input[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if input[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func countAsteroids(grid [][]int, posX, posY int) int {
	rows, cols := len(grid), len(grid[0])
	dirs := make(map[Slope]int)
	checked := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 1 && !(posX == x && posY == y) {
				var dir Slope
				dir.dx = x - posX
				dir.dy = y - posY
				(&dir).shorten()
				dirs[dir]++
				// fmt.Println(x, y, dir)
				checked++
			}
		}
	}
	return len(dirs)
}
func countAsteroids2(grid [][]int, posX, posY int) map[Slope][]Coord {
	rows, cols := len(grid), len(grid[0])
	dirs := make(map[Slope][]Coord)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 1 && !(posX == x && posY == y) {
				var dir Slope
				dir.dx = x - posX
				dir.dy = y - posY
				(&dir).shorten()
				newPos := Coord{x, y}
				dirs[dir] = append(dirs[dir], newPos)
			}
		}
	}
	return dirs
}
func (s *Slope) toAngle() float64 {

	angle := math.Atan2(float64(s.dy), float64(s.dx))/(2*math.Pi)*360.0 + 90
	if angle < 0 {
		angle += 360
	}
	return angle
}

func part1(grid [][]int) (int, Coord) {
	maxAsteroids := 0
	var maxX, maxY int
	rows, cols := len(grid), len(grid[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 1 {
				currAsteroids := countAsteroids(grid, x, y)
				if currAsteroids > maxAsteroids {
					maxAsteroids = currAsteroids
					maxX = x
					maxY = y
				}
			}
		}
	}
	return maxAsteroids, Coord{maxX, maxY}
}
func part2(turret Coord, grid [][]int, target int) {
	slopeAsteroidsMap := countAsteroids2(grid, turret.x, turret.y)
	anglesAsteroidsMap := make(map[float64][]Coord)
	angles := make([]float64, 0)
	for slope := range slopeAsteroidsMap {
		angle := (&slope).toAngle()
		anglesAsteroidsMap[angle] = slopeAsteroidsMap[slope]
		sort.Slice(anglesAsteroidsMap[angle], func(p, q int) bool {
			return anglesAsteroidsMap[angle][p].distance(turret) < anglesAsteroidsMap[angle][q].distance(turret)
		})
		if len(anglesAsteroidsMap[angle]) >= 1 {
			angles = append(angles, angle)
		}
		sort.Float64s(angles)
	}
	i := 0
	for i < target {
		for _, a := range angles {
			if len(anglesAsteroidsMap[a]) > 0 {
				i++
				if i == target {
					fmt.Println(anglesAsteroidsMap[a][0])
					return
				}
				anglesAsteroidsMap[a] = anglesAsteroidsMap[a][1:]
			}
		}
	}
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year = 2019
		day  = 10
	)

	// filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	filePath := fmt.Sprintf("%d/inputs/input10_3.txt", year)

	lines := readAOC.ReadInput(filePath)
	// fmt.Println(lines)
	var (
		solution1, solution2 int64
	)
	sm := makeGrid(lines)
	// fmt.Println(sm)
	// printGrid(sm)
	ll := strings.Split(lines[0], ",")
	code := make([]int64, len(ll))
	for i := range ll {
		code[i], _ = strconv.ParseInt(ll[i], 10, 64)
	}
	// fmt.Println(mathUtils.Gcd(10, 0))
	elapsed := time.Since(start)
	// ts := Slope{dx: -100, dy: 8}
	// ts2 := Slope{dx: 200, dy: -16}
	// (&ts).shorten()
	// (&ts2).shorten()
	// fmt.Println(ts)
	// checkA := make(map[Slope]int, 0)
	// checkA[ts]++
	// checkA[ts2]++
	// fmt.Println(len(checkA))
	// fmt.Println(countAsteroids(sm, 6, 3))
	num1, maxPos := part1(sm)
	fmt.Println(num1, maxPos)
	// fmt.Println("P1:", num1, maxPos)
	// ts := Slope{dx: -20, dy: -1}
	// fmt.Println((&ts).toAngle())
	// fmt.Println(math.Atan2(float64(-ts.dy), float64(-ts.dx)))
	// iv := int(math.Atan2(float64(ts.dy), float64(ts.dx))/(2*math.Pi)*360.0 + 90)
	// fmt.Println((&ts).toAngle())
	slopeAsteroidsMap := countAsteroids2(sm, maxPos.x, maxPos.y)
	fmt.Println("POS:", maxPos)
	angleSlopeMap := make(map[float64]Slope)
	angles := make([]float64, 0)
	anglesNumMap := make(map[float64]int)
	for slope := range slopeAsteroidsMap {
		angle := (&slope).toAngle()

		angleSlopeMap[angle] = slope
		anglesNumMap[angle] += len(slopeAsteroidsMap[slope])
		if anglesNumMap[angle] >= 1 {
			angles = append(angles, angle)
		}
	}
	anglesAsteroidsMap := make(map[float64][]Coord)
	for slope := range slopeAsteroidsMap {
		angle := (&slope).toAngle()
		anglesAsteroidsMap[angle] = slopeAsteroidsMap[slope]
		sort.Slice(anglesAsteroidsMap[angle], func(p, q int) bool {
			return anglesAsteroidsMap[angle][p].distance(maxPos) < anglesAsteroidsMap[angle][q].distance(maxPos)
		})
	}

	sort.Float64s(angles)
	// // fmt.Println(anglesNumMap)
	totals := 0
	for _, v := range anglesNumMap {
		totals += v
	}
	fmt.Println(totals)
	// for _, a := range angles {
	// 	fmt.Print(a, ":")
	// 	if anglesNumMap[a] > 0 {
	// 		fmt.Print(anglesNumMap[a], ":")
	// 		fmt.Print(slopeAsteroidsMap[angleSlopeMap[a]])
	// 	}
	// 	fmt.Println()
	// }
	upTo := 299
	round := 0

	i := 0
	for i < upTo {
		for _, a := range angles {
			if anglesNumMap[a] > 0 {
				if i == upTo {
					fmt.Println("BOOOM: ", upTo, a, round, slopeAsteroidsMap[angleSlopeMap[a]])
				}
				// asteroid found in that direction
				fmt.Println(i, a, slopeAsteroidsMap[angleSlopeMap[a]])
				anglesNumMap[a]--
				// fmt.Println(angle, i)
				// fmt.Println(anglesNumMap)
				i++
			}

		}
		round++
	}
	// fmt.Println(angle, round)
	// fmt.Println(slopeAsteroidsMap[angleSlopeMap[angle]])
	// fmt.Println(mathUtils.Gcd(-4, -2))
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v",
		header, len(lines), solution1, solution2, elapsed)

	// testDist := make([]Coord, 4)
	// testDist[3] = Coord{1, 1}
	// testDist[2] = Coord{1, 2}
	// testDist[1] = Coord{1, 3}
	// testDist[0] = Coord{2, 4}
	// fmt.Println(testDist)
	// sort.Slice(testDist, func(p, q int) bool {
	// 	return testDist[p].distance(Coord{0, 0}) < testDist[q].distance(Coord{0, 0})
	// })
	// fmt.Println(testDist)

	part2(maxPos, sm, 200)
}
