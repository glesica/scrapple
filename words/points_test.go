package words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	for _, c := range []struct {
		word  string
		value int
	}{
		{"", 0},
		{"a", 1},
		{"bad", 6},
	} {
		assert.Equal(t, c.value, Score(Word{Text: c.word}))
	}
}
