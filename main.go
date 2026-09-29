package main

import "example.com/vulnerable"

func unused() int { return vulnerable.Danger() }
func main()       { println(vulnerable.Safe()) }
