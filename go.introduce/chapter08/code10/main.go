package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//遍历目录和文件
func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	return nil
}

func main() {
	//遍历当前目录
	if root, err := os.Getwd(); err == nil {
		err := filepath.Walk(root, visit)
		if err != nil {
			fmt.Printf("%v\n", err)
		}

	}

}
