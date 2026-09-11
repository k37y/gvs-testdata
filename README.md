# patched-multi-range-between

Same CVE as `vuln-multi-range`. `go.mod` requires `grpc v1.56.3`
(fix version of range 1, in the gap before range 2 starts at `v1.57.0`).

Expected: `IsVulnerable=false`.
