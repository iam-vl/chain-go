package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("Hello world\n"))
	sum2 := sha256.Sum256([]byte("Hello world2\n"))
	fmt.Printf("Value: %v\n", sum)
	fmt.Println("===============")
	fmt.Printf("Hex value 1: %x\n", sum)
	fmt.Printf("Hex value 2: %x\n", sum2)

}
