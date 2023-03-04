package service

import (
	"MedicalDataStorage/internal/repository"
	"github.com/ethereum/go-ethereum/core/types"
)

type Queries interface {
	GetUserInfo(_userAddress string, _password string) (string, error)
}

type Writes interface {
	FillUserInfo(info string, password string, _userAddress string) (*types.Transaction, error)
}

type Service struct {
	Queries
	Writes
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Queries: NewQueryService(repository),
		Writes:  NewWriteService(repository),
	}
}
