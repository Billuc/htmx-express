package src

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func ServeRoutes(w http.ResponseWriter, url_path string) error {
	route_path := fmt.Sprintf("./routes/%s.html", url_path)
	route_file, err := os.Open(route_path)

	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(route_file)

	if err != nil {
		return err
	}

	htmxScript := doc.Find("head #htmx-src").First()

	if htmxScript.Length() == 0 {
		doc.Find("head").AppendHtml(`<script id="htmx-src" src="https://unpkg.com/htmx.org@1.9.10"></script>`)
	}

	goquery.Render(w, doc.Find("html"))
	fmt.Printf("Serving %s from routes\n", url_path)
	return nil
}
