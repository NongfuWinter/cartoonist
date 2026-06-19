package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigFile = "config/config.local.yaml"

// Config 全局配置（来自 YAML，并可被环境变量覆盖）。
type Config struct {
	HTTPPort string // HTTP 服务监听端口
	MySQLDSN string // MySQL 数据源名称
}

type yamlConfig struct {
	HTTPPort string `yaml:"http_port"`
	MySQLDSN string `yaml:"mysql_dsn"`
}

// Load 从可选 YAML 加载（默认路径 config/config.local.yaml，或通过 CONFIG_PATH），再用环境变量覆盖。
func Load() Config {
	raw, explicitPath, readErr := readYAMLBytes()
	pathDesc := explicitPathOrDefault(explicitPath)
	if readErr != nil {
		panic(fmt.Sprintf("config: %s: %v", pathDesc, readErr))
	}
	var yc yamlConfig
	if len(raw) > 0 {
		if err := yaml.Unmarshal(raw, &yc); err != nil {
			panic(fmt.Sprintf("config: parse YAML %s: %v", pathDesc, err))
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = yc.HTTPPort
	}
	if port == "" {
		port = "8080"
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = yc.MySQLDSN
	}
	if dsn == "" {
		dsn = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	}

	return Config{
		HTTPPort: port,
		MySQLDSN: dsn,
	}
}

func readYAMLBytes() (data []byte, explicitPath bool, err error) {
	path := os.Getenv("CONFIG_PATH")
	if path != "" {
		raw, readErr := os.ReadFile(path)
		return raw, true, readErr
	}

	raw, readErr := os.ReadFile(defaultConfigFile)
	if readErr != nil {
		if errors.Is(readErr, os.ErrNotExist) {
			return nil, false, nil
		}
		return nil, false, readErr
	}
	return raw, false, nil
}

func explicitPathOrDefault(explicitPath bool) string {
	if explicitPath {
		return os.Getenv("CONFIG_PATH")
	}
	return defaultConfigFile
}

func (c Config) ListenAddr() string {
	return ":" + c.HTTPPort
}
