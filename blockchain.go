package main

import (
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init_hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash string) *Block {
	b := NewBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
}

// func NewBlockchain() *Blockchain {
// 	b := &Block{}
// 	bc := new(Blockchain)
// 	bc.CreateBlock(0, "init hash")
// 	return bc
// }
// func (bc *Blockchain) CreateBlock(nonce int, prevHash string) *Block {
// 	b := NewBlock(nonce, prevHash)
// 	bc.chain = append(bc.chain, b)
// 	return b
// }
