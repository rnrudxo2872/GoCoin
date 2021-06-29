package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rnrudxo2872/GoCoin/blockchain"
)

var port string = ":4000"

type homeData struct {
	PageTitle string
	Block     []*blockchain.Block
}

func handleHome(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().GetAllBlock()}

	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", handleHome)

	fmt.Printf("Now Listening Server! http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
