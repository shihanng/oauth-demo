package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[len("/protected/"):])
}

func main() {
	http.HandleFunc("/protected/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
