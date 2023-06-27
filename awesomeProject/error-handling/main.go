package main

import (
	"errors"
	"fmt"
	"io/fs"
)

func main() {
	// Read a single file

	_, err := ReadFile("no/file")
	fmt.Println("Error:", err)

	// Unwrap the error returned by os.Open()
	fmt.Println("err unwrapped:", errors.Unwrap(err))

	// Read multiple files

	// Passing an empty slice triggers a plain error
	_, err = ReadFiles([]string{})
	fmt.Println("Error:", err)

	// Passing multiple paths of non-existing files triggers a joined error
	_, err = ReadFiles([]string{"no/file/a", "no/file/b", "no/file/c"})
	fmt.Println("Error:", err)

	// Unwrap the errors inside the joined error
	// A joined error does not have the method "func Unwrap() error"
	// because it does not wrap a single error but rather a slice of errors.
	// Therefore, errors.Unwrap() cannot unwrap the errors and returns nil.
	fmt.Println("err unwrapped:", errors.Unwrap(err))

	// To unwrap a joined error, you can type-assert that err has
	// an Unwrap() method that returns a slice of errors.
	e, ok := err.(interface{ Unwrap() []error })
	if ok {
		fmt.Println("err unwrapped after type assertion:", e.Unwrap())
	}

	// Confirm that the error is, or wraps, an fs.ErrNotExist error
	fmt.Println("err is fs.ErrNotExist:", errors.Is(err, fs.ErrNotExist))

	// Confirm that the error is, or wraps, an fs.PathError.
	// errors.As() assigns the unwrapped PathError to target.
	// This allows reading PathError's Path field.
	target := &fs.PathError{}
	if errors.As(err, &target) {
		fmt.Printf("err as PathError: path is '%s'\n", target.Path)
	}
}
