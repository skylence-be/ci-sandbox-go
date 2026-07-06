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

// run writes the program output to w; main stays a thin, untested shell.
func run(w io.Writer) {
	fmt.Fprintln(w, greet.Greeting("skylence"))
	fmt.Fprintln(w, "version:", version)
}

func main() {
	run(os.Stdout)
}
