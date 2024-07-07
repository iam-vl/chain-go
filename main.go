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
	walletM := wallet.NewWallet() // miner
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()
	// fmt.Println(w.PrivateKey())
	// fmt.Println(w.PublicKey())
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
	// fmt.Println(w.BlockchainAddress())

	// Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)
	// fmt.Printf("signature %s\n", t.GenerateSignature())

	// Blockchain 
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added:", isAdded)
	blockchain.Mining()
	blockchain.Print()


	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))

}

// func mainCh1() {
// 	myBlockchainAddr := "my_blockchain_address"
// 	blockchain := block.NewBlockchain(myBlockchainAddr)
// 	blockchain.AddTransaction("A", "B", 1.0)
// 	blockchain.Mining()

// 	blockchain.AddTransaction("C", "D", 2.0)
// 	blockchain.AddTransaction("X", "Y", 3.0)

// 	blockchain.Mining()
// 	blockchain.Print()
// 	fmt.Printf("MY %1.f\n", blockchain.CalculateTotalAmount("my_blockchain_address"))
// 	fmt.Printf("C %1.f\n", blockchain.CalculateTotalAmount("C"))
// 	fmt.Printf("D %1.f\n", blockchain.CalculateTotalAmount("D"))
// }
