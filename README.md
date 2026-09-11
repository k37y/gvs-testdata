# patched-stdlib-between-ranges

Same CVE as `vuln-stdlib-multi-range`. `go.mod` toolchain directive: `go 1.21.9`
(the fix version of range 1, in the gap before range 2 starts at `1.22.0`).

Validates the gap between two vulnerable ranges is correctly reported as not vulnerable.

Expected: `IsVulnerable=false`.
