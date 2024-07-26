package main

import (
	"errors"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var ghiaTxt = `ghia (` + pkg.SemVer + `) - PoST Freedom.

Ghia is a tool suite for interacting with the chia network.
`

var (
	cfgFile   string
	cfgSave   bool
	logsDir   string
	logsLevel string
	dataDir   string
	network   string

	rootCmd = &cobra.Command{
		Use:   "ghia",
		Short: "PoST Freedom.",
		Long:  ghiaTxt,
	}
)

func init() {

	viper.SetEnvPrefix("GHIA")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogging)
	cobra.OnInitialize(initData)

	cobra.OnFinalize(persistConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config-file", "C", "", "config file (default is $HOME/.ghia/config.toml)")
	rootCmd.PersistentFlags().BoolVarP(&cfgSave, "config-save", "", false, "saves the config file with any eligible envs/flags passed")
	rootCmd.PersistentFlags().StringVarP(&logsDir, "logs-dir", "L", "", "logging directory (default is $HOME/.ghia/logs)")
	rootCmd.PersistentFlags().StringVarP(&logsLevel, "logs-level", "", "info", "set logging level  off|fatal|error|warn|info|check|debug|trace")
	rootCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "D", "", "data directory (default is $HOME/.ghia/data)")
	rootCmd.PersistentFlags().StringVarP(&network, "network", "N", "mainnet", "selects the network  mainnet|testnet|simnet")

	viper.BindPFlag("logs-dir", rootCmd.PersistentFlags().Lookup("logs-dir"))
	viper.BindPFlag("logs-level", rootCmd.PersistentFlags().Lookup("logs-level"))
	viper.BindPFlag("data-dir", rootCmd.PersistentFlags().Lookup("data-dir"))
	viper.BindPFlag("network", rootCmd.PersistentFlags().Lookup("network"))
}

func initData() {

	if viper.GetString("data-dir") == "" {
		home, err := os.UserHomeDir()

		cobra.CheckErr(err)

		viper.Set("data-dir", home+"/.ghia/data")
	}

}

func initLogging() {

	if logsDir == "" {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		logsDir = home + "/.ghia/logs"
	}

	//log2.SetLogLevel(log2.GetLevelByString(viper.GetString("logs-level"), log2.Debug))
	//log2.CodeLocations(false)
	//
	//if log2.GetLogLevel() == log2.Debug {
	//	log2.CodeLocations(true)
	//}
}

func initConfig() {

	if cfgFile == "" {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		cfgFile = home + "/.ghia/config.toml"
	}

	viper.SetConfigFile(cfgFile)

	if _, err := os.Stat(cfgFile); errors.Is(err, os.ErrNotExist) {
		return
	}

	if err := viper.ReadInConfig(); err != nil {
		//log.E.Ln("failed to read config file:", err)
		os.Exit(1)
	}

}

func persistConfig() {

	if !cfgSave {
		return
	}

	if err := viper.WriteConfig(); err != nil {
		//log.E.Ln("failed to save config file:", err)
		os.Exit(1)
	}
}

func main() {

	var err error

	if err = rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
