package main

import (
	"fmt"
	"os"
	"strconv"
	// "strconv"
	// "strings"
	// "github.com/adventOfCode2019_go/utils/mathUtils"
)

func looknsay(s string) (result string) {
	inRune := []rune(s)
	outRune := make([]rune, 0, 2*len(inRune))
	for i, j := 0, 0; i < len(inRune); i = j {

		j = i + 1
		for j < len(inRune) && inRune[j] == inRune[i] {
			j++
		}
		outRune = append(outRune, []rune(strconv.Itoa(j-i))...)
		outRune = append(outRune, inRune[i])

	}
	return string(outRune)
}
func main() {
	fmt.Println(os.Getwd())
	//lines := readAOC.ReadInput("2015/inputs/input07_2015b.txt")
	// lines := readAOC.ReadInput("../inputs/input07_2015.txt")
	// lines := readAOC.ReadInput("2015/day07/test2.txt")
	var (
		//solution1, solution2 int64
		solution1 int64
		solution2 int64
	)
	//fmt.Printf("Length of input (in lines): %v\n", len(lines))
	input := "1113122113"
	// input := "111221"
	for i := 0; i < 50; i++ {
		if i == 40 {
			solution1 = int64(len(input))
		}
		input = looknsay(input)
		//fmt.Println(i, input)

	}

	// solution1 = looknsay(input)
	solution2 = int64(len(input))
	fmt.Printf("AoC 2015 - Day 07\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
