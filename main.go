package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func main() {
	// Read the ABI file into a memory buffer
	// abi, err := os.ReadFile("YourContract.abi")
	// if err != nil {
	// 	log.Fatalf("Failed to read ABI file: %v", err)
	// }

	// Read the private key from an environment variable
	privateKey, ok := os.LookupEnv("PRIVATE_KEY")
	if !ok || privateKey == "" {
		log.Fatalf("Private key not found in environment")
	}

	// Connect to foundry's Anvil
	sdk, err := thirdweb.NewThirdwebSDK("http://localhost:8545", &thirdweb.SDKOptions{
		SecretKey: privateKey,
	})
	if err != nil {
		panic(err)
	}

	chainID, err := sdk.GetChainID(context.Background())
	if err != nil {
		wErr := fmt.Errorf("failed to get chain ID: %w", err)
		panic(wErr)
	}

	fmt.Printf("Chain ID: %s\n", chainID)
}
