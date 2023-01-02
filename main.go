package main

import (
	"fmt"
	"github.com/glesica/scrapple/query"
	"github.com/glesica/scrapple/words"
	"os"
)

func main() {
	w := words.NewDict()

	var h query.Hand
	var q query.T

	if len(os.Args) > 2 {
		h = query.NewHand(os.Args[1])
		q = query.New(os.Args[2])
	} else if len(os.Args) > 1 {
		q = query.New(os.Args[1])
	} else {
		fmt.Printf("usage: scrapple [hand] <board>")
	}

	var matches []words.Word

	if h != nil {
		matches = query.ForHand(q, h, w)
	} else {
		matches = query.All(q, w)
	}

	for _, m := range matches {
		score := words.Score(m)

		if m.Definition == "" {
			fmt.Printf("%d\t%s\n", score, m.Text)
		} else {
			fmt.Printf("%d\t%s\t%s\n", score, m.Text, m.Definition)
		}
	}
}
