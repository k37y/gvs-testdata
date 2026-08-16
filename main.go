package main

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/html"

	"github.com/k37y/gvs-testdata/helper"
)

func main() {
	// Uses helper which requires x/net v0.33.0 (patched)
	// But root go.mod still says x/net v0.23.0 (untidy)
	doc, err := helper.ParseDoc("<html><body><p>hello</p></body></html>")
	if err != nil {
		fmt.Println("html parse error:", err)
		return
	}
	fmt.Println("parsed:", doc.Type)

	// Also use html.Parse directly
	_, _ = html.Parse(nil)

	// CVE-2024-45337: golang.org/x/crypto/ssh.NewServerConn
	config := &ssh.ServerConfig{}
	conn, _ := net.Pipe()
	_, _, _, _ = ssh.NewServerConn(conn, config)
}
