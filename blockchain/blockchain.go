package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	Hash     string
	prevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getPrevHash() string {
	totalBlockLen := len(b.blocks)

	if totalBlockLen == 0 {
		return ""
	}

	return b.blocks[totalBlockLen-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getPrevHash()}
	newBlock.calculateHash()
	return &newBlock
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}
