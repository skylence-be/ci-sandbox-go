// Command ci-sandbox-go prints a greeting. Logic lives in the app package so
// this entrypoint stays a thin, inspectable shell (excluded from the coverage
// gate); app.Run carries everything testable.
package main

import (
	"os"

	"github.com/skylence-be/ci-sandbox-go/app"
)

// version is stamped by goreleaser via -ldflags "-X main.version=..." and handed
// to the app package at startup.
var version = "dev"

func main() {
	if err := app.Run(os.Stdout, version); err != nil {
		os.Exit(1)
	}
}
