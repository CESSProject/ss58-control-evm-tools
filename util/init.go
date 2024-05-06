package ss58_control_tools

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

var config_path = "./config.json"

type BaseConfig struct {
	Rpc          []string `json:"rpc"`
	Mnemonic     string   `json:"mnemonic"`
	Source       string
	Contract     string `json:"contract"`
	GasLimit     uint64 `json:"gasLimit"`
	MaxFeePerGas uint64 `json:"maxFeePerGas"`
}

func LoadConfig(path string) *BaseConfig {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	base_config := &BaseConfig{}
	err = json.Unmarshal(buf, base_config)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	keyring, _ := signature.KeyringPairFromSecret(base_config.Mnemonic, 0)
	h160 := types.NewH160(keyring.PublicKey)
	base_config.Source = string(h160.Hex())

	contract, err := hex.DecodeString(base_config.Contract)
	if err != nil {
		log.Fatalln(err)
	}
	base_config.Contract = string(contract)

	return base_config
}

func Init() (*chain.ChainClient, *BaseConfig, error) {
	base_config := LoadConfig(config_path)

	fmt.Printf("BaseConfig Rpc: %v \n", base_config.Rpc)
	fmt.Printf("BaseConfig Memric: %v \n", base_config.Mnemonic)
	fmt.Printf("BaseConfig Source: %v \n", base_config.Source)
	fmt.Printf("BaseConfig Contract: %v \n", base_config.Contract)
	fmt.Printf("BaseConfig GasLimit: %v \n", base_config.GasLimit)
	fmt.Printf("BaseConfig MaxFeePerGas: %v \n", base_config.MaxFeePerGas)

	chain, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(base_config.Rpc),
		cess.Mnemonic(base_config.Mnemonic),
		cess.TransactionTimeout(time.Second*10),
	)

	return chain, base_config, err
}
