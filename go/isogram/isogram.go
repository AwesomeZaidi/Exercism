package isogram

import (
	"strings"
)

// IsIsogram is a function woo waa.
func IsIsogram(word string) bool {
	// declare empty arr
	if word == "" {
		return true
	}
	var letters []string

	// initialize empty letters array
	// for every character in the word
	// 		go through every element in the array of letters
	// 			if the character in the word already exists in letters array
	// 				return false
	// 		append the character in the word to the letters array

	for i := 0; i < len(word); i++ {
		char := strings.ToLower(string(word[i]))
		// check if the letter is in the letters array already
		for j := 0; j < len(letters); j++ {
			letterChar := strings.ToLower(string(letters[j]))
			if letterChar == char {
				return false
			}
		}
		if char == "-" || char == " " {
			continue
		}
		letters = append(letters, char)
	}
	// check if that char already exists in arr
	// return False
	// store the char in an arr

	return true
}
