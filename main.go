package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	blockchain := NewBlockchain()
	blockchain.CreateBlock(5, "hash 1")
	blockchain.CreateBlock(2, "hash 2")
	blockchain.Print()

	// b := &Block{nonce: 1}
	// fmt.Println("%x\n", b.Hash())
	// fmt.Println("%x\n", b.Hash())
}
