package main

import "net"

func main() {
	_, _ = net.LookupCNAME("example.org")
}
