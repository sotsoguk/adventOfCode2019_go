package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	// "strconv"
	// "strings"
	// "github.com/adventOfCode2019_go/utils/mathUtils"
	readAOC "github.com/adventOfCode2019_go/utils"
)

func parseFilter(d interface{}, filter string) []int {
	r := make([]int, 0)

	switch dd := d.(type) {
	case float64:
		r = append(r, int(dd))
	case []interface{}:
		for _, v := range dd {
			si := parseFilter(v, filter)
			r = append(r, si...)
		}
	case map[string]interface{}:
		var cr []int
		var skipCurr bool
		for _, v := range dd {
			if v == filter {
				skipCurr = true
				break
			}
			si := parseFilter(v, filter)
			cr = append(cr, si...)
		}
		if !skipCurr {
			r = append(r, cr...)
		}
	case string:
		//ignore
	default:
		fmt.Println("OOPS")
	}
	return r
}
func main() {
	fmt.Println(os.Getwd())
	lines := readAOC.ReadInput("2015/inputs/input12_2015.txt")
	// lines := readAOC.ReadInput("../inputs/input07_2015.txt")
	// lines := readAOC.ReadInput("2015/day07/test2.txt")
	var (
		//solution1, solution2 int64
		solution1 int64
		solution2 int64
	)
	fmt.Printf("Length of input (in lines): %v\n", len(lines))
	fmt.Println(len(lines[0]))

	numbers, _ := regexp.Compile(`-?[0-9]+`) // dont forget the minus sign
	// input := "111221"

	// solution1 = looknsay(input)
	// test := "adhjh 123jfj 2hj33kjjk"
	// var result []string
	result := numbers.FindAllString(lines[0], -1)
	//fmt.Println("TESTING")
	// for _, v := range result {
	// 	fmt.Println(v)
	// }
	for _, num := range result {
		n, _ := strconv.Atoi(num)
		solution1 += int64(n)
	}
	//parse json and unmarshal
	var d interface{}
	err := json.Unmarshal([]byte(lines[0]), &d)
	if err != nil {
		fmt.Println("ERROR JSON PARSING")
	}
	r2 := parseFilter(d, "red")
	for _, v := range r2 {
		solution2 += int64(v)
	}

	fmt.Printf("AoC 2015 - Day 12\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
