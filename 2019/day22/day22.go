package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	// . "github.com/lukechampine/uint128"
)

const n2 int64 = 119315717514047
const nn2 int64 = 101741582076661

func mul(a int64, b int64, n int64) int64 {
	if b == 0 {
		return 0

	}
	if b == 1 {
		return a
	}
	var result int64
	result = mul(a, b/2, n)
	result = (result + result) % n
	if (b & 1) == 1 {
		return (result + a) % n
	}
	return result
}

func pow(base int64, exp int64, n int64) int64 {
	if exp == 0 {
		return 1
	}
	result := pow(base, exp/2, n)
	// result = result * result % n
	result = mul(result, result, n)
	if (exp & 1) == 1 {
		// return (result * base) % n
		return mul(result, base, n)
	}
	return result

}

func powMatrix(m Mat, exp int64, n int64) Mat {
	var result Mat
	if exp == 0 {
		return Mat{1, 0, 0, 1}
	}
	result = powMatrix(m, exp/2, n)
	result = result.multiply(result, n)
	if (exp & 1) == 1 {
		return result.multiply(m, n)
		// return m.multiply(result, n)
	}
	return result
}

func inverseModulo(b int64, n int64) int64 {
	return pow(b, n-2, n)
}

type Mat struct {
	a, b, c, d int64
}

func cutM(m int64) Mat {
	var result Mat
	result.a = 1
	result.b = -m
	result.d = 1
	return result
}

func invCutM(m int64) Mat {
	var result Mat
	result.a = 1
	result.b = m
	result.d = 1
	return result
}

func Rev(n int64) Mat {
	var result Mat
	result.a = -1
	result.b = n - 1
	result.d = 1
	return result
}
func inc(num int64) Mat {
	var result Mat
	result.a = num
	result.d = 1
	return result
}
func invInc(num int64, n int64) Mat {
	var result Mat
	result.a = inverseModulo(num, n)
	result.d = 1
	return result
}
func (m Mat) multiply(m2 Mat, n int64) Mat {
	// var result Mat
	// result.a = ((m.a*m2.a+m.b*m2.c)%n + n) % n
	// result.b = ((m.a*m2.b+m.b*m2.d)%n + n) % n
	// result.c = ((m.c*m2.a+m.d*m2.c)%n + n) % n
	// result.d = ((m.c*m2.b+m.d*m2.d)%n + n) % n
	// return result
	var result Mat
	// result.a = ((m.a*m2.a%n+m.b*m2.c%n)%n + n) % n
	// result.b = ((m.a*m2.b%n+m.b*m2.d%n)%n + n) % n
	// result.c = ((m.c*m2.a%n+m.d*m2.c%n)%n + n) % n
	// result.d = ((m.c*m2.b%n+m.d*m2.d%n)%n + n) % n
	a1 := mul(m.a, m2.a, n)
	a2 := mul(m.b, m2.c, n)
	b1 := mul(m.a, m2.b, n)
	b2 := mul(m.b, m2.d, n)
	c1 := mul(m.c, m2.a, n)
	c2 := mul(m.d, m2.c, n)
	d1 := mul(m.c, m2.b, n)
	d2 := mul(m.d, m2.d, n)
	result.a = ((a1+a2)%n + n) % n
	result.b = ((b1+b2)%n + n) % n
	result.c = ((c1+c2)%n + n) % n
	result.d = ((d1+d2)%n + n) % n
	return result

}
func dealIntoNewStack(cards []int) []int {
	for left, right := 0, len(cards)-1; left < right; left, right = left+1, right-1 {
		cards[left], cards[right] = cards[right], cards[left]
	}
	return cards
}
func cutNCards(cards []int, n int) []int {
	m := len(cards)
	if n < 0 {
		n = m + n
	}
	if n > 0 {
		tmp := cards[0:n]
		cards = cards[n:]
		cards = append(cards, tmp...)
		// cards[m-n:] = tmp

	}
	return cards
}
func dealWithN(cards []int, n int) []int {
	m := len(cards)
	newCards := make([]int, m)
	cnt := 0
	for i := 0; i < m; i++ {
		newCards[cnt] = cards[i]
		cnt = (cnt + n) % m
	}
	return newCards
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 22
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	// filePath := fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	// ll := strings.Split(lines[0], ",")
	// code := make([]int64, len(ll))
	// for i := range ll {
	// 	code[i], _ = strconv.ParseInt(ll[i], 10, 64)
	// }
	numCards := n2

	// cards := make([]int64, numCards)
	// var i int64
	// for i = 0; i < numCards; i++ {
	// 	cards[i] = i
	// }
	var m Mat
	m.a = 1
	m.d = 1
	var mi Mat
	mi.a = 1
	mi.d = 1
	for _, l := range lines {
		toks := strings.Split(l, " ")
		if toks[0] == "cut" {

			nn, _ := strconv.Atoi(toks[1])
			fmt.Println(nn)
			// cards = cutNCards(cards, nn)
			// m = (cutM(mnt64(nn))).multiply(m, int64(numCards))
			mi = mi.multiply(invCutM(int64(nn)), int64(numCards))
		} else if toks[1] == "into" {
			// cards = dealIntoNewStack(cardsm
			// m = (Rev(int64(numCards))).multiply(m, int64(numCards))
			mi = mi.multiply(Rev(int64(numCards)), int64(numCards))
		} else if toks[1] == "with" {
			nn, _ := strconv.Atoi(toks[3])
			// cards = dealWithN(cards, nn)
			// m = (inc(int64(nn))).multiplyMmm int64(numCards))
			mi = mi.multiply(invInc(int64(nn), int64(numCards)), int64(numCards))
		}
	}
	mip := powMatrix(mi, nn2, n2)
	// for i := range cards {
	// 	if cards[i] == 2019 {
	// 		fmt.Println(cards[i], " == ", i)
	// 	}
	// }
	// fmt.Println("Matrix:", (m.a*2019+m.b)%int64(numCards))
	fmt.Println("Matrix:", (mip.a*2020+mip.b)%int64(numCards))
	// fmt.Println(dealIntoNewStack(cards))
	// fmt.Println(cutNCards(cards, -4))
	// fmt.Println(dealWithN(cards, 3))
	elapsed := time.Since(start)
	fmt.Println(n2)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)

	fmt.Println(mul(129, 38, 1000))

}
