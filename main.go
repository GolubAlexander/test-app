package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello")
	})

	address := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(address, nil))
}
