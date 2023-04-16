package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprintf(w, "Hello, Golang Workshop @ HAW!\nURL = %s\n", req.URL)
}

func main() {
	fmt.Println("please connect to localhost:7777/hello")
	// adding our function to an endpoint
	http.HandleFunc("/hello", HelloServer)
	// starting the server
	log.Fatal(http.ListenAndServe(":7777", nil))
}
