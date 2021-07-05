package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rnrudxo2872/GoCoin/blockchain"
	"github.com/rnrudxo2872/GoCoin/utils"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	finUrl := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(finUrl), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"-"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func document(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "12DQW2F41124",
		},
		{
			URL:         "/blocks/{id}",
			Method:      "GET",
			Description: "See A Block",
		},
	}
	fmt.Println(data)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().GetAllBlock())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}

}

type AddBlockBody struct {
	Message string
}

func main() {
	http.HandleFunc("/", document)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
