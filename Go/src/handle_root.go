package src

import (
	"fmt"
	"html"
	"net/http"
	"strings"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	url_path := html.EscapeString(r.URL.Path)

	if strings.HasSuffix(url_path, "/") {
		url_path = url_path + "index"
	}
	url_path = strings.TrimPrefix(url_path, "/")

	err := ServeRoutes(w, url_path)

	if err == nil {
		return
	}

	err = ServePublic(w, url_path)

	if err == nil {
		return
	}

	fmt.Printf("Error: %s\n", err)
	http.NotFound(w, r)
}
