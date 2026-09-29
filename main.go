package main
import "example.com/vulnerable"
func init() { println(vulnerable.Danger()) }
func main() {}
