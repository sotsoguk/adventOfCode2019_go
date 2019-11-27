package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	// "strconv"
	// "strings"
	// "github.com/adventOfCode2019_go/utils/mathUtils"
)

func main() {
	//lines := readAOC.ReadInput("2015/inputs/input03_2015.txt")
	input := "iwrupvqb"
	// input := "abcdef"
	prefix1 := "00000"
	prefix2 := "000000"
	var (
		solution1, solution2 int64
		number               int64
		found1, found2       bool
		toCheck, checkSum    string
		data                 []byte
		checkSumBytes        [16]byte
	)

	for !(found1 && found2) {
		number++
		toCheck = input + strconv.FormatInt(number, 10)
		data = []byte(toCheck)
		checkSumBytes = md5.Sum(data)
		//checkSum = string(checkSumBytes[:16])
		checkSum = hex.EncodeToString(checkSumBytes[:16])

		// fmt.Println(toCheck, checkSum)
		if !found1 && strings.HasPrefix(checkSum, prefix1) {
			found1 = true
			solution1 = number
		}
		if !found2 && strings.HasPrefix(checkSum, prefix2) {
			found2 = true
			solution2 = number
		}
	}
	fmt.Printf("\nAoC 2015 - Day 04\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
