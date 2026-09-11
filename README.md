# vuln-build-constraint

CVE-2024-45338 (GO-2024-3333): `golang.org/x/net/html`, vulnerable range `0 → v0.33.0`.

`constrained.go` has `//go:build windows` and is the only file calling `html.Parse`;
`main.go` only calls the unrelated `golang.org/x/crypto/ssh.NewServerConn`.

On a non-windows analysis host, `constrained.go` is excluded from the build, so
reachability of `html.Parse` cannot be statically determined.

Expected: `IsVulnerable=unknown`, `Errors` contains "Need manual analysis" referencing `constrained.go`.
