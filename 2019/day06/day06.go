package main

import (
	"fmt"
	"strings"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type Node struct {
	name         string
	orbitsAround *Node
	planets      []*Node
}

func processInput(lines []string) map[string]string {
	planetMap := make(map[string]string, 0)
	for _, ll := range lines {
		parts := strings.Split(ll, ")")
		planetMap[parts[1]] = parts[0]
		// if ok {
		// 	fmt.Println(parts[1], " already in Map ?!?")
		// }
	}
	return planetMap
}

func findPlanet(root *Node, planetName string) *Node {
	if root.name == planetName {
		return root
	} else {
		if len(root.planets) > 0 {
			for _, children := range root.planets {
				if pl := findPlanet(children, planetName); pl != nil {
					return pl
				}
			}
		}
	}
	return nil
}
func printSolarSystem(root *Node) {
	fmt.Println(root.name)
	if len(root.planets) > 0 {
		for _, c := range root.planets {
			fmt.Print(root.name, "->")
			printSolarSystem(c)
		}
	}
}
func pathToPlanet(root *Node, name string, path []string) ([]string, bool) {
	if root.name == name {
		return append(path, root.name), true
	}
	if len(root.planets) > 0 {
		for _, c := range root.planets {
			newPath, found := pathToPlanet(c, name, path)
			if found {
				// fmt.Println("F: ", root.name, path, newPath)
				newPath := append(newPath, root.name)
				return append(path, newPath...), true
			}
		}
	}
	return path, false
}
func countOrbits(root *Node, level int) int {
	num := level
	if len(root.planets) > 0 {
		for _, c := range root.planets {
			num += countOrbits(c, level+1)
		}
	}
	return num
}
func reverseSlice(a []string) []string {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
func lca(a []string, b []string) (string, int) {
	level := 0
	minLen := mathUtils.Min32(len(a), len(b))
	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			if i > 0 {
				return a[i-1], level - 1
			} else {
				return "ERROR", 0
			}
		}
		level++
	}
	return "-1", minLen
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input02_2019.txt")
	// fmt.Println(os.Getwd())

	const (
		year = 2019
		day  = 6
	)

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	// filePath := fmt.Sprintf("%d/inputs/input%02d_test.txt", year, day)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)

	var (
		solution1, solution2 int64
	)
	com := &Node{name: "COM"}
	planets := processInput(lines)
	// fmt.Println(planets)

	// insert planets
	for len(planets) > 0 {
		for planet, orbitOf := range planets {
			if o := findPlanet(com, orbitOf); o != nil {
				newPlanet := &Node{name: planet, orbitsAround: o}
				o.planets = append(o.planets, newPlanet)
				delete(planets, planet)
			}
		}
	}
	// fmt.Println(findPlanet(com, "D") != nil)
	// fmt.Println(findPlanet(com, "F") != nil)
	// fmt.Println(findPlanet(com, "PP") != nil)
	// //printSolarSystem(com)
	fmt.Println(countOrbits(com, 0))
	pathYou := make([]string, 0)
	pathSanta := make([]string, 0)
	pathYou, _ = pathToPlanet(com, "YOU", pathYou)
	pathSanta, _ = pathToPlanet(com, "SAN", pathSanta)
	pathSanta = reverseSlice(pathSanta)
	pathYou = reverseSlice(pathYou)
	// fmt.Println(pathSanta)
	// fmt.Println(pathYou)
	// fmt.Println(lca(pathYou, pathSanta))
	_, levelLCA := lca(pathYou, pathSanta)
	levelYou := len(pathYou) - 1
	levelSanta := len(pathSanta) - 1
	diff := levelYou + levelSanta - 2*levelLCA
	fmt.Println("2:", diff-2)
	// fmt.Println(pathToPlanet(com, "YOU", path))
	// prepare input
	// code8 := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v",
		header, len(lines), solution1, solution2)
}
