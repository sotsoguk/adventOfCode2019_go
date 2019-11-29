package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	// "strconv"
	// "strings"

	// "github.com/adventOfCode2019_go/utils/mathUtils"

	readAOC "github.com/adventOfCode2019_go/utils"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseInput(s string) ingredient {
	var i ingredient
	ss := strings.Split(s, " ")
	i.name = strings.Trim(ss[0], ":")
	regexpNumbers, _ := regexp.Compile(`-?[0-9]+`) // dont forget the minus sign
	numbers := regexpNumbers.FindAllString(s, -1)
	i.capacity, _ = strconv.Atoi(numbers[0])
	i.durability, _ = strconv.Atoi(numbers[1])
	i.flavor, _ = strconv.Atoi(numbers[2])
	i.texture, _ = strconv.Atoi(numbers[3])
	i.calories, _ = strconv.Atoi(numbers[4])
	return i
}
func main() {
	lines := readAOC.ReadInput("2015/inputs/input15_2015.txt")
	var (
		solution1, solution2 int64
		maxScore1, maxScore2 int64
		is                   []ingredient
	)
	fmt.Println(len(lines))

	// test2 := []string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"}

	// for _, l := range test2 {
	// 	fmt.Println(l, nicePart2(l))
	// }
	// parse inputs
	for _, l := range lines {
		ing := parseInput(l)
		is = append(is, ing)
	}
	for _, i := range is {
		fmt.Println(i)
	}

	c500 := make([]int64, 0)
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100-a; b++ {
			for c := 0; c <= 100-(a+b); c++ {
				d := 100 - (a + b + c)
				var currScore int64
				cap := a*is[0].capacity + b*is[1].capacity + c*is[2].capacity + d*is[3].capacity
				if cap < 0 {
					cap = 0
				}
				dur := a*is[0].durability + b*is[1].durability + c*is[2].durability + d*is[3].durability
				if dur < 0 {
					dur = 0
				}
				fla := a*is[0].flavor + b*is[1].flavor + c*is[2].flavor + d*is[3].flavor
				if fla < 0 {
					fla = 0
				}
				tex := a*is[0].texture + b*is[1].texture + c*is[2].texture + d*is[3].texture
				if tex < 0 {
					tex = 0
				}
				currScore = int64(cap) * int64(dur) * int64(fla) * int64(tex)
				cal := a*is[0].calories + b*is[1].calories + c*is[2].calories + d*is[3].calories
				if currScore > maxScore1 {
					maxScore1 = currScore
				}
				if cal == 500 {
					c500 = append(c500, currScore)
				}
			}
		}
	}
	solution1 = maxScore1
	// solve part2
	for _, p2 := range c500 {
		if p2 > maxScore2 {
			maxScore2 = p2
		}
	}
	solution2 = maxScore2
	fmt.Printf("AoC 2015 - Day 05\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
