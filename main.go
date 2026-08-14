package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("status:", resp.StatusCode)
}
