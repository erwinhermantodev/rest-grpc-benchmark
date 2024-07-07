// rest_server.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Millisecond) // simulate some processing time
	fmt.Fprintf(w, "Hello, REST!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Starting REST server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
