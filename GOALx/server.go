package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	url_path := html.EscapeString(r.URL.Path)

	if url_path[0] == '/' {
		url_path = url_path[1:]
	}

	route_path := fmt.Sprintf("./routes/%s", url_path)
	route_data, err := os.ReadFile(route_path)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		http.NotFound(w, r)
	}

	fmt.Fprint(w, string(route_data))
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8091", nil)
}
