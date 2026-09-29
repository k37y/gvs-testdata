package main
import "example.com/vulnerable"
func invoke[T ~int](n T) int { return int(n) + vulnerable.Danger() }
func main() { println(invoke(1)) }
