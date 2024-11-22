package cmd

import (
	"github.com/Heartfilia/litetools/litestr"
	"github.com/Heartfilia/litetools/utils/litedir"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

var REG *regexp.Regexp
var OK_COUNT = 0

func ScanFiles(InputFolder, OutputFolder, goalDate string) {
	if !litedir.FileExists(InputFolder) {
		log.Printf("%s 输入路径不存在: %s", litestr.E(), InputFolder)
		return
	}
	var datFolders = InputFolder
	if litedir.IsDir(InputFolder) {
		datFolders = path.Join(InputFolder, "MsgAttach")
	}
	compile, err := regexp.Compile("\\d{4}-\\d{2}")
	if err != nil {
		return
	}
	REG = compile

	if goalDate != "" {
		log.Println("3秒后开始恢复指定年月的数据:" + litestr.ColorString(goalDate, "red"))
		time.Sleep(time.Second * 3)
	}
	readAndQuery(datFolders, OutputFolder, goalDate, false)
}

func judgeDat(fileName string) bool {
	if strings.Index(fileName, ".dat") == -1 {
		return false
	}
	return true
}

func readAndQuery(pathName, outFolder, goalDate string, last bool) {
	if litedir.IsDir(pathName) {
		thisDir, err := os.ReadDir(pathName)
		if err != nil {
			log.Fatalln("读取文件路径报错.")
			return
		}
		for _, entry := range thisDir {
			if entry.IsDir() {
				thisFolder := entry.Name()
				if thisFolder == "Thumb" {
					continue
				}
				dateFolder := REG.FindString(thisFolder)
				if dateFolder != "" {
					if goalDate != "" && dateFolder != goalDate {
						// 如果传入了指定的年月数据 那么就只跑这个年月的数据
						continue
					}
					saveFolder := path.Join(outFolder, dateFolder)
					if !litedir.FileExists(saveFolder) {
						err := os.Mkdir(saveFolder, 0755)
						if err != nil {
							return
						}
					}
					// 这个地方的下一级就是需要的文件了
					readAndQuery(path.Join(pathName, thisFolder), saveFolder, goalDate, true)
				} else {
					readAndQuery(path.Join(pathName, thisFolder), outFolder, goalDate, false)
				}
			} else if last && judgeDat(entry.Name()) {
				err := ParseAndSave(path.Join(pathName, entry.Name()), outFolder)
				if err != nil {
					return
				}
				OK_COUNT++
				log.Printf("[%d]成功转换了:%s", OK_COUNT, entry.Name())
			}
		}
	} else {
		// 这个地方是兼容 直接传入了dat文件
		if judgeDat(pathName) {
			err := ParseAndSave(pathName, outFolder)
			if err != nil {
				return
			}
			OK_COUNT++
			log.Printf("[%d]成功转换了:%s", OK_COUNT, pathName)
		}
	}
}
