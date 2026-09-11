# multi-module

Two independent Go modules in one repo:
- `svc-a/`: `golang.org/x/net@v0.23.0`, `golang.org/x/crypto@v0.23.0` (both vulnerable)
- `svc-b/`: `golang.org/x/net@v0.33.0`, `golang.org/x/crypto@v0.31.0` (both patched)

Each is `package main`, calling `html.Parse` and `ssh.NewServerConn`.

Used for two CVEs: CVE-2024-45338 (GO-2024-3333, `golang.org/x/net/html`) and
CVE-2024-45337 (GO-2024-3321, `golang.org/x/crypto/ssh`).

Validates per-module job dispatch/deduplication across multiple `go.mod` modules in a
single repo, aggregated as vulnerable if any module is vulnerable.

Expected: `IsVulnerable=true` overall (driven by `svc-a`) for both CVEs; `svc-b` reports
its patched versions individually in `UsedImports`.
