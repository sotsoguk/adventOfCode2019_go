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
