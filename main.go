package main

import (
	"fmt"
	"log"
	"net/http"
)

func ResponseOk() string {
	return "OK"
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ResponseOk())
}

func main() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
