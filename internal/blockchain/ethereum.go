package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	client *ethclient.Client
}

func EthConnect() (*EthereumClient, error) {
	// Connect to quicknode
	client, err := ethclient.Dial("https://quicknode.com/eth/mainnet") // TODO: use config
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	return &EthereumClient{client: client}, nil
}

func (c *EthereumClient) SwapTokens(sourceChain, destinationChain string, amount float64) error {
	// TODO:Implement the logic for swapping tokens between sourceChain and destinationChain
	return nil
}