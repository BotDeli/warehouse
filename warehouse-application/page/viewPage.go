package page

import (
	"fmt"
	"net/http"
	"os"
	"warehouse-application/errors"
)

func loadPage(title string) []byte {
	fileName := fmt.Sprintf("/Users/kare/GolandProjects/warehouse-application/page/%s.html", title)
	body, err := os.ReadFile(fileName)
	if err != nil {
		return []byte("Error")
	}
	return body
}

func ShowPage(w http.ResponseWriter, r *http.Request) {
	title := "_" + r.URL.Path[1:]
	page := loadPage(title)
	_, err := w.Write(page)
	errors.CheckWarning(err)
}
