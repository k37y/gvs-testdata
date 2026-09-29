# callback-dispatch

Scanner fixture for [gvs](https://github.com/k37y/gvs).

Scan in manual mode with `-library example.com/vulnerable -symbols Danger -fixversion v1.1.0`.
The synthetic dependency is replaced by the checked-in `dep` module, so no module downloads are needed.
The version in each application's `go.mod` controls version comparison.

Modules: .: callback, v1.0.0.

The `gvs_integration_excluded` build tag intentionally creates unknown reachability.
Reflection is reported as reachable by RTA; other algorithms still report reflection-risk evidence.
Static analysis does not follow interface dispatch or callback calls.
