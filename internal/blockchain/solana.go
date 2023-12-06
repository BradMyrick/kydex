package blockchain

import (
	"github.com/gagliardetto/solana-go"
)

type SolanaClient struct {
	client *solana.Client
}

func SolanaConnect() (*SolanaClient, error) {
	// TODO: figure out solana client
	client := solana.NewClient("TODO") // TODO: use config

	return &SolanaClient{client: client}, nil
}

func (c *SolanaClient) SwapTokens(sourceChain, destinationChain string, amount float64) error {
	// TODO:Implement the logic for swapping tokens between sourceChain and destinationChain
	return nil
}
