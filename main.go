package main
import "example.com/vulnerable"
func main() { done := make(chan struct{}); go func() { println(vulnerable.Danger()); close(done) }(); <-done }
