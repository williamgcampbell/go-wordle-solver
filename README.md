# Description
This is a simple Wordle solver, that works by making a guess and receiving the feedback in the form of a letter inputs. 
- g - the letter at this position is green
- y - the letter at this position is yellow
- x - the letter at this position is grey

A typical game may look something like this:
```
$ ./go-wordle-solver
arose
> xxgxx
clout
> xygxx
loony
> yxgyx
knoll
> ggggg
We did it!
```

# Dictionary
Because the IOGuesser iterates the provided dictionary of words sequentially, the efficacy of this program relies heavily on the construction of this dictionary.
The primary `words.txt` dictionary has gone through some pre-processing for this reason.

## 5-letter words
The original dictionary was filtered down to only 5-letter words.

## Ordered
The dictionary has been ordered from most valuable to least valuable words. 
The value of a word is determined in the following way:
1. Letters are assigned a weight based on the number of words the letter appears in. 
2. Words are assigned a weight based on the sum of all weights of its letters.
3. All words with duplicate letters are given less weight than all words with unique sets of letters because they limit the discoverability of a guess.

## Words not in the Wordle Dictionary
Occasionally a guess will be a word that isn't in the Wordle DB. 
In this case the word should be removed from the DB and the game restarted. 