package main

import (
	"errors"
	"fmt"
)

func ReadFiles(paths []string) ([][]byte, error) {
	var errs error
	var contents [][]byte

	if len(paths) == 0 {
		// Create a new error with fmt.Errorf() (but without using %w):
		return nil, fmt.Errorf("no paths provided: paths slice is %v", paths)
	}

	for _, path := range paths {
		content, err := ReadFile(path)
		if err != nil {
			// Join all errors that occur into errs.
			// The returned error type implements method "func Unwrap() []error".
			// (Note that the return type is a slice.)
			errs = errors.Join(errs, fmt.Errorf("reading %s failed: %w", path, err))
			continue
		}
		contents = append(contents, content)
	}

	// Some files may have been read, some may have failed to be read.
	// Therefore, ReadFiles returns both return values, regardless
	// of whether there have been errors.
	return contents, errs
}
