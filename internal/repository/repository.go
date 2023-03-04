package repository

import (
	encryption "MedicalDataStorage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractInteractionData interface {
	NewAuth(*ethclient.Client) *bind.TransactOpts
	NewClient() *ethclient.Client
	ContractInstance(*ethclient.Client) *encryption.MedicalDataStorage
}

//type ContractInteractions interface {
//	GetUserInfo(_userAddress string, _password string) (string, error)
//	FillUserInfo(data []byte) (*types.Transaction, error)
//}

type Repository struct {
	ContractInteractionData
}

func NewRepository(cfg *ContractConfig) *Repository {
	return &Repository{
		ContractInteractionData: NewBlockchain(cfg),
	}
}
