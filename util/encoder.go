package ss58_control_tools

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/log"
)

// use setString convert to big.Int
func Uint256Encode(u big.Int) []byte {
	hex_u := u.Text(16)
	hex_u_string, err := hex.DecodeString(hex_u)
	if err != nil {
		log.Error("Uint256Encode Failed: %v \n", err)
		panic(err)
	}

	fill := make([]byte, 32-len(hex_u_string))

	fmt.Printf("fill is %v \n", fill)

	output := append(fill, hex_u_string...)
	return output
}

func AddressEncode(addr string) []byte {
	addr_hex, err := hex.DecodeString(addr)
	if err != nil {
		log.Error("AddressEncode Failed: %v \n", err)
		panic(err)
	}

	fill := make([]byte, 32-len(addr_hex))

	output := append(fill, addr_hex...)

	return output
}
