package main
import "example.com/vulnerable"
func main() { defer vulnerable.Danger() }
