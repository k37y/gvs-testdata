# patched-single-range

Same structure as `vuln-single-range` but with fixed dependency versions:
`golang.org/x/net@v0.33.0`, `golang.org/x/crypto@v0.31.0`.

Expected: `IsVulnerable=false` for CVE-2024-45338 (GO-2024-3333) and CVE-2024-45337 (GO-2024-3321).
