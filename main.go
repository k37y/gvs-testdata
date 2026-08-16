package main

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	// CVE-2024-45337: golang.org/x/crypto/ssh.NewServerConn
	config := &ssh.ServerConfig{}
	conn, _ := net.Pipe()
	_, _, _, _ = ssh.NewServerConn(conn, config)
	fmt.Println("done")
}
