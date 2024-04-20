package main

import (
	"fmt"
	"net/http"
)
	
func HelloWorld(w http.responseWriter, r *http.request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
