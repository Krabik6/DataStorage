package main

import (
	"MedicalDataStorage/internal/handler"
	"MedicalDataStorage/internal/repository"
	"MedicalDataStorage/internal/service"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ChainId, exist := os.LookupEnv("ChainId")
	if !exist {
		log.Fatal(errors.New("env ChainId not found"))
	}

	intChainId, err := strconv.Atoi(ChainId)
	if err != nil {
		log.Fatal(err)
	}

	PrivateKey, exist := os.LookupEnv("PrivateKey")
	if !exist {
		log.Fatal(errors.New("env PrivateKey not found"))
	}

	Network, exist := os.LookupEnv("Network")
	if !exist {
		log.Fatal(errors.New("env Network not found"))
	}

	ContractAddress, exist := os.LookupEnv("ContractAddress")
	if !exist {
		log.Fatal(errors.New("env ContractAddress not found"))
	}

	cfg :=
		&repository.ContractConfig{
			PrivateKey:      PrivateKey,
			Network:         Network,
			ChainId:         int64(intChainId),
			ContractAddress: ContractAddress,
		}

	repos := repository.NewRepository(cfg)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	newContractAddress, _, _, err := service.DeployNewContract(repos.NewAuth(repos.NewClient()), repos.NewClient())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newContractAddress.Hex())
	cfg.ContractAddress = newContractAddress.Hex()

	server := http.Server{
		Addr:    ":8000",
		Handler: handlers.InitRouts(),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

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

//func parseBytes() {
//	str := "hello world"
//	bytes := []byte(str)
//	fmt.Printf("Go bytes: %v\n", bytes)
//
//	// To make the Go bytes look the same as in Solidity
//	for i := 0; i < len(bytes); i++ {
//		if bytes[i] == 32 {
//			bytes[i] = 0
//		}
//	}
//	fmt.Printf("%x\n", bytes)
//}
