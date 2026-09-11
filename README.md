# vuln-replace-directive

CVE-2024-45338 (GO-2024-3333): `golang.org/x/net/html`, vulnerable range `0 → v0.33.0`.

`go.mod` requires `golang.org/x/net v0.23.0` with `replace golang.org/x/net => golang.org/x/net v0.24.0`.

Validates the scanner resolves the `replace`d version (`v0.24.0`), not the `require` line.

Expected: `IsVulnerable=true` (v0.24.0 is still below the v0.33.0 fix).
