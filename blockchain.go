package main

import (
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
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
	b := &Block{} // new line
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash()) // change
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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
