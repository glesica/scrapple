package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeRegex(t *testing.T) {
	for _, c := range []struct {
		query string
		regex string
	}{
		{".a.a.", "^[A-Z]{0,1}A[A-Z]A[A-Z]{0,1}$"},
		{"..a.a..", "^[A-Z]{0,2}A[A-Z]A[A-Z]{0,2}$"},
		{"...a.a..", "^[A-Z]{0,3}A[A-Z]A[A-Z]{0,2}$"},
		{"a.a", "^A[A-Z]A$"},
		{"a a", "^A[A-Z]A$"},
		{"...", "^[A-Z]{0,3}$"},
	} {
		assert.Equal(t, c.regex, AsRegex(New(c.query)))
	}
}
