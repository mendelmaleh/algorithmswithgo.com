package module01

import "math"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) (s int) {
	var charsMap = make(map[rune]int)
	const chars = "0123456789ABCDEF"

	for i, v := range chars {
		charsMap[v] = i
	}

	l, b := len(value), float64(base)
	for i, v := range value {
		s += charsMap[v] * int(math.Pow(b, float64(l-i-1)))
	}

	return s
}
