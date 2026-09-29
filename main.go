package main
import ("example.com/vulnerable"; _ "example.com/helper")
func main() { println(vulnerable.Danger()) }
