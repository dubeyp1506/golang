package main

import (
	"fmt"
	"net/http"
)

func main() {
	const serverAdd string = "localhost:3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is a response")
	})

	fmt.Println("Server started at ", serverAdd)

	err := http.ListenAndServe(serverAdd, nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
