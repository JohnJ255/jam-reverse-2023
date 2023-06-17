package helper

type numbers interface{ int | float64 }

func Limited[T numbers](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
