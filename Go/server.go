package main

import (
	"goalx/src"
	"net/http"
)

func main() {
	http.HandleFunc("/", src.HandleRoot)
	http.ListenAndServe(":8091", nil)
}
