# ci-sandbox-go

Minimal Go project used to validate the Skylence CI shape on GitHub-hosted
runners. The private product repos run the same staged pipeline (lint, tests,
vulnerability scan, workflow lint) on self-hosted hardware; this sandbox
exists so the hosted-runner variant has visible, public run results.

## Local git hooks (lefthook)

`lefthook.yml` runs the same gates as CI, locally, before code leaves your
machine. Install once per clone (hooks live in `.git/hooks`, which git never
clones):

```
brew install lefthook   # or: go install github.com/evilmartians/lefthook@latest
lefthook install
```

- **pre-commit** (instant): `gofmt` check on staged files.
- **pre-push** (mirrors CI, parallel): `golangci-lint run`,
  `go test -race` with the 100% line gate, and `govulncheck`.

Prereqs on PATH: `go` (ships `gofmt`), `golangci-lint`, `govulncheck`. Skip the
coverage step ad hoc with `LEFTHOOK_EXCLUDE=coverage lefthook run pre-push`.

Not a product. Expect force-pushes and churn.

<!-- demo PR to exercise dependency-review + PR-title checks -->
