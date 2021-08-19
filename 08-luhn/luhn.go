package luhn

import (
	"strconv"
	"strings"
)

// Valid returns whether a credit card number is valid or not
func Valid(number string) bool {
	// remove whitespaces
	stripped := strings.ReplaceAll(number, " ", "")
	// strings with less than 2 digits are not valid
	if len(stripped) < 2 {
		return false
	}
	// check whether every character is numerical
	_, err := strconv.Atoi(stripped)
	if err != nil {
		return false
	}
	var sum int
	for i, num := range stripped {
		// convert runes to the actual digit
		digit, _ := strconv.Atoi(string(num))
		// modify numbers in specific positions, depending on the string length
		if i%2 == len(stripped)%2 {
			double := digit * 2
			if double > 9 {
				sum += double - 9
			} else {
				sum += double
			}
		} else {
			sum += digit
		}
	}
	return sum%10 == 0
}
