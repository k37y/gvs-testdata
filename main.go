package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(strings.NewReader("<html><body><p>hello</p></body></html>"))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("parsed:", doc.Type)
}
