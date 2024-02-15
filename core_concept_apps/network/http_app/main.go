package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /path/{id}", func(w http.ResponseWriter, r *http.Request) {
		params, _ := url.ParseQuery(r.URL.RawQuery)
		fmt.Println(params)

		fmt.Fprint(w, "got path\n")
	})

	mux.HandleFunc("/task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	http.ListenAndServe("localhost:8090", mux)
}
