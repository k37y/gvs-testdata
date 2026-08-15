package main

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(strings.NewReader("<html><body><p>hello</p></body></html>"))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("parsed:", doc.Type)

	config := &ssh.ServerConfig{}
	conn, _ := net.Pipe()
	_, _, _, _ = ssh.NewServerConn(conn, config)
}
