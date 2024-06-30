package main

import (
	"fmt"
	"strings"
)

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previousHash    %s\n", b.previousHash)
	fmt.Printf("transactions    %d\n", b.transactions)
}

func (bc *Blockchain) Print() {
	fmt.Println("printing blockchain")
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", m15("="), i, m15("="))
		block.Print()
	}
	fmt.Printf("%s\n", m15("*"))
}

func m15(s string) string {
	return strings.Repeat(s, 15)
}

// func (b *Block) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Timestamp    int64    `json:"timestamp"`
// 		Nonce        int      `json:"nonce"`
// 		PreviousHash string   `json:"previous_hash"`
// 		Transactions []string `json:"transactions"`
// 	}{
// 		Timestamp:    b.timestamp,
// 		Nonce:        b.nonce,
// 		PreviousHash: b.previousHash,
// 		Transactions: b.transactions,
// 	})
// }
// func (b *Block) Hash() [32]byte {
// 	m, _ := json.Marshal(b)
// 	return sha256.Sum256([]byte(m))
// }
