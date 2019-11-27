package main

import (
	"fmt"
	"strconv"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func main() {
	lines := readAOC.ReadInput("2018/inputs/input01_2018.txt")
	// for _, l := range lines {
	// 	fmt.Println(l)
	// }
	var freq int
	var freqs = make([]int, 0)
	for _, l := range lines {
		v, _ := strconv.Atoi(l)
		freq += v
		freqs = append(freqs, freq)
	}
	fmt.Println("Part 1: ", freq)

	offset := freq
	seenFreqs := make(map[int]int)
	// part 2

	var solutionPart2 int
Loop:
	for {
		for _, v := range freqs {
			// check if value is already in map
			_, ok := seenFreqs[v]
			if ok {
				//duplicateFreq = true
				solutionPart2 = v
				break Loop
			} else {
				seenFreqs[v] = 0
			}
		}
		for i := range freqs {
			freqs[i] += offset
		}

	}
	fmt.Println("Part 2: ", solutionPart2)
}
