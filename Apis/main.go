package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Main is running")
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handling orders")
	})

	port := ":3000"

	certFile := "cert.pem"
	keyFile := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		TLSConfig: tlsConfig,
		Handler:   nil,
	}

	//Enable TLS

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port", port)

	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		fmt.Println(err)
	}
}
