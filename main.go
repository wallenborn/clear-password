package main

import (
	"fmt"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: clearpassword <p12file> <password>")
		os.Exit(1)
	}

	inputFile := args[0]
	password := args[1]
	outputFile := args[2]

	// Read the PKCS12 file
	p12Data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Failed to read PKCS12 file: %v\n", err)
		return
	}

	// Decode the PKCS12 file
	privateKey, cert, caCerts, err := pkcs12.DecodeChain(p12Data, password)
	if err != nil {
		fmt.Printf("Failed to decode PKCS12 file: %v\n", err)
		return
	}

	// Re-encode without password
	newP12Data, err := pkcs12.Modern.Encode(privateKey, cert, caCerts, "")
	if err != nil {
		fmt.Printf("Failed to encode PKCS12 file: %v\n", err)
		return
	}

	// Save the new PKCS12 file
	err = os.WriteFile(outputFile, newP12Data, 0644)
	if err != nil {
		fmt.Printf("Failed to write output PKCS12 file: %v\n", err)
		return
	}

}
