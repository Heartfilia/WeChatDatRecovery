package cmd

import (
	"fmt"
	"github.com/Heartfilia/litetools/litestr"
	"github.com/Heartfilia/litetools/utils/litedir"
	"log"
	"os"
	"path"
)

func ScanFiles(InputFolder string, OutputFolder string) {
	if !litedir.FileExists(InputFolder) {
		log.Printf("%s 输入路径不存在: %s", litestr.E(), InputFolder)
		return
	}
	datFolders := path.Join(InputFolder, "MsgAttach")
	readAndQuery(datFolders, OutputFolder)
}

func readAndQuery(pathName, outFolder string) {
	if litedir.IsDir(pathName) {
		thisDir, err := os.ReadDir(pathName)
		if err != nil {
			return
		}
		for _, entry := range thisDir {
			if entry.IsDir() {
				readAndQuery(path.Join(pathName, entry.Name()), outFolder)
			} else {
				fmt.Printf("Found file: %s\n", entry.Name())
			}
		}
	} else {
		fmt.Println(pathName)
	}
}
