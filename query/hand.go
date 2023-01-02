package query

import (
	"github.com/glesica/scrapple/words"
	"strings"
)

// TODO: Support blank tiles throughout

type Hand map[string]int

func NewHand(letters string) Hand {
	hand := make(Hand)

	for _, letter := range New(letters) {
		if _, exists := hand[letter]; exists {
			hand[letter] += 1
		} else {
			hand[letter] = 1
		}
	}

	return hand
}

func ForHand(q T, hand Hand, dict words.Dict) words.Dict {
	var results words.Dict

	for _, word := range dict {
		wordLen := len(word.Text)
		wordChars := strings.Split(word.Text, "")
		queryLen := len(q)

		// Start the word at the left-most end of the query, then shift
		// it right until it hits the right-most end of the query. At
		// each position, check if they match, counting characters along
		// the way. If they match, verify that the character counts are
		// satisfied by the hand. If so, add the word to the result set.
		for queryIndex := 0; queryIndex <= queryLen-wordLen; queryIndex++ {
			// We can't start a word immediately after a fixed letter so
			// just skip to the next position in that case.
			if queryIndex > 0 && q[queryIndex-1] != "." {
				continue
			}

			// We can't run a word into a fixed letter on the right-hand
			// side, either, so skip to the next position in that case.
			if queryIndex < queryLen-wordLen && q[queryIndex+wordLen] != "." {
				continue
			}

			matched := true
			counts := make(map[string]int)

			// Determine whether the word, in this position, matches the
			// query, and count how many of each letter would be required
			// to play the word.
			for wordIndex := 0; wordIndex < wordLen; wordIndex++ {
				queryLetter := q[queryIndex+wordIndex]
				wordLetter := wordChars[wordIndex]

				// If we found a wildcard, so we need to use a letter
				// the hand to fill it in with the character found in
				// the word at this position.
				if queryLetter == "." {
					if _, exists := counts[wordLetter]; exists {
						counts[wordLetter] += 1
					} else {
						counts[wordLetter] = 1
					}

					continue
				}

				// If we don't have a wildcard, and the characters don't
				// match, then this alignment is invalid.
				if q[queryIndex+wordIndex] != wordChars[wordIndex] {
					matched = false
					break
				}
			}

			if !matched {
				continue
			}

			// Check counts
			sufficient := true
			for letter, required := range counts {
				available, exists := hand[letter]

				// The letter is missing from the hand
				if !exists {
					sufficient = false
					break
				}

				// Too many of the letter are required
				if available < required {
					sufficient = false
					break
				}
			}

			if sufficient {
				results = append(results, word)
				// Only include a word in the result set once.
				// TODO: Continue on the outer loop instead of just breaking here
				break
			}
		}
	}

	return results
}
