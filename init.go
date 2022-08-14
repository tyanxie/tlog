package tlog

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configFilename = "config.yaml"

func init() {
	var configs []*Config

	// 读取配置文件，如果文件不存在则使用默认配置，否则如果失败就panic
	configFile, err := os.ReadFile(configFilename)
	if err != nil {
		if os.IsNotExist(err) {
			configs = defaultConfigs
		} else {
			panic(fmt.Errorf("init tlog failed: open config file failed: %w", err))
		}
	} else {
		var wrapper *configWrapper
		if err := yaml.Unmarshal(configFile, &wrapper); err != nil {
			panic(fmt.Errorf("init tlog failed: parse config file failed: %w", err))
		}
		if wrapper == nil { // 如果配置文件中没有log项，则使用默认配置
			configs = defaultConfigs
		} else {
			configs = wrapper.Log
		}
	}
	// 如果配置文件中有log项，但其下内容为空，则使用默认配置
	if configs == nil || len(configs) == 0 {
		configs = defaultConfigs
	}

	// 创建默认日志器，失败时panic
	defaultLogger, err = NewWithCallerSkip(1, configs...)
	if err != nil {
		panic(fmt.Errorf("init tlog failed: %w", err))
	}
}
