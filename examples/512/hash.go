package main

import (
	"fmt"

	"github.com/dnonakolesax/streebog"
)

func main() {
	message := []byte("any-message")
	hash := streebog.New(64)
	hash.Write(message)
	fmt.Printf("Hash: %x\n", hash.Sum(nil))
}
