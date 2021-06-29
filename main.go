package main

import (
	"fmt"

	"github.com/rnrudxo2872/GoCoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Seconde Block")
	chain.AddBlock("세번째 블록")

	for _, block := range chain.GetAllBlock() {
		fmt.Println("Data : " + block.Data)
		fmt.Println("Hash : " + block.Hash)
		fmt.Println("PrevHash : " + block.PrevHash)
	}
}
