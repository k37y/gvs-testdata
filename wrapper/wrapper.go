package wrapper

import (
	"io"

	"golang.org/x/net/html"
)

func ParseHTML(r io.Reader) (*html.Node, error) {
	return html.Parse(r)
}
