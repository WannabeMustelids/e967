package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var r = httprouter.New()

func init() {
	r.GET("/", Index)
	r.GET("/hello/:name", Hello)

	r.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/html")
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, `Page not found. You should <a href="//%s">go back to the homepage.</a>`, req.Host)
	})

	http.Handle("/", r)
}

func Index(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(rw, "Dook!\n")
}

func Hello(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(rw, "Dook, %s!\n", ps.ByName("name"))
}

func main() {
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
