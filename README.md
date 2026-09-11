# vuln-stdlib-multi-range

CVE-2023-45288 (GO-2024-2687): `net/http`, vulnerable ranges `0 → 1.21.9` and `1.22.0 → 1.22.2`.

`go.mod` toolchain directive: `go 1.21.4` (inside the first range).

Expected: `IsVulnerable=true`.
