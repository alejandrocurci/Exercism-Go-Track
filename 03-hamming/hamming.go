// hamming package implements operations for ADN sequences
package hamming

import "errors"

// Distance func calculates the amount of differences between two ADN sequence
func Distance(a, b string) (int, error) {
	var dif int
	var err error
	if len(a) != len(b) {
		err = errors.New("sequences have different lenghts")
		return dif, err
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dif++
		}
	}
	return dif, err
}
