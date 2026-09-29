//go:build gvs_integration_excluded

package main

import "example.com/vulnerable"

func excluded() int { return vulnerable.Danger() }
