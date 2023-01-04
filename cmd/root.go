package cmd

import (
	"fmt"
	"os"

	"interviewDemo/internal/config"
	"interviewDemo/internal/logger"
	"interviewDemo/internal/server"

	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to config file")
}

var rootCmd = &cobra.Command{
	Use:   "web server",
	Short: "test web server",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		if err := cfg.GetConfig(cfgFile); err != nil {
			shutDown(err)
		}
		if err := logger.CreateLogger(cfg.LogFile, cfg.LogLevel); err != nil {
			shutDown(err)
		}
		logger.Log.Sugar().Infof("Using config file: %s", cfg.LogCfgFile)
		logger.Log.Sugar().Debugf("WEB server is runnig %s", cfg.HTTPListen)
		if err := server.StartSRV(cfg.HTTPListen); err != nil {
			shutDown(err)
		}
	},
}

// Execute the app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		shutDown(err)
	}
}

func shutDown(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
