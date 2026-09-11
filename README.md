# vuln-multi-range

GO-2023-2153: `google.golang.org/grpc`, vulnerable ranges `0 → v1.56.3`,
`v1.57.0 → v1.57.1`, `v1.58.0 → v1.58.3`. `go.mod` requires `grpc v1.57.0` (inside range 2).

Expected: `IsVulnerable=true`.
