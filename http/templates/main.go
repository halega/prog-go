package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", editHandler)
	log.Println("App started at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "edit.html")
		return
	}

}
