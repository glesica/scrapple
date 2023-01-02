package query

import "strings"

type T []string

// New converts a string to a slice of single-character
// strings representing the letters in the string. A blank
// is converted to a period for convenience (it reads better
// in debugging output). All letters are also converted to
// upper case.
//
// Example: "a b.d" -> "A.B.D"
func New(letters string) T {
	letters = strings.ToUpper(letters)
	letters = strings.ReplaceAll(letters, " ", ".")
	return strings.Split(letters, "")
}
