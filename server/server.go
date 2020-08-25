package main

import "fmt"
import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OLA")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8888", nil)
}