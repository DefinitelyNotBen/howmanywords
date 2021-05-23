package main

import (
	"fmt"
	"strings"
)

var punct = []rune{'!', '?', '.', ','}

func main() {
	// insert string into here
	s := "How many strings-ret hello And then there were 5"

	fmt.Print(howMany(s))

}

// check if character is one of sentence ending punctuation
func isPunct(r rune) bool {
	for _, p := range punct {
		if p == r {
			return true
		}
	}
	return false
}

// check if character is a letter
func isLetter(r rune) bool {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return true
	}

	return false
}

// counts words in sentence that contain only letters or are hyphenated. Sentence ends with full stop, comma, excalation or question marks
func howMany(sentence string) int {

	w := strings.Fields(sentence)
	n := len(w)

	for i := 0; i < n; i++ {

		r := []rune(w[i])

		word := true

		// loop over every character in word, except last
		for j := 0; j < len(r)-1; j++ {
			if !isLetter(r[j]) && r[j] != '-' {
				word = false
				break
			}
		}

		// if all characters are letters so far, last character isn't a letter but is punct then end sentence
		lastChar := r[len(r)-1]
		if word && !isLetter(lastChar) {
			if isPunct(lastChar) {

				w = w[:i+1]
				break
			} else {
				word = false
			}
		}

		// if not a word then remove from slice, unless final word
		if !word {
			if i == n-1 {
				w = w[:i]
				break
			}

			w = append(w[:i], w[i+1:]...)
			n -= 1
			continue
		}

	}

	return len(w)
}
