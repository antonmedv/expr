package runtime

import "math"

// Compare the integral part before examining the fraction, so converting a
// large integer to float64 cannot round it into equality. The bounds keep
// float-to-integer conversions within their defined range.
func compareIntFloat(x int64, y float64) (int, bool) {
	if math.IsNaN(y) {
		return 0, false
	}
	if y >= 0x1p63 {
		return -1, true
	}
	if y < -0x1p63 {
		return 1, true
	}
	whole := int64(y)
	if x < whole {
		return -1, true
	}
	if x > whole {
		return 1, true
	}
	return compareFraction(float64(whole), y), true
}

func compareUintFloat(x uint64, y float64) (int, bool) {
	if math.IsNaN(y) {
		return 0, false
	}
	if y >= 0x1p64 {
		return -1, true
	}
	if y < 0 {
		return 1, true
	}
	whole := uint64(y)
	if x < whole {
		return -1, true
	}
	if x > whole {
		return 1, true
	}
	return compareFraction(float64(whole), y), true
}

func compareFraction(whole, value float64) int {
	if whole < value {
		return -1
	}
	if whole > value {
		return 1
	}
	return 0
}
