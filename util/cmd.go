package ss58_control_tools

import (
	"fmt"
	"math/big"
	"os"

	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cess-tools-erc20",
	Short: "CESS TOOLS is a tool for calling smart contracts",
	Long: `CESS TOOLS is a tool for calling smart contracts. 
			It can use the cess address to control its mapped 
			evm address to call smart contracts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cess tools...")
	},
}

func InitCmd(sdk *chain.ChainClient, config *BaseConfig) {
	var transferCMD = &cobra.Command{
		Use:   "transfer",
		Short: "CESS TOOLS is a tool for calling smart contracts",
		Long: `CESS TOOLS is a tool for calling smart contracts. 
				It can use the cess address to control its mapped 
				evm address to call smart contracts.`,
		Run: func(cmd *cobra.Command, args []string) {
			target := args[0]
			arg_amount := args[1]

			amount := new(big.Int)
			amount.SetString(arg_amount, 10)

			client, err := InitERC20Clinet()
			if err != nil {
				fmt.Println("Init Erc20 Clinent Failed")
				panic(err)
			}

			hash, err := client.TransferExtrinsic(sdk, config, target, *amount)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Transfer Result: %v \n", hash)
		},
		Args: cobra.MaximumNArgs(2),
	}

	rootCmd.AddCommand(transferCMD)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
