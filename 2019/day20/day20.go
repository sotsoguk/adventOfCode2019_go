package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
)

func NewDeque() *Deque {
	return &Deque{}
}

type Deque struct {
	Items []interface{}
}

func (s *Deque) Push(item interface{}) {
	temp := []interface{}{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) Inject(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Deque) Pop() interface{} {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) Eject() interface{} {
	i := len(s.Items) - 1
	defer func() {
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
	}()
	return s.Items[i]
}

func (s *Deque) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

type Vec2i struct {
	x int
	y int
}
type Path struct {
	length     int
	keysNeeded []string
}

type Mem struct {
	bitmap int
	pos    string
}

type Locations map[string][]Vec2i

var (
	Up    = Vec2i{0, -1}
	Down  = Vec2i{0, 1}
	Left  = Vec2i{-1, 0}
	Right = Vec2i{1, 0}
)

func (v Vec2i) add(w Vec2i) Vec2i {
	return Vec2i{v.x + w.x, v.y + w.y}
}

type SGrid [][]string

func (g *SGrid) print() {
	for y := 0; y < len(*g); y++ {
		for x := 0; x < len((*g)[0]); x++ {
			tile := (*g)[y][x]
			if len(tile) > 1 {
				fmt.Print("*")
			} else {
				fmt.Print((*g)[y][x])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func input2grid(lines []string) SGrid {
	grid := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]string, len(lines[0]))
	}
	for y := 0; y < len(lines); y++ {
		ll := strings.Split(lines[y], "")
		for i, c := range ll {
			grid[y][i] = c
		}
	}
	return grid

}
func findKeysAndDoors(grid SGrid) (keyPos map[string]Vec2i, doorPos map[string]Vec2i) {
	keyPos = make(map[string]Vec2i)
	doorPos = make(map[string]Vec2i)
	isUpper := func(s string) bool {
		return "A" <= s && s <= "Z"
	}
	isLower := func(s string) bool {
		return "a" <= s && s <= "z"
	}
	h, w := len(grid), len(grid[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if isLower(grid[y][x]) {
				keyPos[grid[y][x]] = Vec2i{x, y}
			}
			if isUpper(grid[y][x]) {
				doorPos[grid[y][x]] = Vec2i{x, y}
			}
		}
	}
	return keyPos, doorPos
}

func findStartPos(grid SGrid) Vec2i {
	h, w := len(grid), len(grid[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == "@" {
				return Vec2i{x, y}
			}
		}
	}
	return Vec2i{-1, -1}
}
func findMultipleStartPos(grid SGrid) []Vec2i {
	result := make([]Vec2i, 0)
	h, w := len(grid), len(grid[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == "@" {
				result = append(result, Vec2i{x, y})
			}
		}
	}
	return result
}
func shortestPath(grid SGrid, start Vec2i, goal Vec2i, portals map[Vec2i][]Vec2i) int {
	// isPortal := func(s string) bool {
	// 	return len(s) > 1
	// }
	isPortal := func(v Vec2i) bool {
		_, ok := portals[v]
		return ok
	}
	distances := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		distances[i] = make([]int, len(grid[0]))
	}
	for y := 0; y < len(distances); y++ {
		for x := 0; x < len(distances[0]); x++ {
			distances[y][x] = 10000
		}

	}
	dirs := []Vec2i{Up, Right, Down, Left}
	// PrintGridNumbers(distances)
	// find shortes path

	deq := NewDeque()
	//start := Vec2i{21, 21}
	// goal := Vec2i{39, 1}
	cpos := start
	answer := -1
	distances[cpos.y][cpos.x] = 0
	deq.Push(cpos)
	for !deq.IsEmpty() {
		u := deq.Pop().(Vec2i)
		if u.y == goal.y && u.x == goal.x {
			answer = distances[u.y][u.x]
			fmt.Println(answer)
			break
		}
		//Walk up
		// UP
		v := u
		for _, dir := range dirs {
			v = u
			v = v.add(dir)
			// fmt.Println(v)
			if grid[v.y][v.x] == "#" || grid[v.y][v.x] == " " || grid[v.y][v.x] == "AA" || distances[v.y][v.x] <= distances[u.y][u.x] {
				continue
			}
			if distances[v.y][v.x] == 10000 {
				if isPortal(v) {
					distances[v.y][v.x] = distances[u.y][u.x]
					w := portals[v][0]
					distances[w.y][w.x] = distances[u.y][u.x]
					v = portals[v][1]

				}
				distances[v.y][v.x] = distances[u.y][u.x] + 1
				deq.Inject(v)
			}
		}
		//Right
	}

	// fmt.Println(answer)
	//construct path and look for needed keys
	// backPos := goal
	// doorsPassed := make([]string, 0)
	// for !(backPos.x == start.x && backPos.y == start.y) {
	// 	if lastPos := backPos.add(Up); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
	// 		backPos = lastPos
	// 		// fmt.Println(backPos)
	// 		if isUpper(grid[backPos.y][backPos.x]) {
	// 			doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
	// 		}
	// 	} else if lastPos := backPos.add(Right); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
	// 		backPos = lastPos
	// 		// fmt.Println(backPos)
	// 		if isUpper(grid[backPos.y][backPos.x]) {
	// 			doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
	// 		}
	// 	} else if lastPos := backPos.add(Down); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
	// 		backPos = lastPos
	// 		// fmt.Println(backPos)
	// 		if isUpper(grid[backPos.y][backPos.x]) {
	// 			doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
	// 		}
	// 	} else if lastPos := backPos.add(Left); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
	// 		backPos = lastPos
	// 		// fmt.Println(backPos)
	// 		if isUpper(grid[backPos.y][backPos.x]) {
	// 			doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
	// 		}
	// 	}
	// }
	// return answer, doorsPassed
	// for i := range doorsPassed {
	// 	doorsPassed[i] = strings.ToLower(doorsPassed[i])
	// }
	// return Path{answer, doorsPassed}
	return 0
}

func sufficientKeys(keysNeeded []string, keysHave []string) bool {
	if len(keysNeeded) > len(keysHave) {
		return false
	}
	sort.Strings(keysNeeded)
	sort.Strings(keysHave)
	// fmt.Println(keysNeeded, keysHave)
	for i := range keysNeeded {
		keyFound := false
		for j := range keysHave {
			if keysNeeded[i] == keysHave[j] {
				keyFound = true
				continue
			}
		}
		if !keyFound {
			return false
		}
	}
	return true
}
func removeKey(key string, keys []string) []string {
	newKeys := make([]string, 0)
	for _, k := range keys {
		if k != key {
			newKeys = append(newKeys, k)
		}
	}
	sort.Strings(newKeys)
	return newKeys
}
func collectKeys(from string, haveKeys2 []string, remainingKeys2 []string, lut *map[string]map[string]Path, steps int, mem *map[Mem]int) int {
	if len(remainingKeys2) == 0 {
		// fmt.Println("All Done:", steps)
		return steps
	}
	// fmt.Print(from, "<")
	haveKeys := make([]string, len(haveKeys2))
	copy(haveKeys, haveKeys2)
	remainingKeys := make([]string, len(remainingKeys2))
	copy(remainingKeys, remainingKeys2)
	sort.Strings(remainingKeys)
	sort.Strings(haveKeys)
	// create bitmap from havekeys
	// fmt.Println(haveKeys)
	bm := setBitKeys(haveKeys)
	if val, ok := (*mem)[Mem{bm, from}]; ok {
		// fmt.Println("LookUp:", haveKeys, steps)
		return steps + val
	}
	//collect all keys which we can collect
	possibleKeys := make([]string, 0)
	for _, key := range remainingKeys {
		// fmt.Println((*lut)[from][key].keysNeeded)
		if sufficientKeys((*lut)[from][key].keysNeeded, haveKeys) {
			possibleKeys = append(possibleKeys, key)
		}
	}
	sort.Strings(possibleKeys)
	// fmt.Println("Possible:", steps, from, remainingKeys, possibleKeys, haveKeys)
	answer := 1000000
	for _, possKey := range possibleKeys {
		remKeys := removeKey(possKey, remainingKeys)
		havKeys := append(haveKeys, possKey)
		newSteps := (*lut)[from][possKey].length
		tmpSteps := collectKeys(possKey, havKeys, remKeys, lut, newSteps, mem)
		//save state
		//(*mem)[Mem{setBitKeys(havKeys), possKey}] = answer - newSteps
		if tmpSteps < answer {
			answer = tmpSteps

		}
	}
	// fmt.Println(">", from)
	(*mem)[Mem{bm, from}] = answer
	// fmt.Println(possibleKeys)
	return steps + answer
}
func findLocations(grid [][]string) Locations {
	h, w := len(grid), len(grid[0])
	locs := make(map[string][]Vec2i)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			currTile := grid[y][x]
			if len(currTile) < 2 {
				continue
			} else {
				if _, ok := locs[currTile]; !ok {
					locs[currTile] = make([]Vec2i, 0)
				}
				locs[currTile] = append(locs[currTile], Vec2i{x, y})
			}
		}
	}
	return locs
}
func setBitKeys(keys []string) int {
	bitMap := 0
	for i := range keys {
		b := int(keys[i][0]) - 97
		m := 1 << b
		bitMap = bitMap | m
	}
	return bitMap
}
func parseGrid20(grid [][]string) SGrid {
	isUpper := func(s string) bool {
		return "A" <= s && s <= "Z"
	}
	h, w := len(grid), len(grid[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			currTile := grid[y][x]
			if isUpper(currTile) {
				// check letter below
				//upper border
				if y == 0 {
					nextTile := grid[1][x]
					grid[1][x] = strings.Join([]string{currTile, nextTile}, "")
					grid[0][x] = " "
				} else if y == h-2 {
					nextTile := grid[h-1][x]
					grid[h-2][x] = strings.Join([]string{currTile, nextTile}, "")
					grid[h-1][x] = " "
				} else if x == 0 {
					if nextTile := grid[y][1]; isUpper(nextTile) {
						grid[y][1] = strings.Join([]string{currTile, nextTile}, "")
						grid[y][0] = " "
					}
				} else if x == w-2 {
					if nextTile := grid[y][w-1]; isUpper(nextTile) {
						grid[y][w-2] = strings.Join([]string{currTile, nextTile}, "")
						grid[y][w-1] = " "
					}

				} else if nextTile := grid[y+1][x]; isUpper(nextTile) {
					if grid[y+2][x] == " " {
						grid[y][x] = strings.Join([]string{currTile, nextTile}, "")
						grid[y+1][x] = " "
					} else {
						grid[y+1][x] = strings.Join([]string{currTile, nextTile}, "")
						grid[y][x] = " "
					}

				} else if nextTile := grid[y][x+1]; isUpper(nextTile) {
					if grid[y][x+2] == " " {
						grid[y][x] = strings.Join([]string{currTile, nextTile}, "")
						grid[y][x+1] = " "
					} else {
						grid[y][x+1] = strings.Join([]string{currTile, nextTile}, "")
						grid[y][x] = " "
					}
				}

			}
		}
	}
	return grid

}

// func part2(grid SGrid) {
// 	sPos := findMultipleStartPos(grid)
// 	fmt.Println(sPos)
// 	numRobots := len(sPos)
// 	luts := make([]map[string]map[string]Path, numRobots)

// 	keyMap, _ := findKeysAndDoors(grid)
// 	for i := 0; i < numRobots; i++ {
// 		luts[i] = make(map[string]map[string]Path, 0)
// 		luts[i]["@"] = map[string]Path{}
// 		for k := range keyMap {
// 			luts[i][k] = map[string]Path{}
// 		}
// 	}
// 	for key, pos := range keyMap {
// 		for i := 0; i < numRobots; i++ {
// 			p, ok := shortestPath2(grid, sPos[i], pos)
// 			if ok {
// 				luts[i]["@"][key] = p
// 			}
// 		}
// 	}
// 	for i := 0; i < numRobots; i++ {
// 		fmt.Println(luts[i]["@"])
// 	}
// }

// func splitGrid(grid1 SGrid, grid2 SGrid) []SGrid {
// 	start := findStartPos(grid1)
// 	grids := make([]SGrid, 4)
// 	grids[0] = make([][]string, start.y)
// 	for i := 0; i < start.y; i++ {
// 		grids[0][i] = make([]string, start.x)
// 	}
// 	for y := 0; y < start.y; y++ {
// 		for x := 0; x < start.x; x++ {
// 			grids[0][y][x] = grid2[y][x]
// 		}
// 	}
// 	grids[1] = make([][]string, start.y)
// 	for i := 0; i < start.y; i++ {
// 		grids[0][i] = make([]string)
// 	}
// 	for y := 0; y < start.y; y++ {
// 		for x := 0; x < start.x; x++ {
// 			grids[0][y][x] = grid2[y][x]
// 		}
// 	}
// }
func nextFreeTile(grid [][]string, pos Vec2i) Vec2i {
	// var freePos Vec2i
	x, y := pos.x, pos.y
	if grid[y][x+1] == "." {
		return Vec2i{x + 1, y}
	} else if grid[y][x-1] == "." {
		return Vec2i{x - 1, y}

	} else if grid[y+1][x] == "." {
		return Vec2i{x, y + 1}
	} else if grid[y-1][x] == "." {
		return Vec2i{x, y - 1}
	}
	return Vec2i{}
}
func findStartEnd(grid [][]string, locs Locations) (start Vec2i, end Vec2i) {
	start = nextFreeTile(grid, locs["AA"][0])
	end = nextFreeTile(grid, locs["ZZ"][0])
	return
}
func parsePortals(grid [][]string, locs Locations) map[Vec2i][]Vec2i {
	portals := make(map[Vec2i][]Vec2i)
	for k, v := range locs {
		if k == "AA" || k == "ZZ" {
			continue
		}
		p1 := v[0]
		p2 := v[1]
		portals[p1] = []Vec2i{p2, nextFreeTile(grid, p2)}
		portals[p2] = []Vec2i{p1, nextFreeTile(grid, p1)}

	}
	return portals
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 20
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	// filePath2 := fmt.Sprintf("%d/inputs/input%02d_%d_p2.txt", year, day, year)
	// filePath := fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)

	sg := input2grid(lines)
	sg.print()
	sg = parseGrid20(sg)
	sg.print()
	locs := findLocations(sg)

	fmt.Println(locs)
	startPos, endPos := findStartEnd(sg, locs)
	fmt.Println(startPos, endPos)
	portals := parsePortals(sg, locs)
	elapsed := time.Since(start)
	ss1 := shortestPath(sg, startPos, endPos, portals)
	fmt.Println(ss1)
	// a := []int{1, 2, 3, 4, 5}
	// b := append(a, 6)
	// fmt.Println(a, b)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)
	// fmt.Println(removeKey("A", []string{"B", "C", "A", "D", "E"}))

	// fmt.Println(setBitKeys([]string{"j"}))

	// lines2 := readAOC.ReadInput(filePath2)
	// create array
	// grid := input2grid(lines)
	// grid2 := input2grid(lines2)
	// (&grid).print()

	// keyMap, _ := findKeysAndDoors(grid)
	// startPos := findStartPos(grid)
	// fmt.Println(startPos)
	// lut := make(map[string]map[string]Path, 0)
	// // create from startpos to all keys (lowercase,a,bc)
	// lut["@"] = map[string]Path{}
	// for key := range keyMap {
	// 	lut[key] = map[string]Path{}
	// }
	// for key, pos := range keyMap {

	// 	lut["@"][key] = shortestPath(grid, startPos, pos)
	// }
	// for skey, spos := range keyMap {
	// 	for gkey, gpos := range keyMap {
	// 		if gkey == skey {
	// 			continue
	// 		} else {
	// 			tmp := shortestPath(grid, spos, gpos)
	// 			lut[skey][gkey] = tmp
	// 			lut[gkey][skey] = tmp
	// 		}
	// 	}
	// }
	// allkeys := make([]string, 0)
	// for i := range keyMap {
	// 	allkeys = append(allkeys, i)
	// }

	// fmt.Println(lut)
	// for k, v := range lut {
	// 	fmt.Println(k, v)
	// }

	// memory := make(map[Mem]int, 0)
	// fmt.Println(collectKeys("@", []string{}, allkeys, &lut, 0, &memory))
	// a1 := []string{"a", "c", "d"}
	// a2 := []string{"c", "d", "a"}
	// fmt.Println(setBitKeys(a1))
	// fmt.Println(setBitKeys(a2))
	// part2(grid2)
}
