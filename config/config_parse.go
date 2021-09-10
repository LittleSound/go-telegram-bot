package config

import (
	"go-bot/assertion"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config 配置文件的格式
type Config struct {
	BotToken                   string `yaml:"botToken"`
	Debug                      bool   `yaml:"debug"`
	IgnoreThePrelaunchMessages bool   `yaml:"ignoreThePrelaunchMessages"`
}

// InitConfig 初始化配置文件
func InitConfig(configPath string) Config {
	var err error
	var config Config

	data, err := ioutil.ReadFile(configPath)
	assertion.Panic(err)

	err = yaml.Unmarshal([]byte(data), &config)
	assertion.Panic(err)
	logrus.Println(config)
	return config
}
