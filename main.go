package main

import (
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

func main() {
	_, _ = net.LookupCNAME("example.org")
	var parser dnsmessage.Parser
	parser.Answer()
}
