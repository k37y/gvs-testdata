package main
import "example.com/vulnerable"
func alpha() { println(vulnerable.Danger()) }
func beta() { println(vulnerable.Other()) }
func main() { alpha(); beta() }
