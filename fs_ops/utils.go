package fs_ops

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

/* 创建文件，并填充指定内容

Args:
    p string,  路径
	content string, 内容

Returns：
    error

*/
func createAndFillFile(p string, content string) error {
	file, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0666)
	if nil != err {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
	return nil
}

/* 找出指定目录下所有文件的路径

Args:
    root string, 指定的目录

Returns:
	[]string, 找到的文件路径
	error, 报错信息

*/
func findAllFiles(root string) ([]string, error) {
	if !isDir(root) {
		return nil, fmt.Errorf("not exist directory: %v", root)
	}
	var files []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if isFile(path) {
			files = append(files, path)
		}
		return nil
	})
	if nil != err {
		return nil, err
	}
	return files, nil
}

/* 判断给定路径是否为已存在的目录

Args:
	path string, 给定的路径

Returns:
    bool, 是否为已存在的目录

*/
func isDir(path string) bool {
	s, err := os.Stat(path)
	if nil != err {
		return false
	}
	return s.IsDir()
}

/* 判断给定路径是否为已存在的文件

Args:
	path string, 给定的路径

Returns:
    bool, 是否为已存在的文件

*/
func isFile(path string) bool {
	s, err := os.Stat(path)
	if nil != err {
		return false
	}
	return !s.IsDir()
}
