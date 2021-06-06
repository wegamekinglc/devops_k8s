package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, 世界")
	log.Println("This is my hello world example")
}

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am fine")
	log.Println("This is the health check")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", handlerHealthz)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
