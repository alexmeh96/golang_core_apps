package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Book struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("welcome"))
	})
	r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Max-Age", "300")
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Expose-Headers", "Date,Qqq")

		w.Write([]byte("hello alex"))
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		println(book.Name)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(book)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})

	println("server started!")
	err := http.ListenAndServeTLS(":8081", "./config/server.crt", "./config/server.key", r)
	if err != nil {
		panic(err)
	}
}
