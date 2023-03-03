package fs_ops

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jielyu/river/models"
)

func ParseToml(tomlPath string) (models.TOMLConfig, error) {
	config := models.TOMLConfig{}
	// 检查配置文件是否存在
	if _, err := os.Stat(tomlPath); err != nil {
		return config, fmt.Errorf("failed to read River.toml, e:%v", err)
	}
	_, err := toml.DecodeFile(tomlPath, &config)
	if err != nil {
		log.Fatal("failed to parse River.toml")
	}
	return config, nil
}
