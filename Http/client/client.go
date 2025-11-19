package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	// res, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	res, err := client.Get("https://swapi.dev/api/people/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
