# Wallet 

Plan:

* Private key, public key, and ECDSA
* How to create wallets 
* How to create blockchain addresses
* Creating blockchain addresses
* Signatures for transactions
* Transaction verification


## Private key, public key, and ECDSA 
Encrypt w/ public key
Decrypt with private key


## How to create wallets  

```go
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	return w
}
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {	return w.privateKey }
func (w *Wallet) PrivateKeyStr() string {	return fmt.Sprintf("%x", w.privateKey.D.Bytes()) }
func (w *Wallet) PublicKey() *ecdsa.PublicKey {	return w.publicKey }
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes, w.publicKey.Y.Bytes())
}
```
Main:
```go
func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKey())
	fmt.Println(w.PublicKey())
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
```

## Creating blockchain addresses 

Add dependencies:
```
go get golang.org/x/crypto/ripemd160
go get github.com/btcsuite/btcutil/base58
```

```go
func NewWallet() *Wallet {
	// 1. Create ECDSA privateK (32B) publicK (64B) - DONE before
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	// 2. Perform SHA-256 hashing on the publicK
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// 3. Perform RIPEMD-160 on the SHA-256 result (20B)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// 4. Add version byte in front of RIPE-160 hash (0x00 for main network)
	// (there are two networks - main and test)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// 5. Perform SHA-256 on the extended RIPEMD-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	// 6. Perform SHA-256 on the previous SHA-256 result
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	// 7. Take the first 4 bytes of the second SHA-256 for checksum
	chsum := digest6[:4]
	// 8. Add the 4 checksum bytes from 7 at the ned of extended RIPEMD-160 hash from 4 (25 bytes)
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	// 9 convert the res from a byte string to base58
	address := base58.Encode(dc8)
	w.blockchainAddress = address
	return w
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}
```
Main:  
```go
func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
}
```
Output:  
```
bf3ded9ea28a298a9f114fa28df0e29012d4434d0b899a9800a4b94c3c7399a1
10952801e61db3fe1ef6519c4cddb8e873091767edaed9d4c3a58ad829b9857e24d0397
16dWLUTUa6YG2gaV6houVK5Cvpdf3dUkJi
```


## Signatures for transactions

## Transaction verification 
