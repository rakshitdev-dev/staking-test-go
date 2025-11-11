package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"staking-contract-function-test/mycontract"
)

func main() {
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	if err != nil {
		log.Fatal("Failed to connect RPC:", err)
	}

	contractAddress := common.HexToAddress("0x87a54a4cC6bfAc0D1Fd3286619508d1299753c7b")
	contract, err := mycontract.NewMycontract(contractAddress, client)
	if err != nil {
		log.Fatal("Failed to load contract:", err)
	}

	// Call a view function
	stakeDuration, err := contract.StakeDuration(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatal("Call failed:", err)
	}

	fmt.Println("Stake Duration:", stakeDuration)
}
