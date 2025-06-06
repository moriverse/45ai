package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	JWT       JWTConfig       `mapstructure:"jwt"`
	WeChat    WeChatConfig    `mapstructure:"wechat"`
	AIService AIServiceConfig `mapstructure:"ai_service"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type JWTConfig struct {
	SecretKey      string `mapstructure:"secret_key"`
	ExpiresInHours int    `mapstructure:"expires_in_hours"`
}

type WeChatConfig struct {
	AppID          string `mapstructure:"app_id"`
	MchID          string `mapstructure:"mch_id"`
	APIv3Key       string `mapstructure:"api_v3_key"`
	PrivateKeyPath string `mapstructure:"private_key_path"`
}

type AIServiceConfig struct {
	APIKey      string `mapstructure:"api_key"`
	APIEndpoint string `mapstructure:"api_endpoint"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
