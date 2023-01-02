package query

import (
	"fmt"
)

const letterRange = "[A-Z]"

func AsRegex(q T) string {
	leading := 0
	for _, c := range q {
		if c != "." && c != " " {
			break
		}
		leading++
	}

	trailing := 0
	// If we consumed the whole query with the leading matcher,
	// then the query is just wildcards, and we don't need
	// anything else.
	if leading < len(q) {
		for i := len(q) - 1; i >= 0; i-- {
			c := q[i]
			if c != "." && c != " " {
				break
			}
			trailing++
		}
	}

	expr := ""

	if leading > 0 {
		expr = fmt.Sprintf("%s{0,%d}", letterRange, leading)
	}

	for i := leading; i < len(q)-trailing; i++ {
		c := q[i]
		if c == "." || c == " " {
			c = letterRange
		}

		expr = fmt.Sprintf("%s%s", expr, c)
	}

	if trailing > 0 {
		expr = fmt.Sprintf("%s%s{0,%d}", expr, letterRange, trailing)
	}

	return fmt.Sprintf("^%s$", expr)
}
