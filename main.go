package main

import (
	"encoding/json"
	"fmt"
	"test-adv-tech/db"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting HTTP server")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", root)
	router.HandleFunc("/request", request)

	log.Fatal(http.ListenAndServe(":"+getenv("LISTEN_PORT", "10000"), router))
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root Page, access /request for application logic\n")
}

func request(w http.ResponseWriter, r *http.Request) {
	var data = db.ReadData()
	json.NewEncoder(w).Encode(data)
}
