package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	r.Get("/app2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app2"))
	})
	r.Get("/app2/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app2 hello"))
	})

	println("app2 started!")
	err := http.ListenAndServe(":8086", r)
	if err != nil {
		panic(err)
	}
}
