package cmd

import (
	"bufio"
	"fmt"
	"github.com/Heartfilia/litetools/utils/litedir"
	"os"
	"path/filepath"
)

var KEY = [3][3]byte{
	{0x89, 0x50, 0x4e},
	{0x47, 0x49, 0x66}, // 第三个 0x46 也可以应该
	{0xff, 0xd8, 0xff},
}

type FormatResult struct {
	Value byte
	Index int
}

func ParseAndSave(datPath, saveFolder string) error {
	//fmt.Printf("%s --> %s\n", datPath, saveFolder)
	// 打开输入文件
	datRead, err := os.Open(datPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer func(datRead *os.File) {
		err := datRead.Close()
		if err != nil {
		}
	}(datRead)

	// 调用 Format 函数获取模式
	formatResult := FormatFile(datPath)
	var extension string
	switch formatResult.Index {
	case 1:
		extension = ".png"
	case 2:
		extension = ".gif"
	default:
		extension = ".jpg"
	}
	fileLength := len(datPath)
	datFileName := datPath[fileLength-36 : fileLength-4]

	// 构建输出文件路径
	outFileName := datFileName + extension
	outPathFull := filepath.Join(saveFolder, outFileName)

	if litedir.FileExists(outPathFull) {
		// 这里是已经转好了的 不需要再操作了
		return nil
	}

	// 创建输出文件
	pngWrite, err := os.Create(outPathFull)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer func(pngWrite *os.File) {
		err := pngWrite.Close()
		if err != nil {
		}
	}(pngWrite)

	// 读取输入文件并应用异或运算
	reader := bufio.NewReader(datRead)
	for {
		now, err := reader.ReadByte()
		if err != nil {
			break
		}
		newByte := now ^ formatResult.Value
		if _, err := pngWrite.Write([]byte{newByte}); err != nil {
			return fmt.Errorf("failed to write to output file: %v", err)
		}
	}

	return nil
}

func FormatFile(datPath string) FormatResult {
	// 打开文件
	file, err := os.Open(datPath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return FormatResult{Value: 0, Index: 0}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	// 使用 bufio.Reader 读取文件
	reader := bufio.NewReader(file)

	// 读取前三个字节
	firstThreeBytes, err := reader.Peek(3)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return FormatResult{Value: 0, Index: 0}
	}

	// 遍历每个模式进行异或运算
	for j, xor := range KEY {
		res := make([]byte, 3)
		for i := 0; i < 3; i++ {
			res[i] = firstThreeBytes[i] ^ xor[i]
		}
		// 检查结果是否所有字节都相同
		if res[0] == res[1] && res[1] == res[2] {
			return FormatResult{
				Value: res[0],
				Index: j + 1,
			}
		}
	}

	// 如果没有找到匹配项，返回默认值
	return FormatResult{Value: 0, Index: 0}
}
