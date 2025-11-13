package main

import (
	"context"
	"fmt"
	"log"
	// after including write
	// "crypto/ecdsa"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// after including write
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"staking-contract-function-test/mycontract"
)

func main() {
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	if err != nil {
		log.Fatal("Failed to connect RPC:", err)
	}

	contractAddress := common.HexToAddress("0xBE88Adb48eC38188336D3Fb1f81Ba34f623DC728")
	contract, err := mycontract.NewMycontract(contractAddress, client)
	if err != nil {
		log.Fatal("Failed to load contract:", err)
	}

	// Call a view function
	minStakeDuration, err := contract.MinStakeDuration(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatal("Call failed:", err)
	}

	fmt.Println("Minimum Stake Duration:", minStakeDuration)
	maxStakeDuration, err := contract.MaxStakeDuration(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatal("Call failed:", err)
	}

	fmt.Println("Maximum Stake Duration:", maxStakeDuration)

	owner, err := contract.Owner(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatal("Call failed:", err)
	}

	fmt.Println("Owner Address:", owner)

	_ = godotenv.Load(".env") //usually at start of main

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("PRIVATE_KEY not found in environment")
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// -----------------------------
	// Execute transaction: Unstake
	// -----------------------------
	chainID := big.NewInt(97) // BSC Testnet
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}
	userAddress := auth.From
	count, err := contract.GetUserStakesCount(&bind.CallOpts{
		Context: context.Background(),
	}, userAddress)
	if err != nil {
		log.Fatal("Failed to fetch stake count:", err)
	}
	fmt.Println("Stake count for", userAddress.Hex(), "=", count)

	if count.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("No stakes found — nothing to unstake.")
	}

	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	auth.Context = context.Background()
	auth.GasLimit = 300000 // optional override
	if count.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("❌ No active stakes for this address.")
	}

	// Always use 0 if only one stake
	index := big.NewInt(0) // <-- change if needed
	tx, err := contract.Unstake(auth, index)
	if err != nil {
		log.Fatal("Unstake transaction failed:", err)
	}

	fmt.Println("🔗 Transaction sent!")
	fmt.Println("TX Hash:", tx.Hash().Hex())

	// -----------------------------
	// Wait for confirmation (like await tx.wait())
	// -----------------------------
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Error waiting for transaction:", err)
	}

	if receipt.Status == 1 {
		fmt.Println("✅ Unstake transaction succeeded!")
	} else {
		fmt.Println("❌ Unstake transaction failed!", receipt)
	}
}
