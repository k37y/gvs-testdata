package main

import (
	"fmt"
	"strings"

	"github.com/k37y/gvs-testdata/wrapper"
)

func main() {
	doc, err := wrapper.ParseHTML(strings.NewReader("<html><body><p>hello</p></body></html>"))
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}
	fmt.Println("parsed:", doc.Type)
}
