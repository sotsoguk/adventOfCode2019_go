package mathUtils

func Min(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}
func Min32(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}
func Max32(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func Minxyz(args ...int64) (minValue int64) {
	minValue = args[0]
	for _, v := range args {
		if v < minValue {
			minValue = v
		}
	}
	return
}

func Maxxyz(args ...int64) (maxValue int64) {
	maxValue = args[0]
	for _, v := range args {
		if v > maxValue {
			maxValue = v
		}
	}
	return
}

func Swap(x, y int) (int, int) {
	return y, x
}

func Abs32(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func Gcd(x, y int) int {
	for y != 0 {

		x, y = y, x%y
	}
	return x
}

func MakeDigits(n int64) []int64 {
	digits := make([]int64, 5)
	for i := 0; i < 5; i++ {
		digits[4-i] = int64(n % 10)
		// digits = append(digits, int(n%10))
		n /= 10
	}
	return digits
}
func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
