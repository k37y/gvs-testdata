package main

import "example.com/vulnerable"

func invoke(f func() int) int { return f() }
func main()                   { println(invoke(vulnerable.Danger)) }
