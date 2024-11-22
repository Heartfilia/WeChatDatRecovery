package main

import (
	"WeChatDatRecovery/cmd"
	"fmt"
)

func main() {
	path := cmd.GetChoice()
	fmt.Println("-----------------------------------------------------")

	cmd.ScanFiles(path.Input, path.Output, path.GoalDate)
}
