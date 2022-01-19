package strings

import "strings"

// Lines returns a slice of lines in str.
func Lines(str string) []string {
	return strings.Split(str, "\n")
}
