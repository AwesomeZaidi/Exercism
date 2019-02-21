package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is the datastructure representing the frequency of words
type Frequency map[string]int

// better at var state so its more reusablre
var wordRegex = regexp.MustCompile(`[\w']+`)

// WordCount function can...
func WordCount(s string) Frequency {

	frequency := Frequency{}

	words := wordRegex.FindAllString(s, -1)

	for _, word := range words {
		frequency[strings.Trim(strings.ToLower(word), "'")]++
	}

	return frequency

}
