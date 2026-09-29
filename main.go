package main

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/html"
)

func main() {
	// CVE-2024-45338: golang.org/x/net/html.Parse
	doc, err := html.Parse(strings.NewReader("<html><body><p>hello</p></body></html>"))
	if err != nil {
		fmt.Println("html parse error:", err)
		return
	}
	fmt.Println("parsed:", doc.Type)

	// CVE-2024-45337: golang.org/x/crypto/ssh.NewServerConn
	config := &ssh.ServerConfig{}
	conn, _ := net.Pipe()
	_, _, _, _ = ssh.NewServerConn(conn, config)
}
