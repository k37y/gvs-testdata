# vuln-untidy-gomod

CVE-2024-45338 (GO-2024-3333): `golang.org/x/net/html`, vulnerable range `0 → v0.33.0`.

Root `go.mod` requires `golang.org/x/net v0.23.0` directly, and also requires a local
nested module `helper/` (own `go.mod`, no `main` package) which itself requires
`golang.org/x/net v0.33.0` (patched).

Go's MVS would actually build with `v0.33.0`, but gvs reads the root `go.mod`'s literal
`require` line (`v0.23.0`) — a known, intentional false positive (safer than a false
negative; running `go mod tidy` resolves the discrepancy).

Expected: `IsVulnerable=true` (false positive by design), `Files={".": [["main.go"]], "helper": null}`, `Errors=[]`.
