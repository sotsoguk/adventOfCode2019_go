package main

import (
	"fmt"
	"sort"
	"strconv"

	// "strconv"
	// "strings"

	// "github.com/adventOfCode2019_go/utils/mathUtils"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func numP(c []int, cap int) int {

	if cap < 0 {
		return 0
	}
	if cap == 0 {
		return 1
	}
	if len(c) < 1 {
		return 0
	}
	return numP(c[1:], cap) + numP(c[1:], cap-c[0])
}
func main() {
	lines := readAOC.ReadInput("2015/inputs/input17_2015.txt")
	var (
		solution1, solution2 int64
	)
	const cap = 150
	fmt.Println(len(lines))
	containers := make([]int, len(lines))
	for i, l := range lines {
		containers[i], _ = strconv.Atoi(l)
	}
	fmt.Println(containers)
	sort.Sort(sort.Reverse(sort.IntSlice(containers)))
	fmt.Println(containers)
	p1 := numP(containers, cap)
	solution1 = int64(p1)
	fmt.Printf("AoC 2015 - Day 17\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
