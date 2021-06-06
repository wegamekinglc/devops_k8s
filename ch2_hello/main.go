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

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
	log.Fatal(http.ListenAndServe(":8888", nil))
}
