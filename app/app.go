// Package app holds the testable program logic behind the ci-sandbox-go binary.
package app

import (
	"fmt"
	"io"

	"github.com/skylence-be/ci-sandbox-go/greet"
)

// Run writes the greeting and version to w, returning any write error.
func Run(w io.Writer, version string) error {
	_, err := fmt.Fprintf(w, "%s\nversion: %s\n", greet.Greeting("skylence"), version)
	return err
}
