# Scabble

Answer interesting questions that come up while playing Scrabble.

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

Then, you can just do `go build -o scabble main.go`.

## Usage

`./scabble <hand> <board>` - find words from the given hand that can
fit in the given board location

`./scabble <board>` - find words that can fit in the given board
location, regardless of hand

## Examples

`./scabble treg '.i...'` - reveals that you can play "tiger", among
other words

`./scabble '.....'` - displays all available five-letter words

`./scabble hello` - shows how many points "hello" is worth
