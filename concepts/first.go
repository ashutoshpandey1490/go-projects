package concepts

import (
	"errors"
	"fmt"
	// or we can run 'go get <package>' which will update the mod file.
)

// In Go, a function whose name starts with a capital letter can be called by
// a function not in the same package. This is known in Go as an exported name
func Print_my_str(name string) string {
	// In Go, the := operator is a shortcut for declaring and initializing a
	// variable in one line (Go uses the value on the right to determine the variable's type).
	message := fmt.Sprintf("Hello, %v, Welcome !!", name)
	return message
}

// error is a build-in type
func Print_my_str_with_error(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name provided")
	}
	message := fmt.Sprintf("Hello, %v, Welcome !!", name)
	return message, nil
}
