package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	// "strconv"
	// "strings"
	// "github.com/adventOfCode2019_go/utils/mathUtils"
	readAOC "github.com/adventOfCode2019_go/utils"
)

func parseLine(s string) (n1 string, n2 string, h int) {
	lose := 1
	if strings.Contains(s, "lose") {
		lose = -1
	}
	ss := strings.Split(s, " ")
	n1 = ss[0]
	n2 = strings.Trim(ss[10], ".")
	htmp, _ := strconv.Atoi(ss[3])
	h = htmp * lose
	return
}

func GeneratePermutations(data []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}
func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func main() {
	fmt.Println(os.Getwd())
	lines := readAOC.ReadInput("2015/inputs/input13_2015.txt")
	// lines := readAOC.ReadInput("../inputs/input07_2015.txt")
	// lines := readAOC.ReadInput("2015/day07/test2.txt")
	var (
		//solution1, solution2 int64
		solution1 int64
		solution2 int64
	)
	fmt.Printf("Length of input (in lines): %v\n", len(lines))
	nlu := map[string]int{
		"Alice": 0, "Bob": 1, "Carol": 2,
		"David": 3, "Eric": 4, "Frank": 5,
		"George": 6, "Mallory": 7, "Me": 8,
	}

	// create table of happiness ;)
	table := make([][]int, 9)
	for i := 0; i < 9; i++ {
		table[i] = make([]int, 9)
	}
	// parse input
	for _, line := range lines {
		n1, n2, h := parseLine(line)
		table[nlu[n1]][nlu[n2]] = h
	}
	for i := 0; i < 9; i++ {
		fmt.Println(i, table[i])
	}
	var highest int
	for perm := range GeneratePermutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}) {
		cur := 0
		for j := 0; j < 9; j++ {
			k := j + 1
			if k >= 9 {
				k = 0
			}
			cur += table[perm[j]][perm[k]]
			cur += table[perm[k]][perm[j]]
			if cur > highest {
				highest = cur
			}
		}
	}
	solution1 = int64(highest)
	// fmt.Println(len(lines[0]))
	fmt.Println(parseLine(lines[3]))
	fmt.Printf("AoC 2015 - Day 12\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
