// Package twofer implements string utility for someone's name
package twofer

import "fmt"

// ShareWith returns a string depending on someone's name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
