package ss58_control_tools

import (
	"encoding/hex"
	"log"
)

var (
	ERC20_HEX_TRANSFER = "a9059cbb"
)

type ERC20 struct {
	Transfer []byte
}

func InitERC20Clinet() (ERC20, error) {
	transfer_bytes, err := hex.DecodeString(ERC20_HEX_TRANSFER)
	if err != nil {
		log.Fatalln(err)
		return ERC20{}, err
	}

	var erc20 = ERC20{
		Transfer: transfer_bytes,
	}
	return erc20, nil
}
