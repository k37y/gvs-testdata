# vuln-stdlib-second-range

Same CVE as `vuln-stdlib-multi-range`. `go.mod` toolchain directive: `go 1.22.1`
(inside the second range, `1.22.0 → 1.22.2`).

Validates the scanner checks all vulnerable ranges, not just the first.

Expected: `IsVulnerable=true`.
