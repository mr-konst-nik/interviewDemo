package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultHTTPListen = ":8080"
	defaultLogFile    = "./log.txt"
	defaultLogLevel   = "info"
)

// Config of the app
type Config struct {
	HTTPListen string
	LogFile    string
	LogLevel   string
	LogCfgFile string
}

// NewConfig create new config item
func NewConfig(cfgFile string) (*Config, error) {
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return &Config{}, fmt.Errorf("failed read config: %v", err)
		}
		return &Config{
			HTTPListen: viper.GetString("http_listen"),
			LogFile:    viper.GetString("log_file"),
			LogLevel:   viper.GetString("log_level"),
			LogCfgFile: viper.ConfigFileUsed(),
		}, nil

	}

	return &Config{
		HTTPListen: defaultHTTPListen,
		LogFile:    defaultLogFile,
		LogLevel:   defaultLogLevel,
		LogCfgFile: "default",
	}, nil
}
