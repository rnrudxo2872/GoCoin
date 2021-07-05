package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rnrudxo2872/GoCoin/blockchain"
	"github.com/rnrudxo2872/GoCoin/utils"
)

var port string

type url string

type addBlockBody struct {
	Message string
}

func (u url) MarshalText() ([]byte, error) {
	finUrl := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(finUrl), nil
}

type URLDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"-"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func document(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
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
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}

}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	handler.HandleFunc("/", document)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
