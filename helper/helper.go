package helper

import (
	"strings"

	"golang.org/x/net/html"
)

func ParseDoc(s string) (*html.Node, error) {
	return html.Parse(strings.NewReader(s))
}
