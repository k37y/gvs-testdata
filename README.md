# gvs-testdata

Test fixtures for [gvs](https://github.com/k37y/gvs) integration tests.

Each branch contains a minimal Go module designed to trigger a specific vulnerability scenario.
All non-stdlib branches import both `golang.org/x/net/html` and `golang.org/x/crypto/ssh` to enable testing multiple CVEs per branch.

## Branches

| Branch | CVEs | Type | Expected |
|--------|------|------|----------|
| `vuln-single-range` | CVE-2024-45338 (`x/net` v0.23.0), CVE-2024-45337 (`x/crypto` v0.23.0) | non-stdlib | vulnerable |
| `patched-single-range` | CVE-2024-45338 (`x/net` v0.33.0), CVE-2024-45337 (`x/crypto` v0.31.0) | non-stdlib | not vulnerable |
| `vuln-stdlib-multi-range` | CVE-2023-45288 (`net/http`, Go 1.21.4) | stdlib | vulnerable |
| `patched-stdlib-multi-range` | CVE-2023-45288 (`net/http`, Go 1.22.5) | stdlib | not vulnerable |
| `vuln-replace-directive` | CVE-2024-45338 (`x/net` v0.23.0→v0.24.0), CVE-2024-45337 (`x/crypto` v0.23.0) | non-stdlib | vulnerable (replace) |
| `multi-module` | CVE-2024-45338, CVE-2024-45337 | non-stdlib | svc-a vuln, svc-b patched |
| `not-a-go-repo` | — | — | error |
