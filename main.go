package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	currentDir, err := os.Getwd() // 獲取目前資料夾路徑
	if err != nil {
		fmt.Println("無法獲取目前資料夾：", err)
		return
	}

	fmt.Println("目前資料夾：", currentDir)
	fmt.Println("檔案結構：")
	err = printDirectoryStructure(currentDir, 0) // 列印資料夾結構
	if err != nil {
		fmt.Println("無法生成檔案結構圖：", err)
	}

	fmt.Scanln() // 等待使用者輸入
}

// 列印資料夾結構
func printDirectoryStructure(dirPath string, indentLevel int) error {
	files, err := ioutil.ReadDir(dirPath) // 讀取資料夾內的檔案和子資料夾
	if err != nil {
		return err
	}

	indent := strings.Repeat("  ", indentLevel) // 根據縮排層級生成縮排字串

	for _, file := range files {
		fileName := file.Name() // 獲取檔案名稱
		if file.IsDir() {
			fmt.Println(indent + "├── " + fileName + "/")             // 列印資料夾名稱
			subDirPath := filepath.Join(dirPath, fileName)            // 子資料夾的路徑
			err := printDirectoryStructure(subDirPath, indentLevel+1) // 遞迴處理子資料夾
			if err != nil {
				fmt.Println("無法生成檔案結構圖：", err)
			}
		} else {
			fmt.Println(indent + "├── " + fileName) // 列印檔案名稱
		}
	}

	return nil
}
