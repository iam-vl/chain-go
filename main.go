package main

import (
	"log"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddr := "my_blockchain_address"
	blockchain := NewBlockchain(myBlockchainAddr)
	// blockchain := NewBlockchain()

	blockchain.AddTransaction("A", "B", 1.0)

	// prevHash := blockchain.LastBlock().Hash()
	// nonce := blockchain.ProofOfWork()
	// blockchain.CreateBlock(nonce, prevHash)
	blockchain.Mining()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)

	// prevHash = blockchain.LastBlock().Hash()
	// nonce = blockchain.ProofOfWork()
	// blockchain.CreateBlock(nonce, prevHash)
	blockchain.Mining()
	blockchain.Print()
}
