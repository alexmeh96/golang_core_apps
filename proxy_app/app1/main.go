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
	r.Get("/app1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app1"))
	})
	r.Get("/app1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app1 hello"))
	})

	println("app1 started!")
	err := http.ListenAndServeTLS(":8085", "../config/server.crt", "../config/server.key", r)
	if err != nil {
		panic(err)
	}
}
