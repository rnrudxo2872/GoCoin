package main

import (
	"github.com/rnrudxo2872/GoCoin/explorer"
	"github.com/rnrudxo2872/GoCoin/rest"
)

const port string = ":4000"

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
