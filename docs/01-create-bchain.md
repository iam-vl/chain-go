# Create blockchain

Plan: 

* Creating a block struct
* Creating a blockchain struct
* Creating the hash
* Adding a transaction
* PoW, consensus, and nonce
* Deriving a nonce 
* All about mining
* Calculating the transaction total 

## Getting started 

```go
func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	log.Println("test")
	fmt.Println("test2")
}
```
Output:  
```  
Blockchain: 2024/06/29 17:51:43 test
test2 
``` 

## Creating a block struct
To include prevHash, timestamp, nonce, trasactions
User: send Coin 
Pool: []transactions
Transactions: recipient, sender, value

```go
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block) // provides a pointer
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}
```
Plus, add b.Print()

## Creating a blockchain struct

```go
type Blockchain struct {
    transactionPool []string
    chain []*Block
}
func (bc *Blockchain) CreateBlock(nonce int, prevHash string) *Block {
    b := NewBlock(nonce, prevHash)
    bc.chain = append(bc.chain, b)
    return b
}

func NewBlockchain() *Blockchain {
    bc := new(Blockchain)
    bc.CreateBlock(0, "init_hash")
    return bc
}
func (bc *Blockchain) Print() {
    fmt.Println("printing blockchain")
    for i, block := range bc.chain {
        fmt.Printf("chain %d \n", i)
        block.Print()
    }
}
```
Mains:  
```go
blockchain := NewBlockchain()
blockchain.CreateBlock(5, "hash 1")
blockchain.CreateBlock(2, "hash 2")
blockchain.Print()
```

## Creating the hash

Packages : crypto/sha256
```go
sum := sha256.Sum256([]byte("hello world\n"))
fmt.Printf("%v\n", sum) // native
fmt.Printf("%x\n", sum) // hex
```

```go
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	// fmt.Printf("M Type: %T\n", m)           // M Type: []uint8
	// fmt.Printf("M string: %s\n", string(m)) // M string: {} - need to marshal
	return sha256.Sum256([]byte(m))
}
func main() {
	b := &Block{nonce: 1}
	fmt.Println("%x\n", b.Hash())
}
```
Output:  
```
M Type: []uint8
M string: {}
[68 19 111 163 85 179 103 138 17 70 173 22 247 232 100 158 148 251 79 194 31 231 126 131 16 192 96 246 28 170 255 138]
``` 
private fields, so: 
```

## Adding a transaction
## PoW, consensus, and nonce
## Deriving a nonce 
## All about mining
## Calculating the transaction total 
