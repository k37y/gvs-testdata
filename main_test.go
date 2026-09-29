package main

import (
	"example.com/vulnerable"
	"testing"
)

func TestDanger(t *testing.T) {
	if vulnerable.Danger() != 42 {
		t.Fatal("unexpected value")
	}
}
