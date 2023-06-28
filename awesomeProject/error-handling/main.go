package main

import (
	"errors"
	"io/fs"
	"log"
)

func main() {
	// Read a single file

	_, err := ReadFile("no/file")
	if err != nil {
		log.Println("Error:", err)
	}

	// Unwrap the error returned by os.Open()
	log.Println("err unwrapped:", errors.Unwrap(err))

	// Read multiple files

	// Passing an empty slice triggers a plain error
	_, err = ReadFiles([]string{})
	log.Println("Error:", err)

	// Passing multiple paths of non-existing files triggers a joined error
	_, err = ReadFiles([]string{"no/file/a", "no/file/b", "no/file/c"})
	log.Println("Error:", err)

	// Unwrap the errors inside the joined error
	// A joined error does not have the method "func Unwrap() error"
	// because it does not wrap a single error but rather a slice of errors.
	// Therefore, errors.Unwrap() cannot unwrap the errors and returns nil.
	log.Println("err unwrapped:", errors.Unwrap(err))

	// To unwrap a joined error, you can type-assert that err has
	// an Unwrap() method that returns a slice of errors.
	e, ok := err.(interface{ Unwrap() []error })
	if ok {
		log.Println("err unwrapped after type assertion:", e.Unwrap())
	}

	// Confirm that the error is, or wraps, an fs.ErrNotExist error
	log.Println("err is fs.ErrNotExist:", errors.Is(err, fs.ErrNotExist))

	// Confirm that the error is, or wraps, an fs.PathError.
	// errors.As() assigns the unwrapped PathError to target.
	// This allows reading PathError's Path field.
	target := &fs.PathError{}
	if errors.As(err, &target) {
		log.Printf("err as PathError: path is '%s'\n", target.Path)
	}

	// Recover from a panic

	// This example is at the end of main, because the panic
	// causes main to exit. Only the deferred function is
	// called before exiting.

	defer func() {
		// Is this func invoked from a panic?
		if r := recover(); r != nil {
			// Yes: recover from the panic
			log.Printf("Recovering from error '%v'\n", r)
			// ...
		}
	}()

	// isValidPath panics because of an invalid regexp.
	if isValidPath("/path/to/file") {
		_, _ = ReadFile("/path/to/file")
	}
}
