package hamming
import (
	"errors"
)

func Distance(a, b string) (int, error) {
	// Error Handling
	var count = 0
	if len(a) != len(b) {
		err := errors.New("Error in the DNA strands")
		return 0, err
	}
	// for each letter in string a
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	
	return count, nil
}


// PROGRAM PSUEDO CODE:
//  for each letter in string b
	// if letters are same
		// increment count
	
	// return count