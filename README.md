# gvs-testdata

Test fixtures for [gvs](https://github.com/k37y/gvs) integration tests.

Each branch contains a minimal Go module designed to trigger a specific vulnerability scenario.

## Branches

| Branch | CVE | Type | Ranges | Expected |
|--------|-----|------|--------|----------|
| `vuln-single-range` | CVE-2024-45338 (GO-2024-3333) | non-stdlib (`golang.org/x/net/html`) | 0→0.33.0 | vulnerable |
| `patched-single-range` | CVE-2024-45338 (GO-2024-3333) | non-stdlib (`golang.org/x/net/html`) | 0→0.33.0 | not vulnerable |
| `vuln-stdlib-multi-range` | CVE-2023-45288 (GO-2024-2687) | stdlib (`net/http`) | 0→1.21.9, 1.22.0-0→1.22.2 | vulnerable |
| `patched-stdlib-multi-range` | CVE-2023-45288 (GO-2024-2687) | stdlib (`net/http`) | 0→1.21.9, 1.22.0-0→1.22.2 | not vulnerable |
| `vuln-replace-directive` | CVE-2024-45338 (GO-2024-3333) | non-stdlib (`golang.org/x/net/html`) | 0→0.33.0 | vulnerable (replace) |
| `multi-module` | CVE-2024-45338 (GO-2024-3333) | non-stdlib (`golang.org/x/net/html`) | 0→0.33.0 | svc-a vuln, svc-b patched |
| `not-a-go-repo` | — | — | — | error |
