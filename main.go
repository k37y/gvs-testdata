package main

import (
	"example.com/vulnerable"
	"unsafe"
)

func main() { println(unsafe.Sizeof(vulnerable.Danger())); println(vulnerable.Danger()) }
