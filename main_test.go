package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamgcampbell/go-wordle-solver/internal/guesser"
)

func TestWithStandardDictionary(t *testing.T) {
	r := strings.NewReader(db)
	g := guesser.NewIOGuesser(r)

	tests := map[string]struct {
		expectedGuesses []string
		responses       []string
	}{
		"Day 219 Test": {
			expectedGuesses: []string{"arose", "clout", "loony", "knoll"},
			responses:       []string{"xxgxx", "xygxx", "yxgyx", "ggggg"},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			step := 0
			for {
				guess, valid := g.NextGuess()
				if !valid {
					continue
				}
				require.Equal(t, test.expectedGuesses[step], guess)
				ng, done := processResponse(guess, test.responses[step], g)
				if done {
					break
				}
				g = ng
				step++
			}
		})
	}
}
