package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I'm uiJeong! -v6.0")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
