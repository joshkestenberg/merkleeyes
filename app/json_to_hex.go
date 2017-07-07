package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Print("Enter JSON string: ")
	var input string
	fmt.Scanln(&input)

	bytes := []byte(input)

	hex := hex.EncodeToString(bytes)

	fmt.Println("Your transaction as encoded as a byte string will be: 0x" + hex)
}
