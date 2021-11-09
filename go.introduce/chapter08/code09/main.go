package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CreateDir 根据当前日期来创建文件夹
func createDir(Path string) string {
	//Format对数字是有要求的,2006/01/02 15:04:05
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		//先创建文件夹
		os.MkdirAll(folderPath, os.ModePerm)
		//修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}
func main() {
	//获取当前绝对路径
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	path = createDir(path)
	fmt.Println(path)
	os.MkdirAll(filepath.Join(path, "新文件夹/dic"), os.ModePerm)
	//重命名或者移动目录
	os.Rename(filepath.Join(path, "新文件夹"), filepath.Join(path, "newfolder"))
	//删除目录,如果有子目录则不删除
	os.Remove(filepath.Join(path, "newfolder/dic"))
	//os.RemoveAll(filepath.Join(path, "newfolder"))
}
