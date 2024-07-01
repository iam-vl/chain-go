package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
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
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{} // new line
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
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

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(
			t.senderBlockchainAddress,
			t.recipientBlockchainAddress,
			t.value,
		))
	}
	return transactions
}

func (bc *Blockchain) ValidateProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidateProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}
