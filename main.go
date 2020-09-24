package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func decorator(h func(r *http.Request) Response) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := h(r)
		result.Write(w)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", decorator(Logging(Index, "index"))).Methods("GET")
	router.HandleFunc("/todos", decorator(Logging(TodoIndex, "todo-index"))).Methods("GET")
	router.HandleFunc("/todos/{todoId}", decorator(IDShouldBeInt(TodoShow, "todo-show"))).Methods("GET")
	router.HandleFunc("/todos", decorator(Logging(TodoCreate, "todo-create"))).Methods("POST")
	router.HandleFunc("/todos/{todoId}", decorator(IDShouldBeInt(TodoDelete, "todo-delete"))).Methods("DELETE")

	http.Handle("/", router)

	log.Println("start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
