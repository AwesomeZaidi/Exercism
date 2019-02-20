// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	acronym := ""
	fmt.Printf(s)
	for pos, char := range s {
		if char == " " {
			// acronym.append(s[pos+1])
			acronym += "csd"
		}
	}
	return acronym
}
