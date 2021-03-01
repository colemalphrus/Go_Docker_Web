package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))

	router := mux.NewRouter()

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/about/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
        name := vars["name"]
		fmt.Fprintf(w, "hi %s this is about us", name)
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		fmt.Fprint(w, "Hello Go!")
	})

	http.ListenAndServe(":8080", router)
}
