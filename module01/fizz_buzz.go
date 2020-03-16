package module01

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	fizz = 1
	buzz = 2
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	var s strings.Builder

	for i := 1; i <= n; i++ {
		m := 0

		if i%3 == 0 {
			m |= fizz
		}
		if i%5 == 0 {
			m |= buzz
		}

		switch m {
		case fizz:
			s.WriteString("Fizz")
		case buzz:
			s.WriteString("Buzz")
		case fizz | buzz:
			s.WriteString("Fizz Buzz")
		default:
			s.WriteString(strconv.Itoa(i))
		}

		if i < n {
			s.WriteString(", ")
		}
	}

	fmt.Println(s.String())
}
