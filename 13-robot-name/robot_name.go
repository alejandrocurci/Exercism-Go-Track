package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

type Robot struct {
	name string
}

// memory saves the names already used
var memory = map[string]bool{}

// Name returns the current robot's name
// if the robot has no name, it generates one new for it
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		newName := r.generateName()
		for !validateName(newName) {
			newName = r.generateName()
		}
		r.name = newName
	}
	return r.name, nil
}

// Reset resets the name's robot to the empty string
func (r *Robot) Reset() {
	r.name = ""
}

// validateName returns false when the name has been already used
func validateName(name string) bool {
	if !memory[name] {
		memory[name] = true
		return true
	}
	return false
}

// generateName returns a string with 2 uppercase letters and 3 numbers (AF173, BX920, etc)
func (r *Robot) generateName() string {
	var newName string
	// generate first two letters
	for i := 0; i < 2; i++ {
		newName += randomUpper()
	}
	// generate last three numbers
	for i := 0; i < 3; i++ {
		newName += strconv.Itoa(random(0, 10))
	}
	return newName
}

// random generates a random integer between [min, max)
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// randomUpper generates a random uppercase letter between [A, Z]
// "A" = 65 and "Z" = 90 in ASCII code
func randomUpper() string {
	num := random(65, 91)
	return string(rune(num))
}
