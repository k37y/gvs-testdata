//go:build windows

package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parseHTML() {
	doc, err := html.Parse(strings.NewReader("<html><body><p>hello</p></body></html>"))
	if err != nil {
		fmt.Println("html parse error:", err)
		return
	}
	fmt.Println("parsed:", doc.Type)
}
