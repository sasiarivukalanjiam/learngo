package main

import (
        "fmt"
	"net/http"
	"log"
)

func Hello(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(w, "Hello world!")
}

func main() {
        http.HandleFunc("/", Hello)
	fmt.Println("Starting up on 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
