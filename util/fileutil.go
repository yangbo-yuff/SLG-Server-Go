package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

// SaveBytesToFile 将字节数组保存到指定文件
func SaveBytesToFile(data []byte, path string) error {
	return ioutil.WriteFile(path, data, 0644)
}

// SaveStringToFile 将字符串保存到指定文件
func SaveStringToFile(content, path string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

// ReadFileToString 读取文件内容为字符串
func ReadFileToString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadFileToBytes 读取文件内容为字节数组
func ReadFileToBytes(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// CreateDir 创建目录
func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// DeleteFile 删除指定文件
func DeleteFile(path string) error {
	return os.Remove(path)
}

// Example usage
func Example() {
	err := SaveStringToFile("Hello, World!", "output.txt")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	content, err := ReadFileToString("output.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", content)
}
