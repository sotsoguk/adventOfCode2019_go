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

type instr int

const (
	comAND instr = iota
	comOR
	comNOT
	comLSHIFT
	comRSHIFT
	comSET
)

type instruction struct {
	evaluated bool
	command   instr
	op1       string
	op2       string
	value     uint16
}

var circuit = make(map[string]instruction)

func parseLine(l string) {
	var currInstr instruction
	s := strings.Split(l, " ")
	//AND
	if strings.Contains(l, "AND") {
		currInstr.command = comAND
		currInstr.op1 = s[0]
		currInstr.op2 = s[2]
		circuit[s[4]] = currInstr
	} else if strings.Contains(l, "OR") {
		currInstr.command = comOR
		currInstr.op1 = s[0]
		currInstr.op2 = s[2]
		circuit[s[4]] = currInstr
	} else if strings.Contains(l, "NOT") {
		currInstr.command = comNOT
		currInstr.op1 = s[1]
		circuit[s[3]] = currInstr
	} else if strings.Contains(l, "RSHIFT") {
		currInstr.command = comRSHIFT
		currInstr.op1 = s[0]
		currInstr.op2 = s[2]
		circuit[s[4]] = currInstr

	} else if strings.Contains(l, "LSHIFT") {
		currInstr.command = comLSHIFT
		currInstr.op1 = s[0]
		currInstr.op2 = s[2]
		circuit[s[4]] = currInstr
	} else {
		// check if number is set or not
		if _, err := strconv.ParseInt(s[0], 10, 16); err == nil {
			currInstr.evaluated = true
			tmp, _ := strconv.ParseInt(s[0], 10, 16)
			currInstr.value = uint16(tmp)
			circuit[s[2]] = currInstr
		} else {
			currInstr.command = comSET
			currInstr.op1 = s[0]
			circuit[s[2]] = currInstr
		}
	}

}
func evalWire(wire string) uint16 {
	// TODO
	// CHECK FOR NUMERIC ARGUMENTS IN ALLLL commands
	// 1 AND r
	if v, err := strconv.ParseInt(wire, 10, 16); err == nil {
		return uint16(v)
	}
	currInstr, ok := circuit[wire]
	// fmt.Print(wire, "\n")
	if !ok {
		panic(wire)

	}
	// defer fmt.Println("defer at ", currInstr)
	var v1, v2 uint16
	var newValue uint16
	if currInstr.evaluated {
		// fmt.Println(wire, "ALREDAY IN MEMORY")
		return currInstr.value
	} else if currInstr.command == comAND {
		v1 = evalWire(currInstr.op1)
		v2 = evalWire(currInstr.op2)
		newValue = v1 & v2
	} else if currInstr.command == comOR {
		v1 = evalWire(currInstr.op1)
		v2 = evalWire(currInstr.op2)
		newValue = v1 | v2
	} else if currInstr.command == comNOT {
		v1 = evalWire(currInstr.op1)
		newValue = ^v1
	} else if currInstr.command == comLSHIFT {
		v1 = evalWire(currInstr.op1)
		tmp, _ := strconv.ParseInt(currInstr.op2, 10, 16)
		v2 = uint16(tmp)
		newValue = v1 << v2
	} else if currInstr.command == comRSHIFT {
		v1 = evalWire(currInstr.op1)
		tmp, _ := strconv.ParseInt(currInstr.op2, 10, 16)
		v2 = uint16(tmp)
		newValue = v1 >> v2
	} else if currInstr.command == comSET {
		newValue = evalWire(currInstr.op1)
	}
	currInstr.evaluated = true
	currInstr.value = newValue
	circuit[wire] = currInstr
	return newValue
}
func main() {
	fmt.Println(os.Getwd())
	lines := readAOC.ReadInput("2015/inputs/input07_2015b.txt")
	// lines := readAOC.ReadInput("../inputs/input07_2015.txt")
	// lines := readAOC.ReadInput("2015/day07/test2.txt")
	var (
		solution1, solution2 int64
	)
	fmt.Printf("Length of input (in lines): %v\n", len(lines))
	for _, l := range lines {
		parseLine(l)
	}

	// for k, _ := range circuit {
	// 	fmt.Println(evalWire(k))
	// }
	// for k, v := range circuit {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(circuit["b"])
	// fmt.Println(evalWire("lx"))
	fmt.Println(evalWire("a"))
	fmt.Printf("AoC 2015 - Day 07\n-----------------\nPart1:\t%v\nPart2:\t%v", solution1, solution2)
}
