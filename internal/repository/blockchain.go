package repository

type ContractConfig struct {
	PrivateKey      string
	Network         string
	ChainId         int64
	ContractAddress string
}

//func NewContractConfig(cfg ContractConfig) *ContractConfig {
//	return ContractConfig{PrivateKey: cfg.PrivateKey}
//}
