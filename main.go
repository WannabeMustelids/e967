package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type Submission struct {
	Id      int
	Created time.Time
	Tags    []string
	Unsafe  bool
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html.tmpl"))

	data := []Submission{
		{
			Id:      1,
			Tags:    []string{"test", "dummy"},
			Created: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
