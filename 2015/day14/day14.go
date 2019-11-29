package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	// "strconv"
	// "strings"
	// "github.com/adventOfCode2019_go/utils/mathUtils"
	readAOC "github.com/adventOfCode2019_go/utils"
)

type reindeer struct {
	name     string
	speed    int
	flytime  int
	resttime int
	dist     int
	flying   bool
	next     int
	points   int
}

func distance(rd reindeer, time int) int {
	var dist int
	var totalCycles int
	totalCycles = time / int(rd.flytime+rd.resttime)
	dist += totalCycles * (rd.flytime * rd.speed)
	remTime := time - totalCycles*(rd.flytime+rd.resttime)
	if remTime >= rd.flytime {
		dist += rd.speed * rd.flytime
	} else {
		dist += rd.speed * remTime
	}
	return dist
}
func main() {
	fmt.Println(os.Getwd())
	lines := readAOC.ReadInput("2015/inputs/input14_2015.txt")
	// lines := readAOC.ReadInput("../inputs/input07_2015.txt")
	// lines := readAOC.ReadInput("2015/day07/test2.txt")
	var (
		//solution1, solution2 int64
		solution1 int64
		solution2 int64
	)
	fmt.Printf("Length of input (in lines): %v\n", len(lines))
	var rs []reindeer
	// fmt.Println(len(lines[0]))
	for _, l := range lines {
		var tmp int
		ll := strings.Split(l, " ")
		var currDeer reindeer
		currDeer.name = ll[0]
		tmp, _ = strconv.Atoi(ll[3])
		currDeer.speed = tmp
		tmp, _ = strconv.Atoi(ll[6])
		currDeer.flytime = tmp
		tmp, _ = strconv.Atoi(ll[13])
		currDeer.resttime = tmp
		currDeer.flying = true
		currDeer.next = currDeer.flytime
		rs = append(rs, currDeer)
	}
	time := 2503
	var maxDist int
	for _, v := range rs {
		fmt.Println(v)
		currDist := distance(v, time)
		if currDist > maxDist {
			maxDist = currDist
		}
	}
	solution1 = int64(maxDist)

	// Part 2
	for t := 0; t < time; t++ {
		// advance all reindeers
		for k := range rs {
			rd := &rs[k] // pointer to current reindeer
			// check if cycle change
			if rd.next == t {
				if rd.flying {
					rd.flying = false
					rd.next += rd.resttime
				} else {
					rd.flying = true
					rd.next += rd.flytime
				}

			}
			if rd.flying {
				rd.dist += rd.speed
			}
		}
		// look for maxDistance
		var maxDist int
		for i := range rs {
			if rs[i].dist > maxDist {
				maxDist = rs[i].dist
			}
		}
		// apply points
		for k := range rs {
			rd := &rs[k]
			if rd.dist == maxDist {
				rd.points++
			}
		}

	}
	for _, r := range rs {
		fmt.Println(r.points)
	}
	fmt.Printf("AoC 2015 - Day 12\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
