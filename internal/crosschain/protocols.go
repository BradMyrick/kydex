package crosschain

import (
	"errors"

	"github.com/BradMyrick/kydex/internal/blockchain"
)

type Protocol interface {
	SwapTokens(sourceChain, destinationChain string, amount float64) error
}


func GetProtocol(name string) (Protocol, error) {
	switch name {
	case "ethereum":
		client, err := blockchain.EthConnect()
		if err != nil {
			return nil, err
		}
		return client, nil
	case "solana":
		client, err := blockchain.SolanaConnect()
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, errors.New("unsupported protocol")
	}
}
