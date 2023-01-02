package query

import (
	"github.com/glesica/scrapple/words"
	"regexp"
)

func All(q T, dict words.Dict) []words.Word {
	var results words.Dict

	expr := AsRegex(q)

	re, err := regexp.Compile(expr)
	if err != nil {
		panic(err)
	}

	for _, word := range dict {
		if re.MatchString(word.Text) {
			results = append(results, word)
		}
	}

	return results
}
