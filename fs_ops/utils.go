package fs_ops

import (
	"bufio"
	"os"
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
