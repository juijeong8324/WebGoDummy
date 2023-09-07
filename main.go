package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
<<<<<<< Updated upstream
	fmt.Fprintf(w, "Hello, I'm uiJeong! -v6.0")
=======
	fmt.Fprintf(w, "Hello, UiJeong!-v6.0")
>>>>>>> Stashed changes
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
