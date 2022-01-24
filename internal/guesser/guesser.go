package guesser

import (
	"bufio"
	"io"
)

type Guesser interface {
	// NextGuess returns the next wordle guess.
	NextGuess() (string, bool)
}

var _ Guesser = (GuesserFunc)(nil)

// GuesserFunc is a Guesser expressed as a function.
type GuesserFunc func() (string, bool)

// NextGuess implements Guesser.
func (g GuesserFunc) NextGuess() (string, bool) {
	return g()
}

// ioGuesser is a Guesser that receives its guesses from io input
type ioGuesser struct {
	scanner *bufio.Scanner
}

func NewIOGuesser(reader io.Reader) Guesser {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	return &ioGuesser{
		scanner: scanner,
	}
}

// NextGuess will return the next guess from the reader
func (i *ioGuesser) NextGuess() (string, bool) {
	if i.scanner.Scan() {
		return i.scanner.Text(), true
	}
	return "", false
}

func WithBadLetterRule(letter string, next Guesser) Guesser {
	return GuesserFunc(func() (string, bool) {
		word, valid := next.NextGuess()
		if !valid {
			return word, false
		}

		for _, wl := range word {
			if string(wl) == letter {
				return word, false
			}
		}

		return word, true
	})
}

func WithBadLetterPositionRule(letter string, pos int, next Guesser) Guesser {
	return GuesserFunc(func() (string, bool) {
		word, valid := next.NextGuess()
		if !valid {
			return word, false
		}

		for i, l := range word {
			if string(l) == letter && i == pos {
				return word, false
			}
		}

		return word, true
	})
}

func WithCorrectLetterRule(letter string, pos int, next Guesser) Guesser {
	return GuesserFunc(func() (string, bool) {
		word, valid := next.NextGuess()
		if !valid {
			return word, false
		}

		for i, l := range word {
			if string(l) == letter && i == pos {
				return word, true
			}
		}

		return word, false
	})
}

// WithGoodLetterBadPositionRule wraps a guesser function and adds validation for the "yellow"
// feedback on a wordle letter. This validation will ensure that all words have this letter, but
// they cannot contain the letter at this position.
func WithGoodLetterBadPositionRule(letter string, pos int, next Guesser) Guesser {
	return GuesserFunc(func() (string, bool) {
		word, valid := next.NextGuess()
		if !valid {
			return word, false
		}

		found := false
		for i, l := range word {
			if string(l) == letter {
				if i == pos {
					return word, false
				}
				found = true
			}
		}

		if found {
			return word, true
		}

		return word, false
	})
}

func WithGoodLetterBadPositionSkipPositionRule(letter string, pos int, skip int, next Guesser) Guesser {
	return GuesserFunc(func() (string, bool) {
		word, valid := next.NextGuess()
		if !valid {
			return word, false
		}

		found := false
		for i, l := range word {
			if string(l) == letter {
				if i == pos {
					return word, false
				}
				if i == skip {
					continue
				}
				found = true
			}
		}

		if found {
			return word, true
		}

		return word, false
	})
}
