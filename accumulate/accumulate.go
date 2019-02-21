package accumulate

// Accumulate returns all caps string
func Accumulate(s []string, f func(string) string) []string {

	// runeSlice := []rune(s)

	// for _, word := range runeSlice {
	// 	char := unicode.ToLower(char)
	// 	// if unicode.IsLetter(char) && runeSlice[word] {

	// 	// }
	// }

	// return runeSlice

	// var newSlice []string
	// r := []rune{s}
	for pos, _ := range s {
		// println(word)
		s[pos] = f(s[pos])
		// s = append(s, strings.ToUpper(word))
	}
	return s
}
