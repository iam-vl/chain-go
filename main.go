package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// b := &Block{nonce: 1}
	// // b.Print()//
	// fmt.Printf("Block hash: %x\n", b.Hash())

	blockchain := NewBlockchain()

	prevHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, prevHash)

	prevHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, prevHash)
	blockchain.Print()

	// b := &Block{nonce: 1}
	// fmt.Println("%x\n", b.Hash())
	// fmt.Println("%x\n", b.Hash())
}
