package guesser

import "testing"

var (
	alwaysPositiveGuesser = GuesserFunc(func() (string, bool) {
		return "test", true
	})
	alwaysNegativeGuesser = GuesserFunc(func() (string, bool) {
		return "test", false
	})
)

func TestWithGoodLetterBadPositionRule(t *testing.T) {
	tests := map[string]struct {
		guesser Guesser

		letter string
		pos    int
		wantA  string
		wantB  bool
	}{
		"Positive to negative because e should not be at position 1": {
			guesser: alwaysPositiveGuesser,
			letter:  "e",
			pos:     1,
			wantA:   "test",
			wantB:   false,
		},
		"Positive to negative because x is not present in the word": {
			guesser: alwaysPositiveGuesser,
			letter:  "x",
			pos:     1,
			wantA:   "test",
			wantB:   false,
		},
		"Positive to positive because e is present at a position other than 2": {
			guesser: alwaysPositiveGuesser,
			letter:  "e",
			pos:     2,
			wantA:   "test",
			wantB:   true,
		},
		"Negative to negative because it honors other rules": {
			guesser: alwaysNegativeGuesser,
			letter:  "e",
			pos:     2,
			wantA:   "test",
			wantB:   false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			guesser := WithGoodLetterBadPositionRule(test.letter, test.pos, test.guesser)
			actualA, actualB := guesser.NextGuess()
			if actualA != test.wantA {
				t.Errorf("Got: %s, Want: %s.", actualA, test.wantA)
			}

			if actualB != test.wantB {
				t.Errorf("Got: %t, Want: %t.", actualB, test.wantB)
			}
		})
	}
}

func TestWithCorrectLetterRule(t *testing.T) {
	tests := map[string]struct {
		guesser Guesser

		letter string
		pos    int
		wantA  string
		wantB  bool
	}{
		"Positive to negative because x should be at position 2": {
			guesser: alwaysPositiveGuesser,
			letter:  "x",
			pos:     2,
			wantA:   "test",
			wantB:   false,
		},
		"Positive to positive because e is at position 1": {
			guesser: alwaysPositiveGuesser,
			letter:  "e",
			pos:     1,
			wantA:   "test",
			wantB:   true,
		},
		"Negative to negative because it honors other rules, even though e is at position 1": {
			guesser: alwaysNegativeGuesser,
			letter:  "e",
			pos:     1,
			wantA:   "test",
			wantB:   false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			guesser := WithCorrectLetterRule(test.letter, test.pos, test.guesser)
			actualA, actualB := guesser.NextGuess()
			if actualA != test.wantA {
				t.Errorf("Got: %s, Want: %s.", actualA, test.wantA)
			}

			if actualB != test.wantB {
				t.Errorf("Got: %t, Want: %t.", actualB, test.wantB)
			}
		})
	}
}

func TestWithBadLetterRule(t *testing.T) {
	tests := map[string]struct {
		guesser Guesser

		letter string
		wantA  string
		wantB  bool
	}{
		"Invalid because 'e' should not be present in the word": {
			guesser: alwaysPositiveGuesser,
			letter:  "e",
			wantA:   "test",
			wantB:   false,
		},
		"Valid because 'x' is not present in the word": {
			guesser: alwaysPositiveGuesser,
			letter:  "x",
			wantA:   "test",
			wantB:   true,
		},
		"Invalid because it honors other rules, even though 'x' is not present": {
			guesser: alwaysNegativeGuesser,
			letter:  "x",
			wantA:   "test",
			wantB:   false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			guesser := WithBadLetterRule(test.letter, test.guesser)
			actualA, actualB := guesser.NextGuess()
			if actualA != test.wantA {
				t.Errorf("Got: %s, Want: %s.", actualA, test.wantA)
			}

			if actualB != test.wantB {
				t.Errorf("Got: %t, Want: %t.", actualB, test.wantB)
			}
		})
	}
}
