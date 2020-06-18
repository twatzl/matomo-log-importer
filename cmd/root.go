package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"

	"github.com/twatzl/matomo-log-importer/importer"
)

var logger *logrus.Logger
var cfgFile string

const LOG_LEVEL = "logLevel"
const PARAM1 = "param1"
const PARAM2 = "param2"

var rootCmd = &cobra.Command{
	Use:   "matomo-log-importer",
	Short: "matomo-log-importer is a utility tool to read access log files of Traefik and use the Matomo API to do tracking",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &app.Config{
			Param1: viper.GetString(PARAM1),
			Param2: viper.GetBool(PARAM2),
		}

		myApp := app.New(cfg, logger)
		myApp.PrintHello()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogger)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.appname.yaml)")
	rootCmd.PersistentFlags().StringP(LOG_LEVEL, "", "info", "log level, may be one of [trace, debug, info, warn, error, fatal, panic] (default is info)")
	rootCmd.PersistentFlags().StringP(PARAM1, "", "default", "An example for a string parameter.")
	rootCmd.PersistentFlags().BoolP(PARAM2, "", false, "An example for a boolean parameter.")

	// binding flags to viper config
	_ = viper.BindPFlag(LOG_LEVEL, rootCmd.PersistentFlags().Lookup(LOG_LEVEL))
	_ = viper.BindPFlag(PARAM1, rootCmd.PersistentFlags().Lookup(PARAM1))
	_ = viper.BindPFlag(PARAM2, rootCmd.PersistentFlags().Lookup(PARAM2))

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".appname" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".appname")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		// use this code if config file is required
		//	os.Exit(1)
	} else {
		fmt.Println("config loaded")
	}
}

func initLogger() {
	levelStr := viper.GetString(LOG_LEVEL)
	level, err := logrus.ParseLevel(levelStr);
	if err != nil {
		fmt.Println("Can't parse log level:", err)
		os.Exit(1)
	}

	logger = logrus.New()
	logger.SetLevel(level)

	logger.Info("log level set to: ", level.String())
}
