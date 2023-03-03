package fs_ops

import (
	"fmt"

	"github.com/jielyu/river/models"
)

func BuildProject(cConfig *models.CommandConfig, tConfig *models.TOMLConfig) error {
	// 读取toml文件的配置
	//执行编译
	fmt.Printf("release: %v\r\n", cConfig.Release)
	return nil
}
