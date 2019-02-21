package pangram

import (
	"strings"
)

// IsPangram function runs
func IsPangram(s string) bool {

	// coded with @dy-fi! 💯
	// edge case
	if len(s) < 26 {
		return false
	}

	// faster lookup 👀
	alphabet := make(map[string]bool)
	s = strings.ToLower(s)

	for _, v := range s {
		for i := 'a'; i >= 'a' && i <= 'z'; i++ {

			if v == i {
				alphabet[string(i)] = true
			}
		}
	}

	for _, v := range alphabet {
		if v == false || len(alphabet) < 26 {
			return false // ☹️
		}
	}

	return true // 💫
}

// letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
// 	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
