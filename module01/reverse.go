package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	w := []rune(word)
	l := len(w)
	r := make([]rune, l)

	for i := 0; i < l; i++ {
		r[i] = w[l-i-1]
	}

	return string(r)
}
