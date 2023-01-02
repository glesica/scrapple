package words

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed wordlist.txt
var wordsData string

const WordsPathVar = "SCABBLE_WORDS"

type Word struct {
	Text       string
	Definition string
}

type Dict []Word

func NewDict() Dict {
	var words []Word
	var wordStr string

	if wordPath, exists := os.LookupEnv(WordsPathVar); exists {
		wordBytes, err := os.ReadFile(wordPath)
		if err != nil {
			panic("word file not found")
		}

		wordStr = string(wordBytes)
	} else {
		wordStr = wordsData
	}

	for _, line := range strings.Split(wordStr, "\n") {
		parts := strings.Split(line, "\t")

		w := Word{
			Text: strings.ToUpper(parts[0]),
		}

		// Skip lines that don't parse to a word
		if w.Text == "" {
			continue
		}

		if len(parts) > 1 {
			w.Definition = strings.Join(parts[1:], " ")
		}

		words = append(words, w)
	}

	return words
}
