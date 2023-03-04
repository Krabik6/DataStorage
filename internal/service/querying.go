package service

import (
	"MedicalDataStorage/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

type QueryService struct {
	repo *repository.Repository
}

func NewQueryService(repo *repository.Repository) *QueryService {
	return &QueryService{repo: repo}
}

func (s *QueryService) Decrypt(encryptedData []byte) (string, error) {
	client := s.repo.NewClient()
	instance := s.repo.ContractInstance(client)
	response, err := instance.TestDecryption(nil, encryptedData, "world")
	if err != nil {
		log.Fatal(err)
	}
	return response, err
}

func (s *QueryService) Encrypt(_data string, _password string) ([]byte, error) {
	client := s.repo.NewClient()
	instance := s.repo.ContractInstance(client)
	response, err := instance.TestEncryption(nil, _data, _password)
	if err != nil {
		log.Fatal(err)
	}
	return response, err
}

func (s *QueryService) GetUserInfo(_userAddress string, _password string) (string, error) {
	client := s.repo.NewClient()
	instance := s.repo.ContractInstance(client)
	response, err := instance.GetUSerInfo(nil, common.HexToAddress(_userAddress), _password)
	if err != nil {
		log.Fatal(err)
	}
	return response, err
}
