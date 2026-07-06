package main

import (
	"io"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	if err := run(&sb); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	out := sb.String()
	if !strings.Contains(out, "hello, skylence") {
		t.Fatalf("missing greeting in output: %q", out)
	}
	if !strings.Contains(out, "version: dev") {
		t.Fatalf("missing version line in output: %q", out)
	}
}

// errWriter fails every write so the error path in run is covered.
type errWriter struct{}

func (errWriter) Write([]byte) (int, error) { return 0, io.ErrShortWrite }

func TestRunWriteError(t *testing.T) {
	t.Parallel()
	if err := run(errWriter{}); err == nil {
		t.Fatal("expected a write error, got nil")
	}
}
