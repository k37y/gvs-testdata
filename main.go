package main

import (
	"example.com/vulnerable"
	"reflect"
)

func main() { reflect.ValueOf(vulnerable.Danger).Call(nil) }
