// Command ci-sandbox-go prints a greeting. It exists so the CI/CD pipelines
// have real code to lint, test, build, and release.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/skylence-be/ci-sandbox-go/greet"
)

// version is stamped by goreleaser via -ldflags "-X main.version=...".
var version = "dev"

// run writes the program output to w and returns the first write error, if any.
// main stays a thin shell that surfaces that error as a non-zero exit.
func run(w io.Writer) error {
	bw := bufio.NewWriter(w)
	fmt.Fprintln(bw, greet.Greeting("skylence"))
	fmt.Fprintln(bw, "version:", version)
	return bw.Flush()
}

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "ci-sandbox-go:", err)
		os.Exit(1)
	}
}
