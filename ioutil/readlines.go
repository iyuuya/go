package ioutil

import (
	"errors"
	"io/ioutil"

	"github.com/iyuuya/go/strings"
)

// ReadLines reads the entire file specified by filename as individual lines, and returns those lines in a slice.
func ReadLines(filename string) ([]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	return strings.Lines(string(bytes)), nil
}
