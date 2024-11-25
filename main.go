package main

import (
	"WeChatDatRecovery/core"
	"fmt"
)

func main() {
	path := core.GetChoice()
	fmt.Println("-----------------------------------------------------")

	core.ScanFiles(path.Input, path.Output, path.GoalDate)
}
