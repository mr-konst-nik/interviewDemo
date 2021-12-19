package config

import (
	"fmt"

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
}

func (c *Config) GetConfig(cfgFile string) error {
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		fmt.Println("Using config file:", viper.ConfigFileUsed())
		c.HttpListen = viper.GetString("http_listen")
		c.LogFile = viper.GetString("log_file")
		c.LogLevel = viper.GetString("log_level")
		return nil

	}

	c.setDefaults()
	return nil
}

func (c *Config) setDefaults() {
	c.HttpListen = defaultHttpListen
	c.LogFile = defaultLogFile
	c.LogLevel = defaultLogLevel
	fmt.Println("Set default config")
}
