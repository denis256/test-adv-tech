package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting HTTP server")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", root)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root Page\n")

}
