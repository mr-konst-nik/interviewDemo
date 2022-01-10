package config

import (
	"github.com/spf13/viper"
)

const (
	defaultHttpListen = "127.0.0.1:8888"
	defaultLogFile    = "./log/log.txt"
	defaultLogLevel   = "error"
)

type Config struct {
	HttpListen string
	LogFile    string
	LogLevel   string
	LogCfgFile string
}

func (c *Config) GetConfig(cfgFile string) error {
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		c.HttpListen = viper.GetString("http_listen")
		c.LogFile = viper.GetString("log_file")
		c.LogLevel = viper.GetString("log_level")
		c.LogCfgFile = viper.ConfigFileUsed()
		return nil

	}

	c.setDefaults()
	return nil
}

func (c *Config) setDefaults() {
	c.HttpListen = defaultHttpListen
	c.LogFile = defaultLogFile
	c.LogLevel = defaultLogLevel
	c.LogCfgFile = "default"
}
