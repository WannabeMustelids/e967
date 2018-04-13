package daemon

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/wannabemustelids/e967/model"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html.tmpl"))

	data := []model.File{
		{
			Id:       1,
			Tags:     []string{"test", "dummy"},
			Uploaded: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
	}
	tmpl.Execute(w, data)
}

func Run() {
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
