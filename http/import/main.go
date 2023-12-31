package main

import (
	"io"
	http1 "net/http"
	http2 "net/http"
)

func main() {
	http1.HandleFunc("/", func(w http1.ResponseWriter, r *http1.Request) {
		io.WriteString(w, "Hello http1")
	})
	go http1.ListenAndServe(":8081", nil)

	http2.HandleFunc("/", func(w http2.ResponseWriter, r *http2.Request) {
		io.WriteString(w, "Hello http2")
	})
	http2.ListenAndServe(":8080", nil)
}
