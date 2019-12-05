package main

import (
	"fmt"
	"time"
	//readAOC "github.com/adventOfCode2019_go/utils"
	//"github.com/adventOfCode2019_go/utils/mathUtils"
)

type Result struct {
	part1 int
	part2 int
}

func makeDigits(n int64) []int {
	digits := make([]int, 6)
	for i := 0; i < 6; i++ {
		digits[5-i] = int(n % 10)
		n /= 10
	}
	return digits
}

func validPassword(n int64) (valid1, valid2 bool) {

	digits := makeDigits(n)
	countDigits := map[int]int{}
	for _, d := range digits {
		countDigits[d]++
	}
	for _, v := range countDigits {
		if v == 2 {
			valid2 = true
		}
		if v >= 2 {
			valid1 = true
		}
	}
	return valid1, valid2
}
func isIncreasing(n int64) (bool, int) {
	var (
		pos        int
		increasing = true
	)

	digits := makeDigits(n)
	for i := 0; i < 5; i++ {
		if digits[i] > digits[i+1] {
			increasing = false
			pos = i + 1
			break
		}
	}
	//create next valid number
	newDigits := make([]int, 6)
	if !increasing {

		for i := 0; i <= pos; i++ {
			newDigits[i] = digits[i]
		}
		newDigits[pos] = newDigits[pos-1]
	}
	newNum := newDigits[0]
	for i := 1; i < 6; i++ {
		newNum = newNum*10 + newDigits[i]
	}
	return increasing, newNum
}
func findConcPasswords(lower int, upper int, c chan Result) {
	var solution1, solution2 int
	for i := lower; i <= upper; {
		if inc, newNum := isIncreasing(int64(i)); inc {
			v1, v2 := validPassword(int64(i))
			if v1 {
				solution1++
			}
			if v2 {
				solution2++
			}
			i++
		} else {
			i = newNum

		}
	}
	// c <- solution1
	res := new(Result)
	res.part1 = solution1
	res.part2 = solution2
	c <- *res

}

func main() {
	// Debug path
	// fmt.Println(os.Getwd())
	// filePath = fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)

	start := time.Now()
	const (
		year = 2019
		day  = 4
	)

	var (
		solution1, solution2 int64
	)

	// IO
	//filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	//lines := readAOC.ReadInput(filePath)
	input := []int64{234208, 765869}

	//create equal parts

	t0 := int(input[0])
	t1 := int(input[1])
	parts := 4
	partlen := int(int(t1-t0) / parts)
	rs := make([]int, 0)
	for i := 0; i < parts; i++ {
		if i < parts-1 {
			rs = append(rs, t0+i*partlen, t0+(i+1)*partlen-1)
		} else {
			rs = append(rs, t0+(parts-1)*partlen, t1)
		}
	}
	fmt.Println(rs)
	// for i := input[0]; i <= input[1]; {
	// 	if inc, newNum := isIncreasing(i); inc {
	// 		v1, v2 := validPassword(i)
	// 		if v1 {
	// 			solution1++
	// 		}
	// 		if v2 {
	// 			solution2++
	// 		}
	// 		i++
	// 	} else {
	// 		i = int64(newNum)

	// 	}
	// }
	c := make(chan Result)
	go findConcPasswords(rs[0], rs[1], c)
	go findConcPasswords(rs[2], rs[3], c)
	go findConcPasswords(rs[4], rs[5], c)
	go findConcPasswords(rs[6], rs[7], c)
	s1, s2, s3, s4 := <-c, <-c, <-c, <-c
	solution1 = int64(s1.part1 + s2.part1 + s3.part1 + s4.part1)
	solution2 = int64(s1.part2 + s2.part2 + s3.part2 + s4.part2)

	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nRuntime:\t%v",
		header, 0, solution1, solution2, elapsed)
}
