package broken
import "example.com/vulnerable"
func Run() { vulnerable.Danger(); undefinedFunction() }
