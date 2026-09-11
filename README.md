# vuln-indirect-dep

CVE-2024-45338 (GO-2024-3333): `golang.org/x/net/html`, vulnerable range `0 → v0.33.0`.

Root module (`main.go`) imports a local nested module `wrapper/` (own `go.mod`,
`replace .../wrapper => ./wrapper`, `package wrapper` — no `main` package). `wrapper`
requires `golang.org/x/net@v0.23.0` and calls `html.Parse`.

Validates:
- indirect vulnerable usage through a local module boundary is detected;
- the no-`main` `wrapper` module is skipped by the job dispatcher without producing
  spurious "No entry points found" errors.

Expected: `IsVulnerable=true`, `Files={".": [["main.go"]], "wrapper": null}`, `Errors=[]`.
