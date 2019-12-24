package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/intcode"
)

type nwPackage struct {
	from int64
	to   int64
	x    int64
	y    int64
}
type Queue struct {
	data []nwPackage
}

func newQueue() *Queue {
	return &Queue{}
}
func (q *Queue) isEmpty() bool {
	if len((*q).data) == 0 {
		return true
	}
	return false
}
func (q *Queue) Pop() nwPackage {

	defer func() {
		q.data = q.data[1:]
	}()

	return q.data[0]
}
func (q *Queue) Append(pck nwPackage) {
	q.data = append(q.data, pck)
}

func runNetwork(numClients int, code []int64, part2 bool) int64 {
	//Setup queue and clients
	q := make([]*Queue, numClients)
	for i := range q {
		q[i] = newQueue()
	}
	network := make([]intcode.VM, numClients)
	for i := 0; i < numClients; i++ {
		network[i].LoadCode(code)
		network[i].Reset()
		network[i].OutputWait = 3
		network[i].LoadInput(int64(i))
		network[i].RunCode()
	}

	nat := nwPackage{}
	natPrevious := int64(-10000)

	for {
		// check if network is idle
		isIdle := func() bool {
			for i := range network {
				if !(q[i].isEmpty()) || network[i].Mode != 1 {
					return false
				}
			}
			return true
		}

		for i := 0; i < numClients; i++ {
			natSent := false
			if part2 {
				if isIdle() {
					network[0].LoadInputs([]int64{nat.x, nat.y})
					if nat.y == natPrevious {
						return nat.y
					}
					natPrevious = nat.y
					natSent = true
				}
				if natSent {
					continue
				}
			}
			// check if client has 3 output values waiting
			if network[i].Mode == 3 {

				to, x, y := network[i].Output[0], network[i].Output[1], network[i].Output[2]
				if to == 255 {
					if part2 {
						nat = nwPackage{int64(i), to, x, y}
					} else {
						return y
					}
				} else {
					pck := nwPackage{int64(i), to, x, y}
					q[to].Append(pck)
				}
				network[i].ClearOuput()
			}
			// check if client is waiting for input
			if network[i].Mode == 1 {
				if q[i].isEmpty() {
					network[i].LoadInput(-1)
				} else {
					pck := q[i].Pop()
					network[i].LoadInputs([]int64{pck.x, pck.y})
				}
			}
			network[i].RunCode()
		}
	}

}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year       = 2019
		day        = 23
		output     = false
		numClients = 50
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	// filePath := fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	ll := strings.Split(lines[0], ",")
	code := make([]int64, len(ll))
	for i := range ll {
		code[i], _ = strconv.ParseInt(ll[i], 10, 64)
	}

	solution1 = runNetwork(numClients, code, false)
	solution2 = runNetwork(numClients, code, true)

	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)

}
