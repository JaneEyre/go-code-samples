package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	if path == "" {
		// Create an error with errors.New()
		return nil, errors.New("path is empty")
	}
	f, err := os.Open(path)
	if err != nil {
		// Wrap the error.
		// If the format string uses %w to format the error,
		// fmt.Errorf() returns an error that has the
		// method "func Unwrap() error" implemented.
		return nil, fmt.Errorf("open failed: %w", err)
	}
	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}
	return buf, nil
}
