// Package reverse contains method String.
package reverse

// String function has logic to reverse a given string.
func String(s string) string {

	// edge case
	if len(s) == 0 {
		return ""
	}

	// rune slice
	runeSlice := []rune(s)

	// empty reverse string we return
	var reverse string

	// loop through rune string backwards
	for i := len(runeSlice) - 1; i >= 0; i-- {
		// add it to the reverse
		reverse += string(runeSlice[i])
	}

	return reverse
}
