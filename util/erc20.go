package ss58_control_tools

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func (c *ERC20) TransferExtrinsic(sdk *chain.ChainClient, config *BaseConfig, target string, amount big.Int) (string, error) {
	input_amount := Uint256Encode(amount)
	fmt.Printf("\n slice is %v \n", target[2:])
	input_target := AddressEncode(target[2:])

	inputs := append(c.Transfer, input_target...)
	input := append(inputs, input_amount...)

	fmt.Printf("input is: %v \n", input)

	s_h160, err := hex.DecodeString(config.Source[2:])
	if err != nil {
		log.Fatalln(err)
	}
	source := types.NewH160(s_h160)

	value := types.NewU256(*big.NewInt(0))
	max_fee_pergas := types.NewU256(*big.NewInt(int64(config.MaxFeePerGas)))
	var accessList []pattern.AccessInfo

	return sdk.SendEvmCall(source, types.NewH160([]byte(config.Contract)), input, value, types.U64(config.GasLimit), max_fee_pergas, accessList)
}
