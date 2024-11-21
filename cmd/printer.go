package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PathConf struct {
	Input  string
	Output string
}

func GetChoice() PathConf {
	fmt.Println("                +-----------------------------------+")
	fmt.Println("                +         微信图片恢复工具          +")
	fmt.Println("                + 传入微信的 FileStorage 文件夹路径 +")
	fmt.Println("                +     需要指定恢复存放的文件夹      +")
	fmt.Println("                +-----------------------------------+")

	var p PathConf

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("FileStorage的文件夹完整路径:")
	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败：", err)
		return p
	}
	p.Input = strings.Trim(result, "\n")

	fmt.Printf("输出的文件夹完整路径:")
	result, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败：", err)
		return p
	}
	p.Output = strings.Trim(result, "\n")

	return p

}
