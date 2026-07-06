package app

import (
	"io"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	if err := Run(&sb, "1.2.3"); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	out := sb.String()
	for _, want := range []string{"hello, skylence", "version: 1.2.3"} {
		if !strings.Contains(out, want) {
			t.Fatalf("output %q missing %q", out, want)
		}
	}
}

// errWriter fails every write so Run's error path is covered.
type errWriter struct{}

func (errWriter) Write([]byte) (int, error) { return 0, io.ErrShortWrite }

func TestRunWriteError(t *testing.T) {
	t.Parallel()
	if err := Run(errWriter{}, "1.2.3"); err == nil {
		t.Fatal("expected a write error, got nil")
	}
}
