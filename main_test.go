package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	run(&sb)
	out := sb.String()
	if !strings.Contains(out, "hello, skylence") {
		t.Fatalf("missing greeting in output: %q", out)
	}
	if !strings.Contains(out, "version: dev") {
		t.Fatalf("missing version line in output: %q", out)
	}
}
