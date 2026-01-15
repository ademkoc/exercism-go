// Package twofer creates sharing message
package twofer

import "fmt"

// ShareWith returns sharing phrase for person
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
