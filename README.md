# ci-sandbox-go

Minimal Go project used to validate the Skylence CI shape on GitHub-hosted
runners. The private product repos run the same staged pipeline (lint, tests,
vulnerability scan, workflow lint) on self-hosted hardware; this sandbox
exists so the hosted-runner variant has visible, public run results.

Not a product. Expect force-pushes and churn.
