package main

import "example.com/vulnerable"

type runner interface{ run() int }
type impl struct{}

func (impl) run() int     { return vulnerable.Danger() }
func invoke(r runner) int { return r.run() }
func main()               { println(invoke(impl{})) }
