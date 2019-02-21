package raindrops

import (
	"strconv"
)

// Convert func does converting stuff
func Convert(num int) string {
	var output = ""
	if num%3 == 0 {
		output += "Pling"
	}
	if num%5 == 0 {
		output += "Plang"
	}
	if num%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(num)
	}
	return output
}

// Notes:
// When you write in Go, especially when writing if's the compiler adds a semicolon
// when every 'line' finishes so you have to have them on the same line!
