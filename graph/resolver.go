package graph

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EthClient *ethclient.Client
}
