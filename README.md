# Scrapple

Answer interesting questions that come up while playing Scrabble. It's
a "crappy" way to play Scrabble.

## Building

You'll need a word list. You can search for this, or provide your own.
The official word list is
[apparently](https://boardgames.stackexchange.com/questions/38366/latest-collins-scrabble-words-list-in-text-file)
copyrighted, so it is not included with the source code. The word list
is a text file where each line consists of a word and a definition,
separated by a single tab character. The definition is optional.

By default, the wordlist should be stored in a file called `wordlist.txt`
in the `words` directory, then the build should succeed. There is a
trivial sample list in the root directory called `sample-wordlist.txt`.

You can also set an environment variable, `SCRAPPLE_WORDS` to control the
word list at runtime.

Then, you can just do `go build -o scrapple main.go`.

## Usage

`./scrapple <hand> <board>` - find words from the given hand that can
fit in the given board location

`./scrapple <board>` - find words that can fit in the given board
location, regardless of hand

## Examples

`./scrapple treg '.i...'` - reveals that you can play "tiger", among
other words

`./scrapple '.....'` - displays all available five-letter words

`./scrapple hello` - shows how many points "hello" is worth
