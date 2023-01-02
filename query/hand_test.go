package query

import (
	"github.com/glesica/scrapple/words"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestForHand(t *testing.T) {
	hand := NewHand("bd")

	for _, c := range []struct {
		q      T
		hand   Hand
		words  words.Dict
		result words.Dict
	}{
		{New("a.c"), hand, wordList("abc"), wordList("abc")},
		{New("a.c"), hand, wordList("aec"), nil},
		{New("a.c"), hand, wordList("ab", "bc"), nil},
		{New("..."), NewHand("eey"), wordList("eye"), wordList("eye")},
	} {
		result := ForHand(c.q, c.hand, c.words)
		assert.Equal(t, c.result, result)
	}
}

func wordList(all ...string) words.Dict {
	var list words.Dict
	for _, word := range all {
		list = append(list, words.Word{Text: strings.ToUpper(word)})
	}
	return list
}
