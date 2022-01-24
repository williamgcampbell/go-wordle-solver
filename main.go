package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/williamgcampbell/go-wordle-solver/internal/guesser"
)

const winningResponse = "ggggg"

//go:embed words.txt
var db string

func main() {
	r := strings.NewReader(db)
	g := guesser.NewIOGuesser(r)
	buf := bufio.NewReader(os.Stdin)
	for {
		guess, valid := g.NextGuess()

		if len(guess) == 0 {
			fmt.Println("all out of guesses...")
			break
		}

		if !valid {
			continue
		}

		fmt.Println(guess)
		resp := waitForResponse(buf)

		ng, done := processResponse(guess, resp, g)
		if done {
			fmt.Println("We did it!")
			break
		}

		g = ng
	}
}

func processResponse(guess, resp string, g guesser.Guesser) (guesser.Guesser, bool) {
	if resp == winningResponse {
		return g, true
	}

	for i, c := range resp {
		letter := string(guess[i])
		switch string(c) {
		case "g":
			g = guesser.WithCorrectLetterRule(letter, i, g)
		case "y":
			pos, exists := greenLetterPos(letter, guess, resp)
			if !exists {
				g = guesser.WithGoodLetterBadPositionRule(letter, i, g)
			} else {
				// A letter may be green and yellow in a single guess when guessing a word with two of the same letter.
				// For this reason we need a rule that omits the green position from this rule.
				g = guesser.WithGoodLetterBadPositionSkipPositionRule(letter, i, pos, g)
			}
		case "x":
			pos, exists := greenLetterPos(letter, guess, resp)
			if !exists {
				g = guesser.WithBadLetterRule(letter, g)
			} else {
				// A letter may appear green and grey in a single guess when guessing a word with two of the same
				// letter. For this reason we need to adjust the rule to isolate positions that aren't the correct
				// one.
				for p := range guess {
					if p != pos {
						g = guesser.WithBadLetterPositionRule(letter, p, g)
					}
				}
			}
		}
	}
	return g, false
}

func waitForResponse(buf *bufio.Reader) string {
	for {
		fmt.Print("> ")
		response, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		r := response[0 : len(response)-1]
		if err := validateResponse(r); err != nil {
			fmt.Println(err)
			continue
		}
		return r
	}
}

func validateResponse(s string) error {
	if len(s) != 5 {
		return fmt.Errorf("invalid response length")
	}

	for _, l := range s {
		if string(l) != "g" && string(l) != "y" && string(l) != "x" {
			return fmt.Errorf("invalid response format")
		}
	}

	return nil
}

func greenLetterPos(letter, guess, resp string) (int, bool) {
	for i := range guess {
		if string(guess[i]) == letter && string(resp[i]) == "g" {
			return i, true
		}
	}
	return 0, false
}
