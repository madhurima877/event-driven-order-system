package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"UP"}`)
}
func main() {
	fmt.Println("Order Service starting")
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
