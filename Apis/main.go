package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Main is running")
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handling orders")
	})

	port := ":3000"

	fmt.Println("Server is running on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
