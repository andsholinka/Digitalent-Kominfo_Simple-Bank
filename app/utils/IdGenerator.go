package utils

import "math/rand"

// RangeIn is func for
func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
