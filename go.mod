module github.com/k37y/gvs-testdata

go 1.22.0

require (
	github.com/k37y/gvs-testdata/helper v0.0.0
	golang.org/x/crypto v0.23.0
	golang.org/x/net v0.23.0
)

require golang.org/x/sys v0.20.0 // indirect

replace github.com/k37y/gvs-testdata/helper => ./helper
