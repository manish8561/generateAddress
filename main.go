package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func main() {
	// first step is to create a slice of bytes with the desired length
	buf := make([]byte, 32)
	// then we can call rand.Read.
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}
	// print the bytes (numbers from 0 to 255) with %v format verb (raw value)
	log.Printf("random bytes: %v", buf)
	// print the bytes encoded in hexadecimal with %x format verb
	log.Printf("random hex: %x", buf)
	log.Println("hi: ", hex.EncodeToString(buf))
}