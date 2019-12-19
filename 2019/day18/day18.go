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
			fmt.Print((*g)[y][x])
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
func shortestPath(grid SGrid, start Vec2i, goal Vec2i) Path {
	isUpper := func(s string) bool {
		return "A" <= s && s <= "Z"
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
			break
		}
		//Walk up
		// UP
		v := u
		steps := 0
		for {
			steps++
			v = v.add(Up)
			if grid[v.y][v.x] == "#" || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}
		//Right
		v = u
		steps = 0
		for {
			steps++
			v = v.add(Right)
			if grid[v.y][v.x] == "#" || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}
		//Down
		v = u
		steps = 0
		for {
			steps++
			v = v.add(Down)
			if grid[v.y][v.x] == "#" || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}
		//Left
		v = u
		steps = 0
		for {
			steps++
			v = v.add(Left)
			if grid[v.y][v.x] == "#" || distances[v.y][v.x] <= distances[u.y][u.x] {
				break
			}
			if distances[v.y][v.x] == 10000 {
				distances[v.y][v.x] = distances[u.y][u.x] + steps
				deq.Inject(v)
			}
		}

	}
	// fmt.Println(answer)
	//construct path and look for needed keys
	backPos := goal
	doorsPassed := make([]string, 0)
	for !(backPos.x == start.x && backPos.y == start.y) {
		if lastPos := backPos.add(Up); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
			backPos = lastPos
			// fmt.Println(backPos)
			if isUpper(grid[backPos.y][backPos.x]) {
				doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
			}
		} else if lastPos := backPos.add(Right); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
			backPos = lastPos
			// fmt.Println(backPos)
			if isUpper(grid[backPos.y][backPos.x]) {
				doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
			}
		} else if lastPos := backPos.add(Down); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
			backPos = lastPos
			// fmt.Println(backPos)
			if isUpper(grid[backPos.y][backPos.x]) {
				doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
			}
		} else if lastPos := backPos.add(Left); distances[backPos.y][backPos.x] == distances[lastPos.y][lastPos.x]+1 {
			backPos = lastPos
			// fmt.Println(backPos)
			if isUpper(grid[backPos.y][backPos.x]) {
				doorsPassed = append(doorsPassed, grid[backPos.y][backPos.x])
			}
		}
	}
	// return answer, doorsPassed
	for i := range doorsPassed {
		doorsPassed[i] = strings.ToLower(doorsPassed[i])
	}
	return Path{answer, doorsPassed}
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
	return newKeys
}
func collectKeys(from string, haveKeys []string, remainingKeys []string, lut *map[string]map[string]Path, steps int, mem *map[Mem]int) int {
	if len(remainingKeys) == 0 {
		return steps
	}
	// create bitmap from havekeys
	fmt.Println(haveKeys)
	bm := setBitKeys(haveKeys)
	if val, ok := (*mem)[Mem{bm, from}]; ok {
		fmt.Println("LookUp:", haveKeys, steps)
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
	fmt.Println("Possible:", steps, possibleKeys)
	answer := 1000000
	for _, possKey := range possibleKeys {
		remKeys := removeKey(possKey, remainingKeys)
		havKeys := append(haveKeys, possKey)
		newSteps := (*lut)[from][possKey].length
		tmpSteps := collectKeys(possKey, havKeys, remKeys, lut, newSteps, mem)
		if tmpSteps < answer {
			answer = tmpSteps

		}
	}
	(*mem)[Mem{bm, from}] = answer
	// fmt.Println(possibleKeys)
	return steps + answer
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
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 18
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d_04.txt", year, day, year)
	// filePath := fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	// create array
	grid := input2grid(lines)

	(&grid).print()

	keyMap, _ := findKeysAndDoors(grid)
	startPos := findStartPos(grid)
	lut := make(map[string]map[string]Path, 0)
	// create from startpos to all keys (lowercase,a,bc)
	lut["@"] = map[string]Path{}
	for key := range keyMap {
		lut[key] = map[string]Path{}
	}
	for key, pos := range keyMap {

		lut["@"][key] = shortestPath(grid, startPos, pos)
	}
	for skey, spos := range keyMap {
		for gkey, gpos := range keyMap {
			if gkey == skey {
				continue
			} else {
				tmp := shortestPath(grid, spos, gpos)
				lut[skey][gkey] = tmp
				lut[gkey][skey] = tmp
			}
		}
	}
	allkeys := make([]string, 0)
	for i := range keyMap {
		allkeys = append(allkeys, i)
	}
	fmt.Println(lut)
	for k, v := range lut {
		fmt.Println(k, v)
	}

	memory := make(map[Mem]int, 0)
	fmt.Println(collectKeys("@", []string{}, allkeys, &lut, 0, &memory))
	elapsed := time.Since(start)

	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)
	// fmt.Println(removeKey("A", []string{"B", "C", "A", "D", "E"}))

	// fmt.Println(setBitKeys([]string{"j"}))
}
