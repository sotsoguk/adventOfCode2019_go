package main

import (
	"fmt"
	"strconv"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func getFuel(mass int64) int64 {
	if mass < 8 {
		return 0
	}
	return (mass / 3) - 2

}
func main() {
	lines := readAOC.ReadInput("2019/inputs/input01_2019.txt")
	var (
		solution1, solution2 int64
	)

	for _, l := range lines {
		var tmpFuel int64
		lint, _ := strconv.ParseInt(l, 10, 64)
		solution1 += getFuel(lint)
		for lint > 0 {
			fuel := getFuel(lint)
			tmpFuel += fuel
			lint = fuel

		}
		solution2 += tmpFuel

	}
	fmt.Printf("AoC 2019 - Day 01\n-----------------\nLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v",
		len(lines), solution1, solution2)
}
