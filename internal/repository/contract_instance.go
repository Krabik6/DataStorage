package repository

import (
	encryption "MedicalDataStorage"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type Blockchain struct {
	cfg *ContractConfig
}

func NewBlockchain(cfg *ContractConfig) *Blockchain {
	return &Blockchain{cfg: cfg}
}

func (b *Blockchain) NewClient() *ethclient.Client {
	client, err := ethclient.Dial(b.cfg.Network)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (b *Blockchain) NewAuth(client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(b.cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(b.cfg.ChainId)))
	if err != nil {
		fmt.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth
}

func (b *Blockchain) ContractInstance(_client *ethclient.Client) *encryption.MedicalDataStorage {
	address := common.HexToAddress(b.cfg.ContractAddress)
	instance, err := encryption.NewMedicalDataStorage(address, _client)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}
