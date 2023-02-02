package cmd

import (
	"fmt"
	"os"

	"interviewDemo/internal/config"
	"interviewDemo/internal/delivery/http"
	"interviewDemo/internal/logger"

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

		cfg, err := config.NewConfig(cfgFile)
		if err != nil {
			shutDown(err)
		}

		logger, err := logger.CreateLogger(cfg.LogFile, cfg.LogLevel)
		if err != nil {
			shutDown(err)
		}

		log := logger.Sugar()
		log.Infof("Using config file: %s", cfg.LogCfgFile)
		log.Debugf("WEB server is runnig %s", cfg.HTTPListen)

		if err := http.StartSRV(cfg.HTTPListen, log); err != nil {
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
