package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName     string         `json:"app_name"`
	AppMode     string         `json:"app_mode"`
	AppHost     string         `json:"app_host"`
	AppPort     string         `json:"app_port"`
	Sms         SmsConfig      `json:"sms"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config"`
}
type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}
type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

// ParseConfig 解析配置文件
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取json文件的内容
	reader := bufio.NewReader(file)
	// 生成一个json解析器
	decoder := json.NewDecoder(reader)
	//将json文件的内容解析到全局变量_cfg
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
