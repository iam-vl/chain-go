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
	fmt.Println(m)         
	fmt.Println(string(m))   
	return sha256.Sum256([]byte(m))
}
func main() {
	b := &Block{nonce: 1}
	b.Print()
	fmt.Printf("Block hash: %x\n", b.Hash())
}
```
Output:  
```
[123 125]
{} // because of private field
Block hash: 44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a
``` 
Let's marshal it: 
```go
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    
		Nonce        int      
		PreviousHash string   
		Transactions []string 
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}
```
Result: 
```
[123 34 84 105 109 101 115 116 97 109 112 34 58 48 44 34 78 111 110 99 101 34 58 49 44 34 80 114 101 118 105 111 117 115 72 97 115 104 34 58 34 34 44 34 84 114 97 110 115 97 99 116 105 111 110 115 34 58 110 117 108 108 125]
{"Timestamp":0,"Nonce":1,"PreviousHash":"","Transactions":null}
Block hash: 394397db5179529bdc82ddee2b78021677a0e9ddca5e1da9f2d035491bc48754
```
Update json: 
```go
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash string   `json:"previous_hash"`
		Transactions []string `json:"transactions"`
	}
    // ... 
}
```
Result: 
```
{"timestamp":0,"nonce":1,"previous_hash":"","transactions":null}
```

Update `previousHash` for [32]byte:
```go
type Block struct {}
PreviousHash [32]byte `json:"previous_hash"`
func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {}
```
Update NewBlock:
```go
func NewBlockchain() *Blockchain {
	b := &Block{} // new line
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash()) // change
	return bc
}
```
Getting last block:  
```go
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}
```
Using block hashing in main: 
```go
blockchain := NewBlockchain()
prevHash := blockchain.LastBlock().Hash()
blockchain.CreateBlock(5, prevHash)
prevHash = blockchain.LastBlock().Hash()
blockchain.CreateBlock(2, prevHash)
blockchain.Print()
```

## Adding a transaction 

```go
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	fmt.Println("Adding a transaction...")
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}
```

## PoW, consensus, and nonce

```go
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
```
Main:  
```go
blockchain := NewBlockchain()

blockchain.AddTransaction("A", "B", 1.0)

prevHash := blockchain.LastBlock().Hash()

// blockchain.CreateBlock(5, prevHash)
nonce := blockchain.ProofOfWork()
blockchain.CreateBlock(nonce, prevHash)

blockchain.AddTransaction("C", "D", 2.0)
blockchain.AddTransaction("X", "Y", 3.0)

prevHash = blockchain.LastBlock().Hash()

// blockchain.CreateBlock(2, prevHash)
nonce := blockchain.ProofOfWork()
blockchain.CreateBlock(nonce, prevHash)
```

## All about mining
## Calculating the transaction total 
