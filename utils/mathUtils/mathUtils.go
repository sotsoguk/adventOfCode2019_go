package mathUtils

func Min(x, y int64) int64 {
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
