package service

import (
	encryption "MedicalDataStorage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func DeployNewContract(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, *encryption.MedicalDataStorage, error) {
	address, tx, instance, err := encryption.DeployMedicalDataStorage(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	return address, tx, instance, err
}
