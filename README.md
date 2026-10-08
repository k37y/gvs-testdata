# package-identity-go-bundled

Calls `net.LookupCNAME`, which uses Go's bundled DNS parser. The declared
`golang.org/x/net` dependency is replaced with a local synthetic module but
is never imported by the application.

Run from this branch with a built GVS `cg` binary:

```sh
for algo in rta vta cha static; do
  cg -algo "$algo" -library golang.org/x/net/dns/dnsmessage \
    -symbols Parser.Answer -fixversion v0.56.0 .
done
```

Expected for all four algorithms:

- `IsVulnerable`: `"false"`.
- `UsedImports`: empty/null; the external DNS package is not loaded.
- No graph paths or scan errors.
- No external module version or upgrade suggestion for Go's bundled parser.

The synthetic dependency has no vulnerability. Its name and version are test
inputs for package identity matching. This manual scan does not assess `net`
or establish that the application's Go toolchain is free of vulnerabilities.

The companion `package-identity-external` branch imports and calls the local
parser as well, and must retain its external-module finding.
