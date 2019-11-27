package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adventOfCode2019_go/utils/mathUtils"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func main() {
	lines := readAOC.ReadInput("2015/inputs/input02_2015.txt")

	fmt.Println(len(lines))
	var (
		totalPaper  int64
		totalRibbon int64
	)

	for _, v := range lines {
		//split line into parts
		var (
			currPaper  int64
			currRibbon int64
		)
		dims := strings.Split(v, "x")
		//fmt.Println(dims)
		a, _ := strconv.ParseInt(dims[0], 10, 64)
		b, _ := strconv.ParseInt(dims[1], 10, 64)
		c, _ := strconv.ParseInt(dims[2], 10, 64)

		currPaper = 2*(a*b+a*c+b*c) + mathUtils.Minxyz(a*b, b*c, a*c)
		currRibbon = 2 * mathUtils.Minxyz(a+b, a+c, b+c)
		currRibbon += a * b * c
		totalPaper += currPaper
		totalRibbon += currRibbon

	}
	fmt.Printf("AoC 2015 - Day 02\n-----------------\nPart1:\t%v\nPart2:\t%v", totalPaper, totalRibbon)
}
