package main

import (
	"fmt"
	"strconv"
	"strings"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type lineOrientation int

const (
	horizontal lineOrientation = 0
	vertical   lineOrientation = 1
)

type point struct {
	x int
	y int
}
type line struct {
	p0      point
	p1      point
	o       lineOrientation
	length  int
	flipped bool
}

func intersectWires(w1 []line, w2 []line) ([]point, []int) {
	intersections := make([]point, 0)
	dists := make([]int, 0)
	dist1 := 0
	for _, a := range w1 {
		dist2 := 0
		for _, b := range w2 {
			if a.o == horizontal && b.o == horizontal && a.p0.y == b.p0.y {
				r, ok := findCommenInterval(a.p0.x, a.p1.x, b.p0.x, b.p1.x)
				if ok {
					for i := r[0]; i <= r[1]; i++ {
						intersections = append(intersections, point{i, a.p0.y})
					}
				}
			} else if a.o == vertical && b.o == vertical && a.p0.x == b.p0.x {
				r, ok := findCommenInterval(a.p0.y, a.p1.y, b.p0.y, b.p1.y)
				if ok {
					for i := r[0]; i <= r[1]; i++ {
						intersections = append(intersections, point{a.p0.x, i})
					}
				}
			} else if a.o == vertical && b.o == horizontal {
				if b.p0.x <= a.p0.x && b.p1.x >= a.p0.x && a.p0.y <= b.p0.y && a.p1.y >= b.p0.y {

					intersections = append(intersections, point{a.p0.x, b.p0.y})
					// calc dA
					var da, db int
					if a.flipped {
						da = a.p1.y - b.p0.y
					} else {
						da = b.p0.y - a.p0.y
					}
					if b.flipped {
						db = b.p1.x - a.p0.x
					} else {
						db = a.p0.x - b.p0.x
					}
					dists = append(dists, dist1+dist2+da+db)

				}
			} else if a.o == horizontal && b.o == vertical {
				if a.p0.x <= b.p0.x && a.p1.x >= b.p0.x && b.p0.y <= a.p0.y && b.p1.y >= a.p0.y {

					intersections = append(intersections, point{b.p0.x, a.p0.y})
					var da, db int
					if a.flipped {
						da = a.p1.x - b.p0.x
					} else {
						da = b.p0.x - a.p0.x
					}
					if b.flipped {
						db = b.p1.y - a.p0.y
					} else {
						db = a.p0.y - b.p0.y
					}
					dists = append(dists, dist1+dist2+da+db)
				}
			}
			dist2 += b.length
		}
		dist1 += a.length
	}
	return intersections, dists
}
func findCommenInterval(a0 int, a1 int, b0 int, b1 int) ([]int, bool) {
	intersect := false
	r := []int{0, 0}

	if a1 < a0 {
		a0, a1 = mathUtils.Swap(a0, a1)
	}
	if b1 < b0 {
		b0, b1 = mathUtils.Swap(b0, b1)
	}
	s1 := mathUtils.Min(int64(a1), int64(b1))
	if a0 <= b0 && a1 >= b1 {
		intersect = true
		r[0] = b0
		r[1] = b1
	} else if b0 <= a0 && b1 >= a1 {
		intersect = true
		r[0] = a0
		r[1] = a1
	} else if a0 <= b0 && a1 >= b0 {
		intersect = true
		r[0] = b0
		r[1] = int(s1)
	} else if b0 <= a0 && b1 >= a0 {
		intersect = true
		r[0] = a0
		r[1] = int(s1)
	}
	return r, intersect

}
func manhattanDistance(p point) int {
	return mathUtils.Abs32(p.x) + mathUtils.Abs32(p.y)
}

func parseWire(s string) []line {
	ss := strings.Split(s, ",")
	currPoint := point{0, 0} // start at (0,0)
	wire := make([]line, 0)
	for _, dir := range ss {
		var nextPoint point
		var lineSegment line
		var orientation lineOrientation
		steps, _ := strconv.Atoi(dir[1:])
		switch dir[0] {
		case 68: // 68: DOWN
			nextPoint.x = currPoint.x
			nextPoint.y = currPoint.y - steps
			orientation = vertical
			lineSegment = line{nextPoint, currPoint, orientation, steps, true}
		case 85: // 85: UP
			nextPoint.x = currPoint.x
			nextPoint.y = currPoint.y + steps
			orientation = vertical
			lineSegment = line{currPoint, nextPoint, orientation, steps, false}
		case 82: // 82: RIGHT
			nextPoint.x = currPoint.x + steps
			nextPoint.y = currPoint.y
			orientation = horizontal
			lineSegment = line{currPoint, nextPoint, orientation, steps, false}
		case 76: // 76: LEFT
			nextPoint.x = currPoint.x - steps
			nextPoint.y = currPoint.y
			orientation = horizontal
			lineSegment = line{nextPoint, currPoint, orientation, steps, true}
		}
		wire = append(wire, lineSegment)
		currPoint = nextPoint
	}
	return wire
}
func main() {
	// Debug path
	// fmt.Println(os.Getwd())
	// filePath = fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)

	const (
		year = 2019
		day  = 3
	)

	var (
		solution1, solution2 int64
	)

	// IO
	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)

	wire1 := parseWire(lines[0])
	wire2 := parseWire(lines[1])
	intersectionPoints, intersectionDists := intersectWires(wire1, wire2)
	minDist := 2147483647
	minSteps := 2147483647
	for _, ip := range intersectionPoints {
		if ip.x == 0 && ip.y == 0 {
			continue
		} else {
			minDist = mathUtils.Min32(minDist, manhattanDistance(ip))
			fmt.Println(ip)
		}

	}
	for _, d := range intersectionDists {
		if d == 0 {
			continue
		} else {
			minSteps = mathUtils.Min32(minSteps, d)
		}
	}
	solution1 = int64(minDist)
	solution2 = int64(minSteps)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v",
		header, len(lines), solution1, solution2)
}
