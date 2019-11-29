package main

import (
	"fmt"
	"strconv"
	"strings"

	// "strconv"
	// "strings"

	// "github.com/adventOfCode2019_go/utils/mathUtils"

	readAOC "github.com/adventOfCode2019_go/utils"
)

type aunt map[string]int

func main() {
	lines := readAOC.ReadInput("2015/inputs/input16_2015.txt")
	var (
		solution1, solution2 int64
	)
	fmt.Println(len(lines))

	// test2 := []string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"}

	// for _, l := range test2 {
	// 	fmt.Println(l, nicePart2(l))
	// }
	// parse inputs
	aunts := make([]aunt, 500)
	// var aunts []aunt

	//parse input
	// regexpAunts, _ := regexp.Compile(`\D+: \d+`) // dont forget the minus sign
	// ss := regexpAunts.FindAllString(lines[0], -1)
	// ts := strings.ReplaceAll(lines[0], ",", "")
	// ts = strings.ReplaceAll(ts, ":", "")
	// ss2 := strings.Split(ts, " ")
	// for _, v := range ss2 {
	// 	fmt.Println(v)
	// }
	// sue1 := make(aunt)
	// for i := 0; i < len(ss2)-1; i += 2 {
	// 	if i == 0 {
	// 		i += 2
	// 	}
	// 	val, _ := strconv.Atoi(ss2[i+1])
	// 	fmt.Println("!:", ss2[i], val)
	// 	sue1[ss2[i]] = val
	// }
	// fmt.Println(sue1)
	// aunts = append(aunts, sue1)

	for i, l := range lines {
		ts := strings.ReplaceAll(l, ",", "")
		ts = strings.ReplaceAll(ts, ":", "")
		tss := strings.Split(ts, " ")

		currAunt := make(aunt)
		for i := 0; i < len(tss)-1; i += 2 {
			if i == 0 {
				i += 2
			}
			val, _ := strconv.Atoi(tss[i+1])
			//fmt.Println("!:", ss2[i], val)
			currAunt[tss[i]] = val
		}
		// aunts = append(aunts, currAunt)
		aunts[i] = currAunt
		// fmt.Println(currAunt)
	}
	// for _, a := range aunts {
	// 	fmt.Println(a)
	// }
	evidence := map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3,
		"akitas": 0, "viszlas": 0, "goldfish": 5, "trees": 3, "cars": 2,
		"perfumes": 1}

	var maxMatches, maxIndex int
	for i, a := range aunts {
		var currMatches int
		for k, v := range evidence {
			if n, ok := a[k]; ok && n == v {
				currMatches++

			}
		}
		if currMatches > maxMatches {
			maxMatches = currMatches
			maxIndex = i
		}
	}

	var maxMatches2, maxIndex2 int
	for i, a := range aunts {
		var currMatches int
		for k, v := range evidence {
			switch k {
			case "cats", "trees":
				if n, ok := a[k]; ok && n > v {
					currMatches++

				}

			case "pomeranians", "goldfish":
				if n, ok := a[k]; ok && n < v {
					currMatches++

				}
			default:
				if n, ok := a[k]; ok && n == v {
					currMatches++

				}
			}
		}
		if currMatches > maxMatches2 {
			maxMatches2 = currMatches
			maxIndex2 = i
		}
	}
	solution1 = int64(maxIndex + 1)
	solution2 = int64(maxIndex2 + 1)
	fmt.Printf("AoC 2015 - Day 16\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
