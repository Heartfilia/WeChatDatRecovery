package cmd

import (
	"bufio"
	"fmt"
	"github.com/Heartfilia/litetools/utils/litedir"
	"log"
	"os"
	"strings"
)

type PathConf struct {
	Input  string
	Output string
}

func clearEmpty(s string) string {
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, "\r")
	return s
}

func GetChoice() PathConf {
	fmt.Println("                +-----------------------------------+")
	fmt.Println("                +         微信图片恢复工具          +")
	fmt.Println("                + 传入微信的 FileStorage 文件夹路径 +")
	fmt.Println("                +     需要指定恢复存放的文件夹      +")
	fmt.Println("                +          by: Heartfilia           +")
	fmt.Println("                +-----------------------------------+")

	var p PathConf

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("FileStorage的文件夹完整路径:")
	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败：", err)
		return p
	}
	p.Input = clearEmpty(result)

	if !litedir.FileExists(p.Input) {
		log.Panicln("输入的路径不存在")
	}

	fmt.Printf("输出的文件夹完整路径:")
	result, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败：", err)
		return p
	}
	p.Output = clearEmpty(result)

	if p.Output == "" {
		log.Fatalln("请输入 输出目录的文件夹 路径")
	}

	return p

}
