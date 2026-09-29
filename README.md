# gvs-testdata

Scenario fixtures for [GVS](https://github.com/k37y/gvs), a Go vulnerability scanner.
This is the catalog of all **45 scenario branches**. `main` holds this catalog;
`expand-fixture-coverage` is a documentation branch, not a scanner fixture.

Expected results describe the expanded GVS scanner test suite. Older scanner
versions may reproduce the bugs these fixtures were created to detect.

## Reading the expected results

The tables use the scanner's `IsVulnerable` field:

- `true`: the selected algorithm finds a reachable affected symbol and the version is vulnerable.
- `false`: the scan finds no affected path, or the selected dependency version is outside the affected range.
- `unknown`: the scan cannot establish a verdict, for example because package loading is incomplete or a replacement has an unrelated version history.

These are separate from API task statuses such as completed, failed, and cancelled.
A `false` result for an unresolved dynamic call does not establish that the call
cannot happen. See the algorithm matrix below.

Unless stated otherwise, results use RTA on Linux/macOS with the fixture's default
build constraints. Synthetic reachability and scan-logic cases exercise all four
algorithms where listed below; aggregation cases use RTA. The Go standard-library
cases check the Go version recorded by the scanner from the fixture's go.mod.

## Advisory targets

| Target used below | CVE / Go advisory | Affected package | Fixed versions checked |
| --- | --- | --- | --- |
| x/net | CVE-2024-45338 / GO-2024-3333 | golang.org/x/net/html | v0.33.0 |
| x/crypto | CVE-2024-45337 / GO-2024-3321 | golang.org/x/crypto/ssh | v0.31.0 |
| net/http | CVE-2023-45288 / GO-2024-2687 | net/http | Go 1.21.9 and 1.22.2 |
| gRPC | GO-2023-2153 | google.golang.org/grpc and google.golang.org/grpc/internal/transport | v1.56.3, v1.57.1, and v1.58.3, according to the affected range |

CVE scans fetch live advisory data. GVS keeps expected package/symbol snapshots
under `internal/api/testdata/advisories` to detect upstream changes. Those snapshots
are assertion data, not scanner input.

## Real dependency, version, and repository scenarios

| Branch | Scan target | Expected result | Scenario |
| --- | --- | --- | --- |
| [`vuln-single-range`](https://github.com/k37y/gvs-testdata/tree/vuln-single-range) | x/net and x/crypto | Both `true` | Direct calls with both modules at v0.23.0. |
| [`patched-single-range`](https://github.com/k37y/gvs-testdata/tree/patched-single-range) | x/net and x/crypto | Both `false` | x/net v0.33.0 and x/crypto v0.31.0; reachable calls at fixed versions. |
| [`vuln-stdlib-multi-range`](https://github.com/k37y/gvs-testdata/tree/vuln-stdlib-multi-range) | net/http | `true` | Go directive 1.21.4; first affected range, fixed in 1.21.9. |
| [`vuln-stdlib-second-range`](https://github.com/k37y/gvs-testdata/tree/vuln-stdlib-second-range) | net/http | `true` | Go directive 1.22.1; second affected range, fixed in 1.22.2. |
| [`patched-stdlib-between-ranges`](https://github.com/k37y/gvs-testdata/tree/patched-stdlib-between-ranges) | net/http | `false` | Go directive 1.21.9, at the first fix and before the second affected range. |
| [`patched-stdlib-multi-range`](https://github.com/k37y/gvs-testdata/tree/patched-stdlib-multi-range) | net/http | `false` | Go directive 1.22.5, beyond the relevant fixed versions. |
| [`vuln-multi-range`](https://github.com/k37y/gvs-testdata/tree/vuln-multi-range) | gRPC | `true` | gRPC v1.57.0; validates the v1.57.1 fix and findings in both the root and internal/transport packages. |
| [`patched-multi-range`](https://github.com/k37y/gvs-testdata/tree/patched-multi-range) | gRPC | `false` | gRPC v1.57.1, exactly at the fix for that release line. |
| [`patched-multi-range-between`](https://github.com/k37y/gvs-testdata/tree/patched-multi-range-between) | gRPC | `false` | gRPC v1.56.3, between affected ranges. |
| [`vuln-replace-directive`](https://github.com/k37y/gvs-testdata/tree/vuln-replace-directive) | x/net and x/crypto | Both `true` | x/net v0.23.0 replaced by v0.24.0; the replacement remains vulnerable. x/crypto v0.23.0 is also vulnerable. |
| [`patched-replace-directive`](https://github.com/k37y/gvs-testdata/tree/patched-replace-directive) | x/net and x/crypto | Both `false` | x/net v0.23.0 replaced by fixed v0.33.0; x/crypto is v0.31.0. |
| [`vuln-indirect-dep`](https://github.com/k37y/gvs-testdata/tree/vuln-indirect-dep) | x/net | `true` | Vulnerable x/net v0.23.0 reached through a local wrapper module. |
| [`vuln-build-constraint`](https://github.com/k37y/gvs-testdata/tree/vuln-build-constraint) | x/net and x/crypto | x/net `unknown`; x/crypto `true` | On Linux/macOS, the x/net call is excluded by a Windows build tag and needs manual analysis; the SSH call remains included. |
| [`vuln-untidy-gomod`](https://github.com/k37y/gvs-testdata/tree/vuln-untidy-gomod) | x/net and x/crypto | Both `false` | Declarations say v0.23.0, but package loading selects x/net v0.33.0 and x/crypto v0.31.0 through the helper dependency. Guards against the former false positive. |
| [`multi-module`](https://github.com/k37y/gvs-testdata/tree/multi-module) | x/net and x/crypto | Both overall `true` | svc-a uses vulnerable versions; svc-b uses fixed versions. Checks per-module versions, findings, and fix commands. |
| [`not-a-go-repo`](https://github.com/k37y/gvs-testdata/tree/not-a-go-repo) | x/net | `unknown` | No Go module or analyzable Go entry points; Files and UsedImports are empty. |
| [`replacement-downgrade`](https://github.com/k37y/gvs-testdata/tree/replacement-downgrade) | Manual: x/net/html, Parse, v0.33.0 | `true` | Requires x/net v0.33.0 but replaces it with v0.24.0. The suggested fix must update the replacement. |
| [`fork-replacement`](https://github.com/k37y/gvs-testdata/tree/fork-replacement) | Manual: jwt-go, Parse, v3.2.1+incompatible | `unknown` | github.com/dgrijalva/jwt-go v3.2.0+incompatible is replaced by github.com/golang-jwt/jwt v3.2.2+incompatible. Different module histories prevent a reliable upstream-version comparison; no automatic fix is suggested. |

The two manual replacement cases use these arguments:

```sh
# replacement-downgrade
cg -algo rta -library golang.org/x/net/html -symbols Parse -fixversion v0.33.0 .

# fork-replacement: the threshold is a manual test input, not a claim about a fork fix.
cg -algo rta -library github.com/dgrijalva/jwt-go -symbols Parse -fixversion v3.2.1+incompatible .
```

`vuln-build-constraint` uses `//go:build windows`; its documented unknown result
assumes that file is excluded. The Windows variant is not part of this expectation.

## Synthetic reachability and scan-logic scenarios

These fixtures use a checked-in `example.com/vulnerable` module with a local
replacement. `Danger` and `Other` are synthetic targets, not real vulnerabilities.
These scenarios require no external Go dependency downloads.

The default manual scan is:

```sh
cg -algo rta -library example.com/vulnerable -symbols Danger -fixversion v1.1.0 .
```

| Branch | Expected result | Scenario |
| --- | --- | --- |
| [`reachability-direct`](https://github.com/k37y/gvs-testdata/tree/reachability-direct) | `true` | Direct Danger call with dependency v1.0.0. Also reused for commit checkout and open-ended version-range checks. |
| [`reachability-patched`](https://github.com/k37y/gvs-testdata/tree/reachability-patched) | `false` | Direct Danger call with dependency v1.1.0, exactly at the fix. |
| [`unreachable-symbol`](https://github.com/k37y/gvs-testdata/tree/unreachable-symbol) | `false` | Danger is referenced only in an uncalled function. |
| [`test-only-symbol`](https://github.com/k37y/gvs-testdata/tree/test-only-symbol) | `false` | Danger is called only from a _test.go file. |
| [`dependency-not-imported`](https://github.com/k37y/gvs-testdata/tree/dependency-not-imported) | `false` | The dependency is required by go.mod but is not imported. |
| [`interface-dispatch`](https://github.com/k37y/gvs-testdata/tree/interface-dispatch) | Algorithm-dependent; see matrix | Danger is reached through an interface method. |
| [`callback-dispatch`](https://github.com/k37y/gvs-testdata/tree/callback-dispatch) | Algorithm-dependent; see matrix | Danger is reached through a function parameter. |
| [`reflection-call`](https://github.com/k37y/gvs-testdata/tree/reflection-call) | Algorithm-dependent; see matrix | reflect.ValueOf(Danger).Call(nil) in the main package; checks reflection evidence. |
| [`reflection-helper`](https://github.com/k37y/gvs-testdata/tree/reflection-helper) | Algorithm-dependent; see matrix | The reflection call lives in a helper package; checks that its source location and evidence are included. |
| [`unsafe-call`](https://github.com/k37y/gvs-testdata/tree/unsafe-call) | `true` | Direct Danger call plus unsafe usage; requires the unsafe flag. |
| [`init-call`](https://github.com/k37y/gvs-testdata/tree/init-call) | `true` | Danger is called during package initialization. |
| [`goroutine-call`](https://github.com/k37y/gvs-testdata/tree/goroutine-call) | `true` | Danger is called inside a goroutine. |
| [`deferred-call`](https://github.com/k37y/gvs-testdata/tree/deferred-call) | `true` | Danger is invoked with defer. |
| [`generic-call`](https://github.com/k37y/gvs-testdata/tree/generic-call) | `true` | Danger is reached through an instantiated generic function. |
| [`selected-dependency-version`](https://github.com/k37y/gvs-testdata/tree/selected-dependency-version) | `false` | go.mod declares v1.0.0, while a helper causes Go to select patched v1.1.0. |
| [`prerelease-version`](https://github.com/k37y/gvs-testdata/tree/prerelease-version) | `true` | v1.1.0-rc.1 is before the v1.1.0 fix. |
| [`pseudo-version`](https://github.com/k37y/gvs-testdata/tree/pseudo-version) | `true` | v1.0.1-0.20260101000000-abcdefabcdef is before the v1.1.0 fix. |
| [`broken-package`](https://github.com/k37y/gvs-testdata/tree/broken-package) | `unknown` | A valid main package exists alongside a package with a type error. Incomplete analysis must not report false. |
| [`missing-dependency`](https://github.com/k37y/gvs-testdata/tree/missing-dependency) | `unknown` | A missing local dependency prevents package loading; requires an error diagnostic. |
| [`multi-symbol-paths`](https://github.com/k37y/gvs-testdata/tree/multi-symbol-paths) | `true` | Scan Other,Danger. Each SVG must contain its reported symbol and the corresponding alpha→Danger or beta→Other edge. |

The reachability cases, including initialization, goroutines, deferred calls,
generics, and reflection helpers, run under RTA, VTA, CHA, and static analysis.
The broken-package, missing-dependency, selected-version, prerelease, pseudo-version,
and symbol-to-graph cases also run under all four algorithms.

`broken-package` and `missing-dependency` are intentionally not buildable.
For `multi-symbol-paths`, replace `-symbols Danger` with `-symbols Other,Danger`
and add `-graph ./graphs` to generate the SVGs checked by the API tests.

An additional scan of `reachability-direct` checks a vulnerability with no known
fixed version. It must report `true` without inventing a fix command:

```sh
cg -algo rta -library example.com/vulnerable -symbols Danger -fixversion 'Introduced in 1.0.0 - ' .
```

### Algorithm expectations for dynamic calls

| Fixture pattern | RTA | VTA | CHA | Static |
| --- | --- | --- | --- | --- |
| Direct calls, initialization, goroutines, defer, generics | true | true | true | true |
| Interface dispatch | true | true | true | false |
| Callback dispatch | true | true | true | false |
| Tested reflection call, including the helper-package variant | true | false | false | false |

These values apply to the vulnerable synthetic fixtures and their tested call
patterns. All four algorithms must report reflection-risk evidence for the two
reflection fixtures, including the symbol, package, confidence, location, and
`reflect.ValueOf(Danger)` evidence. Those risks are informational and currently do
not change the vulnerability verdict. RTA is not guaranteed to resolve every
possible reflection or runtime-selected call.

## Multi-module aggregation scenarios

Each module has its own go.mod and source files. Vulnerable modules use v1.0.0;
patched modules use v1.1.0. Unknown modules contain a vulnerable call behind the
`gvs_integration_excluded` build tag. Leave that tag unset.

| Branch | Module a | Module b | Overall result |
| --- | --- | --- | --- |
| [`multi-module-all-patched`](https://github.com/k37y/gvs-testdata/tree/multi-module-all-patched) | Patched | Patched | `false` |
| [`multi-module-vulnerable-patched`](https://github.com/k37y/gvs-testdata/tree/multi-module-vulnerable-patched) | Vulnerable | Patched | `true` |
| [`multi-module-patched-unknown`](https://github.com/k37y/gvs-testdata/tree/multi-module-patched-unknown) | Patched | Unknown | `unknown` |
| [`multi-module-unknown-patched`](https://github.com/k37y/gvs-testdata/tree/multi-module-unknown-patched) | Unknown | Patched | `unknown` |
| [`multi-module-vulnerable-unknown`](https://github.com/k37y/gvs-testdata/tree/multi-module-vulnerable-unknown) | Vulnerable | Unknown | `true` |
| [`multi-module-unknown-vulnerable`](https://github.com/k37y/gvs-testdata/tree/multi-module-unknown-vulnerable) | Unknown | Vulnerable | `true` |
| [`multi-module-all-unknown`](https://github.com/k37y/gvs-testdata/tree/multi-module-all-unknown) | Unknown | Unknown | `unknown` |

These cases check verdict precedence (`true` over `unknown` over `false`),
per-module versions and fixes, and graph URLs that do not overwrite one another.
They use the default synthetic manual scan above with RTA.

## Additional checks in GVS

The fixture branches are inputs to several checks, rather than one test per branch.
The expanded suite lives in
[`internal/api/integration_test.go`](https://github.com/k37y/gvs/blob/main/internal/api/integration_test.go).

| Test group | Scenarios and assertions |
| --- | --- |
| `TestCallgraphIntegration` | CVE and Go-ID scans, manual scans, advisory ranges, direct/indirect/replaced dependencies, multiple modules, build constraints, no-Go repositories, and selected algorithm cases. Checks versions, fixes, symbols, and result fields. |
| `TestCallgraphMatrixIntegration` | Additional x/crypto CVEs and manual x/net scans; patched, replaced, indirect, excluded, untidy, and multi-module combinations across additional algorithms. |
| `TestCallgraphFixturesIntegration` | Synthetic reachability under four algorithms, unsafe/reflect flags, reflection evidence, seven aggregation combinations, and a pinned commit checkout of reachability-direct (`f426435e1da63f0dd405068fbe7a571ff06d9876`). |
| `TestCallgraphScanLogicIntegration` | Incomplete package loading, selected versions, prereleases, pseudo-versions, unfixed ranges, replacement downgrades, fork replacements, and exact symbol/call-edge checks in SVGs. |
| `TestCallgraphConcurrencyIntegration` | Fresh scans with 1, 4, and 8 workers. Multiple modules and multiple affected gRPC packages exercise the actual scanner built with `-race`. |
| `TestCallgraphLifecycleIntegration` | Cache reuse and isolation by repository, revision, CVE, algorithm, library, symbol, and fix version; clone/branch/commit failures and recovery; concurrent-request rejection; progress streaming; cancellation during Go package loading; scanner/child exit; preserved cancelled status; cancellation must not cache a result. |
| `TestCgBinaryValidation` | Actual scanner exit status and diagnostics for missing arguments, invalid directories, and invalid algorithms. |

Shared API assertions check advisory identities and exact affected package/symbol
sets, used-module associations, expected graph presence, distinct graph URLs, and
successful downloads of valid SVG documents. The multi-symbol fixture additionally
checks actual graph endpoints and edges.

Lifecycle tests create private temporary Git repositories so they can modify
inputs and control failures without changing shared scenario branches.

Focused unit regressions in
[`pkg/cmd/cg/scanner_test.go`](https://github.com/k37y/gvs/blob/main/pkg/cmd/cg/scanner_test.go)
cover version introduction/fix boundaries, prereleases, pseudo-versions (including
v0.0.0 pseudo-versions with an introduced-zero range), multiple and open-ended
ranges, advisory transport failures, timeouts, HTTP errors, malformed/trailing JSON,
truncated responses, independent worker verdicts, and cancellation returning unknown.
[`cmd/cg/main_test.go`](https://github.com/k37y/gvs/blob/main/cmd/cg/main_test.go)
checks that symbol sorting/deduplication preserves the matching path in each module.
These checks use local inputs or injected failures and do not need additional
fixture branches.

## Running the scenarios

Install/build `cg` from the GVS repository and put it on PATH. Check out a scenario
branch before scanning; `main` itself contains documentation.

```sh
git clone --branch reachability-direct https://github.com/k37y/gvs-testdata.git
cd gvs-testdata
GVS_CLAUDE=0 cg -algo rta -library example.com/vulnerable -symbols Danger -fixversion v1.1.0 .

# Switch to a real CVE fixture.
git checkout vuln-single-range
GVS_CLAUDE=0 cg -algo rta CVE-2024-45338 .
```

From a **GVS checkout**, run the complete integration suite:

```sh
make test-integration
```

It requires Go, Git, Graphviz (`sfdp`), a C compiler for Go's race detector, and
network access for fixture clones, external dependencies, and live advisories.
The harness uses the real API handlers on an ephemeral server with isolated caches
and clone directories, and disables optional AI verification.

The integration harness always builds the **scanner subprocess** with `-race`.
The default unit command (`make test`) and integration API test process are not
race-instrumented. To enable those as well, run from the GVS checkout:

```sh
go test -race ./...
go test -race -tags integration ./internal/api -run 'TestCallgraph.*Integration|TestCgBinaryValidation' -timeout 45m
```

For unpublished fixture changes, set `GVS_TESTDATA_REPO` to a local Git repository:

```sh
GVS_TESTDATA_REPO=/absolute/path/to/gvs-testdata make test-integration
```

Every scenario used by that run must exist as a local branch in the source
repository; remote-tracking refs alone are not sufficient for local branch cloning.
Use a targeted `go test -run` selection when validating a single new scenario.
For repeatable bug reports, record the fixture commit, GVS revision, Go version,
algorithm, and scan arguments.
