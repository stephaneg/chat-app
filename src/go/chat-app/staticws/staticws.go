package main

import (
	"fmt"
	"log"
	"net/http"
)

func qrCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/qrcode" {
		http.Error(w, "404 not 222 found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./"))

	http.Handle("/", fileServer)
	http.HandleFunc("/qrcode", qrCodeHandler)

	fmt.Printf("Starting server at port 443\n")
	if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); err != nil {
		log.Fatal(err)
	}

}
