package main

import (
	"fmt"
	"log"
	"net/http"
)

var port string = ":4000"

func handleHome(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello! Go Server!")

}

func main() {
	http.HandleFunc("/", handleHome)

	fmt.Printf("Now Listening Server! http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
