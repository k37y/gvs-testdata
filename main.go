package main
import ("example.com/vulnerable"; _ "example.com/missing")
func main() { vulnerable.Danger() }
