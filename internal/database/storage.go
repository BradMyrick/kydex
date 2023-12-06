package database

import (
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func init() {
	// Connect to the IPFS node
	sh = shell.NewShell("localhost:5001") //TODO: use config
	if _, err := sh.ID(); err != nil {
		log.Fatalf("Failed to connect to the IPFS node: %v", err)
	}
}

func StoreSwap(swap *Swap) (string, error) {
	// Convert the Swap struct to JSON
	// Add the JSON to IPFS
	// Return the CID of the added content
	return "", nil
}

func RetrieveSwap(cid string) (*Swap, error) {
	// Retrieve the content from IPFS using the CID
	// Convert the content to a Swap struct
	// Return the Swap struct
	return nil, nil
}
