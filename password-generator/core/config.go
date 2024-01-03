package core

import (
	"github.com/spf13/viper"
)

// Config 结构用于保存配置信息
type Config struct {
	Filename  string
	Passwords map[string]string
}

// NewConfig 创建新的配置实例
func NewConfig(filename string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)

	return &Config{
		Filename:  filename,
		Passwords: make(map[string]string),
	}
}

// Load 从配置文件中加载配置
func (c *Config) Load() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.UnmarshalKey("passwords", &c.Passwords)
}

// Save 将配置保存到配置文件
func (c *Config) Save() error {
	viper.Set("passwords", c.Passwords)
	return viper.WriteConfigAs(c.Filename)
}
