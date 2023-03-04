package service

import (
	"MedicalDataStorage/internal/repository"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type WriteService struct {
	repo *repository.Repository
}

func NewWriteService(repo *repository.Repository) *WriteService {
	return &WriteService{repo: repo}
}

func prepareUserInfo(info string, password string) []byte {
	data := encrypt(convertToBytes(info), convertToBytes(password))
	return data

}

func (w *WriteService) FillUserInfo(info string, password string, _userAddress string) (*types.Transaction, error) {

	data := prepareUserInfo(info, password)
	client := w.repo.NewClient()
	auth := w.repo.NewAuth(client)
	instance := w.repo.ContractInstance(client)
	fmt.Println(info, password, _userAddress)
	tx, err := instance.FillUSerInfo(auth, data, common.HexToAddress(_userAddress))
	if err != nil {
		return tx, err
	}
	return tx, err
}

func convertToBytes(text string) []byte {
	textBytes := []byte(text)
	return textBytes
}

func encrypt(data []byte, password []byte) []byte {
	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		encrypted[i] = data[i] ^ password[i%len(password)]
	}
	return encrypted
}
