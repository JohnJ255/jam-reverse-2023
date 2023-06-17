package helper

import "math"

func KmphToPixelsPerTick(kmph float64) float64 {
	return kmph / 22
}

func ToRadians(angle Degrees) float64 {
	return float64(angle) * math.Pi / 180
}
