package src

import (
	"fmt"
	"net/http"
	"os"
)

func ServePublic(w http.ResponseWriter, url_path string) error {
	public_path := fmt.Sprintf("./public/%s", url_path)
	file_data, err := os.ReadFile(public_path)

	if err != nil {
		return err
	}

	fmt.Fprint(w, string(file_data))
	fmt.Printf("Serving %s from public\n", url_path)
	return nil
}
