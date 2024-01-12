package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
)

func test() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 20; i++ {

		// first step is to create a slice of bytes with the desired length
		buf := make([]byte, 32)
		// then we can call rand.Read.
		_, err = rand.Read(buf)
		if err != nil {
			fmt.Printf("error while generating random string: %s", err)
		}
		// print the bytes (numbers from 0 to 255) with %v format verb (raw value)
		// fmt.Printf("random bytes: %v", buf)
		// print the bytes encoded in hexadecimal with %x format verb
		fmt.Printf("random hex: %x", buf)
		p := hex.EncodeToString(buf)

		fmt.Println("we have a connection")

		privateKey, err := crypto.HexToECDSA(p)

		if err != nil {
			fmt.Println(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		fmt.Println("fromAddress: ", fromAddress)

		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)

		fmt.Println("balance: ", balance)
		if balance.Int64() > 0 {
			writeFile("private key: " + p + " balance: " + balance.String())
		}
	}
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

// write to the file
func writeFile(s string) {
	// Open the file for appending
	file, err := os.OpenFile("myfile.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The string was appended to the file successfully.")
}
func main() {
	// 3
	s := gocron.NewScheduler(time.UTC)
	s.Every(10).Seconds().Do(func() {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
		test()
	})
	// 5
	s.StartBlocking()
}
