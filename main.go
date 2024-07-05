package main

import (
	"fmt"
	"log"

	"github.com/iam-vl/chain-go/block"
	"github.com/iam-vl/chain-go/wallet"
)



func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKey())
	fmt.Println(w.PublicKey())
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())

}

func mainCh1() {
	myBlockchainAddr := "my_blockchain_address"
	blockchain := block.NewBlockchain(myBlockchainAddr)

	// blockchain := block.NewBlockchain(myBlockchainAddr)
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
	fmt.Printf("MY %1.f\n", blockchain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %1.f\n", blockchain.CalculateTotalAmount("C"))
	fmt.Printf("D %1.f\n", blockchain.CalculateTotalAmount("D"))
}
