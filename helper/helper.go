package helper
import ("reflect"; "example.com/vulnerable")
func Run() { reflect.ValueOf(vulnerable.Danger).Call(nil) }
