package config

import (
	"errors"
	"os"

	"github.com/wpx1990/go-base/pkg/log"
	"gopkg.in/yaml.v2"
)

func GetConfig(file_paths []string, config interface{}) error {

	var err error
	var file *os.File
	i := 0
	for ; i < len(file_paths); i++ {
		// 打开yaml文件
		file, err = os.Open(file_paths[i])
		if err == nil {
			log.Info("Succeed to open config file(%s).", file_paths[i])
			break
		}
	}

	if i == len(file_paths) {
		log.Error("Failed to open config file(%v).", file_paths)
		return errors.New("failed to open config file")
	}

	defer file.Close()

	// 创建解析器
	decoder := yaml.NewDecoder(file)

	// 解析 YAML 数据
	err = decoder.Decode(config)
	if err != nil {
		log.Error("Failed to decode config file(%s), err(%v).", file_paths[i], err)
		return err
	}

	return nil
}
