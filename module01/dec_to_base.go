package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	const chars = "0123456789ABCDEF"

	var s []byte
	for dec > 0 {
		s = append([]byte{chars[dec%base]}, s...)

		if dec < base {
			break
		}

		dec /= base
	}

	return string(s)
}
