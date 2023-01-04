package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultHTTPListen = "127.0.0.1:8888"
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

// GetConfig of the app
func (c *Config) GetConfig(cfgFile string) error {
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("failed read config: %v", err)
		}

		c.HTTPListen = viper.GetString("http_listen")
		c.LogFile = viper.GetString("log_file")
		c.LogLevel = viper.GetString("log_level")
		c.LogCfgFile = viper.ConfigFileUsed()
		return nil

	}

	c.setDefaults()
	return nil
}

func (c *Config) setDefaults() {
	c.HTTPListen = defaultHTTPListen
	c.LogFile = defaultLogFile
	c.LogLevel = defaultLogLevel
	c.LogCfgFile = "default"
}
