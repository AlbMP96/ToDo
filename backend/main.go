package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola desde Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
