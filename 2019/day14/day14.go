package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
)

type Recipe struct {
	Ingredient
	inputs []Ingredient
}

type Ingredient struct {
	result string
	cnt    int
}

func calcOre(reactions []string, numFuelDesired int) int {
	recipes := make(map[string]Recipe, len(reactions))
	for _, r := range reactions {
		recipe := parseReaction(r)
		recipes[recipe.result] = recipe
	}
	return produce("FUEL", numFuelDesired, recipes, make(map[string]int))
}

func parseReaction(r string) Recipe {
	parts := strings.Split(r, "=>")
	rResult, rCnt := parseElement(parts[1])

	recipe := Recipe{
		Ingredient: Ingredient{rResult, rCnt},
		inputs:     make([]Ingredient, 0),
	}
	for _, inputStr := range strings.Split(parts[0], ",") {
		input, inputCnt := parseElement(inputStr)
		recipe.inputs = append(recipe.inputs, Ingredient{
			result: input,
			cnt:    inputCnt,
		})
	}
	return recipe
}
func parseElement(element string) (elementName string, elementCount int) {
	element = strings.TrimSpace(element)
	parts := strings.Split(element, " ")
	elementName = parts[1]
	elementCount, _ = strconv.Atoi(parts[0])
	return
}

func part2(reactions []string, ore int) int {
	start := 0
	end := ore
	guesses := 0
	lastGuess := 0
	fuelGuess := 0
	for {
		guesses++
		fuelGuess = (end-start)/2 + start
		reqOre := calcOre(reactions, fuelGuess)
		if reqOre == ore {
			break
		}
		if reqOre > ore {
			end = fuelGuess
		} else {
			start = fuelGuess
		}
		if guesses > 1000 || fuelGuess == lastGuess {
			break
		}
		lastGuess = fuelGuess

	}
	return fuelGuess
}
func produce(desiredElement string, numDesired int, recipes map[string]Recipe, excess map[string]int) int {
	if desiredElement == "ORE" {
		return numDesired
	}
	if excess[desiredElement] >= numDesired {
		excess[desiredElement] -= numDesired
		return 0
	}
	if excess[desiredElement] > 0 {
		numDesired -= excess[desiredElement]
		excess[desiredElement] = 0
	}

	recipe := recipes[desiredElement]
	batches := int(math.Ceil(float64(numDesired) / float64(recipe.cnt)))

	ore := 0
	for _, input := range recipe.inputs {
		ore += produce(input.result, input.cnt*batches, recipes, excess)
	}
	numProduced := batches * recipe.cnt
	excess[desiredElement] += numProduced - numDesired
	return ore

}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 14
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)

	// ll := strings.Split(lines[0], ",")
	// fmt.Println(calcOre(lines, 1))
	fmt.Println(part2(lines, 1000000000000))
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)

}
