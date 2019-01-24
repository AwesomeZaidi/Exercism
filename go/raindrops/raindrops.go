package raindrops
import (
	"strconv"
)

func Convert(num int) (string) {
	var output = ""
	if num % 3 == 0 {
		output += "Pling"
	} else if num % 5 == 0 {
		output += "Plang" 
	} else if num % 7 == 0{
		output += "Plong"
	} else {
		output = strconv.Itoa(123)
	}
	return output
}

// Notes:
	// When you write in Go, especially when writing if's the compiler adds a semicolon
	// when every 'line' finishes so you have to have them on the same line!