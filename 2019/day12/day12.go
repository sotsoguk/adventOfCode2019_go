package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type Vec3i struct {
	X int64
	Y int64
	Z int64
}

type Planet struct {
	Pos Vec3i
	Vel Vec3i
	// Acc Vec3i
}

func (x Vec3i) Add(y Vec3i) Vec3i {
	var z Vec3i
	z.X = x.X + y.X
	z.Y = x.Y + y.Y
	z.Z = x.Z + y.Z
	return z
}

func (p Planet) Energy() int64 {
	var ePot, eKin int64
	ePot += mathUtils.Abs64(p.Pos.X)
	ePot += mathUtils.Abs64(p.Pos.Y)
	ePot += mathUtils.Abs64(p.Pos.Z)
	eKin += mathUtils.Abs64(p.Vel.X)
	eKin += mathUtils.Abs64(p.Vel.Y)
	eKin += mathUtils.Abs64(p.Vel.Z)
	return ePot * eKin
}
func doStep(planets []Planet) {
	numPlanets := len(planets)
	//compute acceleration
	for i := 0; i < numPlanets; i++ {
		for j := 0; j < numPlanets; j++ {
			if i == j {
				continue
			} else {
				xDiff := planets[i].Pos.X - planets[j].Pos.X
				yDiff := planets[i].Pos.Y - planets[j].Pos.Y
				zDiff := planets[i].Pos.Z - planets[j].Pos.Z
				planets[i].Vel.X -= mathUtils.Sign(xDiff)
				planets[i].Vel.Y -= mathUtils.Sign(yDiff)
				planets[i].Vel.Z -= mathUtils.Sign(zDiff)
			}
		}
	}

	// apply velocity
	for i := range planets {
		planets[i].Pos = planets[i].Pos.Add(planets[i].Vel)
	}
}

func planetsEqual(a Planet, b Planet, coord int) bool {
	if coord == 0 {
		return a.Pos.X == b.Pos.X && a.Vel.X == b.Vel.X
	} else if coord == 1 {
		return a.Pos.Y == b.Pos.Y && a.Vel.Y == b.Vel.Y
	} else {
		return a.Pos.Z == b.Pos.Z && a.Vel.Z == b.Vel.Z
	}
}

func part1(planetsInit []Planet, steps int) (energy int64) {

	planets := make([]Planet, len(planetsInit))
	copy(planets, planetsInit)
	for i := 0; i < steps; i++ {
		doStep(planets)
	}

	for i := range planets {
		energy += planets[i].Energy()
	}
	return
}
func part2(planetsInit []Planet) int64 {
	planets := make([]Planet, len(planetsInit))
	copy(planets, planetsInit)
	counters := make([]int64, 3)
	done := make([]bool, 3)
	allEqual := done[0] && done[1] && done[2]

	for !allEqual {
		doStep(planets)

		for c := 0; c < 3; c++ {
			if !done[c] {
				equal := true
				for i := range planets {
					if !planetsEqual(planets[i], planetsInit[i], c) {
						equal = false
					}
				}
				if equal {
					done[c] = true
				} else {
					counters[c]++
				}
			}

		}
		allEqual = done[0] && done[1] && done[2]
	}
	for i := range counters {
		counters[i]++
	}
	return mathUtils.Lcm(counters[0], counters[1], counters[2])
}

func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()
	const (
		year   = 2019
		day    = 12
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO
	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	re := regexp.MustCompile("-?[0-9]+")
	planetInit := make([]Planet, len(lines))
	for i, p := range lines {
		coords := re.FindAllString(p, -1)
		var currPlanet Planet
		currPlanet.Pos.X, _ = strconv.ParseInt(coords[0], 10, 64)
		currPlanet.Pos.Y, _ = strconv.ParseInt(coords[1], 10, 64)
		currPlanet.Pos.Z, _ = strconv.ParseInt(coords[2], 10, 64)
		planetInit[i] = currPlanet
	}

	solution1 = part1(planetInit, 1000)
	solution2 = part2(planetInit)
	elapsed := time.Since(start)

	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)

}
