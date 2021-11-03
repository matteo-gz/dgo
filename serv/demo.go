package main

import (
	"fmt"
	"net/http"
)

func main() {
	if http.ListenAndServe("0.0.0.0:8081", nil) != nil {
		fmt.Println("8081 http listen failed")
	}
}
