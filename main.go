package main

import (
	"WeChatDatRecovery/cmd"
)

func main() {
	//path := cmd.GetChoice()
	//fmt.Println("-----------------------------------------------------")
	var path = cmd.PathConf{
		Input:  "D:\\Cache\\Weixin\\WeChat Files\\xxxxx\\FileStorage",
		Output: "",
	}
	cmd.ScanFiles(path.Input, path.Output)
}
