// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

const sure = "Sure."
const chill = "Whoa, chill out!"
const fine = "Fine. Be that way!"
const whatever = "Whatever."

const minLetter, maxLetter = 'a', 'z'

func hasLetters(s *string) bool {
	toLower := strings.ToLower(*s)
	for _, l := range toLower {
		if l <= maxLetter && l >= minLetter {
			return true
		}
	}
	return false
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	switch {
	case strings.ToUpper(remark) == remark && hasLetters(&remark):
		return chill
	case remark == "":
		return fine
	case strings.HasSuffix(remark, "?"):
		return sure
	default:
		return whatever
	}
}
