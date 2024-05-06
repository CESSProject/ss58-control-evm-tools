package main

import (
	"log"
	ss58_control_tools "ss58-control-tools/util"
)

func main() {
	chain_sdk, base_config, err := ss58_control_tools.Init()
	if err != nil {
		log.Fatalf("SS58 Control Tools init failed:%v", err)
	}

	log.Println("111111111111111111")
	ss58_control_tools.InitCmd(chain_sdk, base_config)
	log.Println("222222222222222222")
	ss58_control_tools.Execute()
	log.Println("333333333333333333")
}
