package main

import (
	"fmt"
	// "strconv"
	// "strings"

	// "github.com/adventOfCode2019_go/utils/mathUtils"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func vowels(s string) bool {
	var cnt int
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			cnt++
		}
	}
	if cnt >= 3 {
		return true
	} else {
		return false
	}
}

func doubleChar(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		if r[i] == r[i+1] {
			return true
		}
	}
	return false

}

func badStrings(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		switch string(r[i : i+2]) {
		case "ab", "cd", "pq", "xy":
			return true
		}
	}
	return false
}

func nicePart1(s string) bool {
	return vowels(s) && doubleChar(s) && !badStrings(s)
}
func doublePair(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-3; i++ {

		for j := i + 2; j < len(r)-1; j++ {
			if r[i] == r[j] && r[i+1] == r[j+1] {
				return true
			}
		}
	}
	return false
}
func pairWithMiddle(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-2; i++ {
		if r[i] == r[i+2] {
			return true
		}
	}
	return false
}
func nicePart2(s string) bool {
	return doublePair(s) && pairWithMiddle(s)
}
func main() {
	lines := readAOC.ReadInput("2015/inputs/input05_2015.txt")
	var (
		solution1, solution2 int64
	)
	fmt.Println(len(lines))

	for _, line := range lines {
		if nicePart1(line) {
			solution1++
		}
		if nicePart2(line) {
			solution2++
		}
	}
	// test2 := []string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"}

	// for _, l := range test2 {
	// 	fmt.Println(l, nicePart2(l))
	// }

	fmt.Printf("AoC 2015 - Day 05\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
