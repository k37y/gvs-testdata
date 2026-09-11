# vuln-single-range

- CVE-2024-45338 (GO-2024-3333): `golang.org/x/net/html`, vulnerable range `0 → v0.33.0`, module at `v0.23.0`.
- CVE-2024-45337 (GO-2024-3321): `golang.org/x/crypto/ssh`, vulnerable range `0 → v0.31.0`, module at `v0.23.0`.

`main.go` calls `html.Parse` and `ssh.NewServerConn` directly (single module, single main package).

Expected: `IsVulnerable=true` for both CVEs.
