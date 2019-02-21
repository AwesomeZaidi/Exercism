package main

import (
	"fmt"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var acronym string

	acronym = string(s[0])

	for pos, char := range s {
		if string(char) == " " {
			acronym = acronym + string(s[pos+1])
		}
	}
	acronym = strings.ToUpper(acronym)
	return acronym
}

func main() {
	fmt.Println(Abbreviate("asim is cool"))
}
