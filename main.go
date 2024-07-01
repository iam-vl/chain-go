package main

import (
	"fmt"
	"log"
)

const MINING_DIFFICULTY = 3

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// b := &Block{nonce: 1}
	// // b.Print()//
	// fmt.Printf("Block hash: %x\n", b.Hash())

	blockchain := NewBlockchain()

	blockchain.AddTransaction("A", "B", 1.0)

	prevHash := blockchain.LastBlock().Hash()
	nonce := blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, prevHash)

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)

	prevHash = blockchain.LastBlock().Hash()
	nonce = blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, prevHash)
	blockchain.Print()
	fmt.Println()

	// b := &Block{nonce: 1}
	// fmt.Println("%x\n", b.Hash())
	// fmt.Println("%x\n", b.Hash())
}
