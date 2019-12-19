package main

import (
	"fmt"

	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

var fftPattern = []int{0, 1, 0, -1}

func getLastDigit(num int) int {
	return mathUtils.Abs32(num) % 10
}

func addSeq(start int, length int) []int {
	adds := make([]int, 0)
	curr := start
	cycle := (start + 1) * 4
	// fmt.Println(cycle)
	reps := start + 1
	for curr < length {
		for i := 0; i < reps; i++ {
			if curr+i >= length {
				break
			} else {
				adds = append(adds, curr+i)
			}
		}
		curr += cycle
	}
	return adds
}
func subSeq(start int, length int) []int {
	subs := make([]int, 0)
	curr := (start+1)*3 - 1
	cycle := (start + 1) * 4
	// fmt.Println(cycle)
	reps := start + 1
	for curr < length {
		for i := 0; i < reps; i++ {
			if curr+i >= length {
				break
			} else {
				subs = append(subs, curr+i)
			}
		}
		curr += cycle
	}
	return subs
}

func fft(signal *[]int, fftLength int) int {
	index := 0
	result := 0
	for index < len(*signal) {
		indexFFT := ((index + 1) / fftLength) % 4
		if fftPattern[indexFFT] == 1 {
			result += (*signal)[index]
		} else if fftPattern[indexFFT] == -1 {
			result -= (*signal)[index]
		}
		index++
	}
	return getLastDigit(result)
}
func toInt(input []int, start, end int) int {
	var result int
	for i := start; i <= end; i++ {
		// fmt.Println(i, input[i])
		result = result*10 + input[i]
	}
	return result
}
func part1(input []int) int64 {
	temp := make([]int, len(input))
	var result int64

	for run := 0; run < 100; run++ {
		for index := 1; index <= len(input); index++ {
			temp[index-1] = fft(&input, index)
		}
		swp := make([]int, len(input))
		copy(swp, temp)
		copy(temp, input)
		copy(input, swp)
	}
	//compute output
	for i := 0; i < 8; i++ {
		result = result*10 + int64(input[i])
	}
	return result
}
func phase2(input *[]int, output *[]int, offset int) {
	// out := make([]int, len(*input))
	// fmt.Println(len(input), len(input), len(out))

	// offset := toInt(input, 0, 6)
	// fmt.Println("OFF:", offset)
	for i := offset; i < len(*input); i++ {
		tmp := 0
		if i == offset {
			for j := i; j < len(*input); j++ {
				tmp += (*input)[j]
			}
			(*output)[i] = tmp % 10
		} else {
			(*output)[i] = (10 + (*output)[i-1] - (*input)[i-1]) % 10
		}
	}
	// return out
}
func part2(input []int) int {
	// signal2 := make([]int, 0)
	start := time.Now()
	// for i := 0; i < 10000; i++ {
	// 	signal2 = append(signal2, input...)
	// }
	l := len(input)
	signal2 := make([]int, len(input)*10000)
	for i := 0; i < 10000; i++ {
		copy(signal2[i*l:(i+1)*l], input)
	}
	out := make([]int, len(signal2))
	elapsed := time.Since(start)
	fmt.Println("create:", elapsed)
	// out := make([]int, len(signal2))
	// fmt.Println(len(input), len(signal2), len(out))
	offset := toInt(input, 0, 6)
	// fmt.Println("OFF:", offset)
	// for i := offset; i < len(signal2); i++ {
	// 	tmp := 0
	// 	if i == offset {
	// 		for j := i; j < len(signal2); j++ {
	// 			tmp += signal2[j]
	// 		}
	// 		out[i] = tmp % 10
	// 	} else {
	// 		out[i] = mathUtils.Abs32(10+out[i-1]-signal2[i-1]) % 10
	// 	}
	// }
	// swp := make([]int, len(out))
	for i := 0; i < 100; i++ {
		phase2(&signal2, &out, offset)

		// copy(swp, out)
		// copy(out, signal2)
		// copy(signal2, swp)
		copy(signal2[offset:], out[offset:])
	}

	return toInt(signal2, offset, offset+7)
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input16_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 16
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO
	// phases := 100
	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	number := lines[0]
	// number := "03036732577212944063491565474664"
	// number := "12345678"
	// number := "80871224585914546619083218645595"
	// number := "19617804207202209144916044189917"
	// ll := strings.Split(lines[0], ",")
	// fmt.Println(calcOre(lines, 1))
	// test := make([]int, 0)
	// test2 := []int{1, 2, 3}
	// for i := 0; i < 3; i++ {
	// 	test = append(test, test2...)
	input := make([]int, len(number))
	for i, d := range number {
		input[i] = int(d - '0')
	}
	//fmt.Println(toInt(input, 0, 6))
	// solution1 = part1(input)
	fmt.Println(part2(input))
	// }
	// fmt.Println(test)
	// input2 := make([]int, len(number))
	// fmt.Println(len(input2))
	// for i, d := range number {
	// 	input2[i] = int(d - '0')
	// }
	// input := make([]int, 0)
	// for i := 0; i < 1000; i++ {
	// 	input = append(input, input2...)
	// }
	// fmt.Println("Input made")
	// // fmt.Println(input)
	// // fmt.Println(addSeq(2, len(input)))
	// // fmt.Println(subSeq(2, len(input)))
	// // precompute all sequences
	// allAdds := make([][]int, len(input))
	// fmt.Println("allAdss made")
	// allSubs := make([][]int, len(input))
	// fmt.Println("allSubs made")
	// allPhases := make([][]int, phases+1)
	// allPhases[0] = input
	// for i := 0; i < len(input); i++ {

	// 	allAdds[i] = addSeq(i, len(input))
	// 	allSubs[i] = subSeq(i, len(input))
	// }
	// // fmt.Println(allAdds)
	// // fmt.Println(allSubs)
	// fmt.Println("PreComp Done")
	// for p := 1; p <= phases; p++ {
	// 	fmt.Println(p)
	// 	currInput := allPhases[p-1]
	// 	currOutput := make([]int, len(input))
	// 	for d := 0; d < len(input); d++ {
	// 		currDigit := 0
	// 		// fmt.Print(p, d, ": ", allAdds[d], allSubs[d])
	// 		for a := 0; a < len(allAdds[d]); a++ {
	// 			currDigit += currInput[allAdds[d][a]]
	// 		}
	// 		for s := 0; s < len(allSubs[d]); s++ {
	// 			currDigit -= currInput[allSubs[d][s]]

	// 		}
	// 		// fmt.Println(" => ", currDigit)
	// 		currOutput[d] = mathUtils.Abs32(currDigit) % 10
	// 	}
	// 	allPhases[p] = currOutput
	// }
	// fmt.Println(allPhases[phases])
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)

}
