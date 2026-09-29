# gvs-testdata

Test fixtures for [gvs](https://github.com/k37y/gvs) integration tests.

Each branch contains a minimal Go module designed to trigger a specific vulnerability scenario.
All non-stdlib branches import both `golang.org/x/net/html` and `golang.org/x/crypto/ssh` to enable testing multiple CVEs per branch.

## Branches

| Branch | CVEs | Type | Expected |
|--------|------|------|----------|
| `vuln-single-range` | CVE-2024-45338 (`x/net` v0.23.0), CVE-2024-45337 (`x/crypto` v0.23.0) | non-stdlib | vulnerable |
| `patched-single-range` | CVE-2024-45338 (`x/net` v0.33.0), CVE-2024-45337 (`x/crypto` v0.31.0) | non-stdlib | not vulnerable |
| `vuln-stdlib-multi-range` | CVE-2023-45288 (`net/http`, Go 1.21.4) | stdlib | vulnerable |
| `patched-stdlib-multi-range` | CVE-2023-45288 (`net/http`, Go 1.22.5) | stdlib | not vulnerable |
| `vuln-replace-directive` | CVE-2024-45338 (`x/net` v0.23.0→v0.24.0), CVE-2024-45337 (`x/crypto` v0.23.0) | non-stdlib | vulnerable (replace) |
| `multi-module` | CVE-2024-45338, CVE-2024-45337 | non-stdlib | svc-a vuln, svc-b patched |
| `not-a-go-repo` | — | — | error |

## Reachability and aggregation fixtures

These branches use a small checked-in `example.com/vulnerable` module and a local
`replace` directive. They require no dependency downloads and isolate call-graph
behavior from changes to external libraries. Scan them with:

```sh
cg -library example.com/vulnerable -symbols Danger -fixversion v1.1.0 -algo rta .
```

| Branch | Expected result | Purpose |
|--------|-----------------|---------|
| `reachability-direct` | true | Direct call with version v1.0.0 |
| `reachability-patched` | false | Direct call with version v1.1.0 |
| `unreachable-symbol` | false | Old dependency, vulnerable function only in an uncalled function |
| `test-only-symbol` | false | Vulnerable function only used in a test file |
| `dependency-not-imported` | false | Old dependency present in go.mod but never imported |
| `interface-dispatch` | true for RTA/VTA/CHA; false for static | Call through an interface |
| `callback-dispatch` | true for RTA/VTA/CHA; false for static | Call through a function parameter |
| `reflection-call` | true for RTA; false for VTA/CHA/static | Reflection call; all algorithms report reflection-risk evidence |
| `unsafe-call` | true | Direct vulnerable call plus unsafe usage |
| `multi-module-all-patched` | false | Both modules patched |
| `multi-module-vulnerable-patched` | true | One vulnerable module and one patched module |
| `multi-module-patched-unknown` | unknown | Patched module plus an excluded source file |
| `multi-module-unknown-patched` | unknown | Same aggregation with module order reversed |
| `multi-module-vulnerable-unknown` | true | Vulnerability takes precedence over unknown reachability |
| `multi-module-unknown-vulnerable` | true | Same aggregation with module order reversed |
| `multi-module-all-unknown` | unknown | Both modules contain excluded vulnerable calls |

Unknown-reachability fixtures use the `gvs_integration_excluded` build tag. Leave
that tag unset when running these checks.

The tests in `gvs/internal/api/integration_test.go` exercise these branches through
the HTTP API and all four call-graph algorithms, including response fields,
reflection evidence, and downloadable SVG graphs. Private temporary repositories
are used only for lifecycle checks such as cancellation and clone failures.
