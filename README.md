# package-identity-external

Calls both `net.LookupCNAME` and `Parser.Answer` from an imported
`golang.org/x/net/dns/dnsmessage` package. The external package comes from a
checked-in synthetic module selected through a local replacement.

Run from this branch with a built GVS `cg` binary:

```sh
for algo in rta vta cha static; do
  cg -algo "$algo" -library golang.org/x/net/dns/dnsmessage \
    -symbols Parser.Answer -fixversion v0.56.0 .
done
```

Expected for all four algorithms:

- `IsVulnerable`: `"true"` for the supplied manual target and version.
- `UsedImports["."]["golang.org/x/net/dns/dnsmessage"]` contains
  `Parser.Answer`, `CurrentVersion: "v0.26.0"`, and fix commands.
- Graphs refer to the external parser, without Go's distinct
  `vendor/golang.org/x/net/dns/dnsmessage` target.
- No scan errors.

The synthetic parser method is empty and has no vulnerability. The package
name and version threshold are test inputs, not an advisory assessment.
This fixture needs no external Go dependency downloads.

The companion `package-identity-go-bundled` branch calls only `net` and must
not produce an external-parser finding.
