package main

import (
	"fmt"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func main() {
	lines := readAOC.ReadInput("2015/inputs/input01_2015.txt")

	fmt.Println(len(lines))
	var (
		currFloor, solution2 int
		basementFound        bool
	)
	// iterate over all characters and save first poistion to reach -1

	for index, symbol := range lines[0] {
		if symbol == '(' {
			currFloor++
		} else {
			currFloor--
		}
		if !basementFound && currFloor == -1 {
			solution2 = index + 1 // 0 index vs 1 index
			basementFound = true
		}
	}
 
	fmt.Printf("AoC 2015 - Day 01\n-----------------\nPart1:\t%v\nPart2:\t%v", currFloor, solution2)
}
