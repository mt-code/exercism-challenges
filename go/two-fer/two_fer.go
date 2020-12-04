// Simple package to test conditional statements
package twofer

import "fmt"

// Returns who "it" is being shared with, if the name is empty
// we share it generically.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
