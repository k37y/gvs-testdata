package main
import "github.com/dgrijalva/jwt-go"
func main() { _, _ = jwt.Parse("token", func(*jwt.Token) (interface{}, error) { return []byte("secret"), nil }) }
