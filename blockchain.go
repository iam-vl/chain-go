package main

import (
	"fmt"
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
	b.transactions = transactions
	return b
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{} // new line
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash()) // change
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	// empting the pool
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	fmt.Println("Adding a transaction...")
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
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
