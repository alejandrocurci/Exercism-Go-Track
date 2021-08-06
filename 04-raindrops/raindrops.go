package raindrops

import (
	"strconv"
)

func Convert(num int) string {
	var sound string
	if num%3 == 0 {
		sound += "Pling"
	}
	if num%5 == 0 {
		sound += "Plang"
	}
	if num%7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		strNum := strconv.Itoa(num)
		sound += strNum
	}
	return sound
}
