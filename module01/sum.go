package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) (s int) {
	for _, v := range numbers {
		s += v
	}

	return
}
