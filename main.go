// Command ci-sandbox-go prints a greeting. It exists so the CI/CD pipelines
// have real code to lint, test, build, and release.
package main

import (
	"fmt"

	"github.com/skylence-be/ci-sandbox-go/greet"
)

// version is stamped by goreleaser via -ldflags "-X main.version=...".
var version = "dev"

func main() {
	fmt.Println(greet.Greeting("skylence"))
	fmt.Println("version:", version)
}
