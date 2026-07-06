// Command ci-sandbox-go prints a greeting. It exists so the CI/CD pipelines
// have real code to lint, test, build, and release.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/skylence-be/ci-sandbox-go/greet"
)

// version is stamped by goreleaser via -ldflags "-X main.version=...".
var version = "dev"

// run writes the program output to w and returns any write error. main stays a
// thin shell that surfaces that error as a non-zero exit.
func run(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s\nversion: %s\n", greet.Greeting("skylence"), version)
	return err
}

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "ci-sandbox-go:", err) //nolint:errcheck // best-effort stderr diagnostic on the exit path
		os.Exit(1)
	}
}
