package main

import (
	"errors"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/ghia-xch/ghia/cmd/ghia/crawler"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/protocol/network"
	"github.com/ghia-xch/ghia/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var ghiaTxt = `ghia (` + pkg.SemVer + `) - PoST Freedom.

Ghia is a tool suite for interacting with the chia net.
`

var (
	cfgFile     string
	cfgSave     bool
	logsDir     string
	logsLevel   string
	dataDir     string
	net         string
	tlsMode     string
	tlsKeyPath  string
	tlsCertPath string
)

var (
	l = log.WithField("component", "main")
)

func init() {

	viper.SetEnvPrefix("GHIA")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	cobra.OnInitialize(initNetwork)
	cobra.OnInitialize(initKeys)
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogging)
	cobra.OnInitialize(initData)

	cobra.OnFinalize(persistConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config-file", "C", "", "config file (default is $HOME/.ghia/config.toml)")
	rootCmd.PersistentFlags().BoolVarP(&cfgSave, "config-save", "", false, "saves the config file with any eligible envs/flags passed")
	rootCmd.PersistentFlags().StringVarP(&logsDir, "logs-dir", "L", "", "logging directory (default is $HOME/.ghia/logs)")
	rootCmd.PersistentFlags().StringVarP(&logsLevel, "logs-level", "", "info", "sets logging level [off|fatal|error|warn|info|check|debug|trace]")
	rootCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "D", "", "data directory (default is $HOME/.ghia/data)")
	rootCmd.PersistentFlags().StringVarP(&net, "network", "N", "mainnet", "selects the network [mainnet|testnet|simnet]")

	rootCmd.PersistentFlags().StringVarP(&tlsKeyPath, "node-tls-mode", "", "public", "selects which embedded keypair to use [public|private]")
	rootCmd.PersistentFlags().StringVarP(&tlsKeyPath, "node-tls-key-path", "", "", "specifies a custom TLS key path for interacting with nodes (defaults to embedded key)")
	rootCmd.PersistentFlags().StringVarP(&tlsCertPath, "node-tls-cert-path", "", "", "specifies a custom TLS cert path for interacting with nodes (defaults to embedded cert)")

	viper.BindPFlag("logs-dir", rootCmd.PersistentFlags().Lookup("logs-dir"))
	viper.BindPFlag("logs-level", rootCmd.PersistentFlags().Lookup("logs-level"))

	viper.BindPFlag("data-dir", rootCmd.PersistentFlags().Lookup("data-dir"))

	viper.BindPFlag("network", rootCmd.PersistentFlags().Lookup("network"))

	viper.BindPFlag("node-tls-mode", rootCmd.PersistentFlags().Lookup("node-tls-mode"))
	viper.BindPFlag("node-tls-key-path", rootCmd.PersistentFlags().Lookup("node-tls-key-path"))
	viper.BindPFlag("node-tls-cert-path", rootCmd.PersistentFlags().Lookup("node-tls-cert-path"))

	crawler.Init(rootCmd)
}

var N *network.Network

func initNetwork() {

	var err error

	if N, err = network.Select(viper.GetString("network")); err != nil {

		l.Fatal(err)

		os.Exit(1)
	}
}

func initKeys() {

	//var err error

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

	log.SetFormatter(
		&nested.Formatter{
			HideKeys:        true,
			FieldsOrder:     []string{"component", "category"},
			TimestampFormat: "2006-01-02 15:04:05",
			CallerFirst:     true,
		},
	)

	log.SetLevel(
		util.GetLogLevel(
			viper.GetString("logs-level"), log.DebugLevel,
		),
	)

	if log.GetLevel() == log.DebugLevel {
		log.SetReportCaller(true)
	}
}

func initConfig() {

	if cfgFile == "" {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		cfgFile = home + "/.ghia/" + N.String.String() + "/config.toml"
	}

	viper.SetConfigFile(cfgFile)

	if _, err := os.Stat(cfgFile); errors.Is(err, os.ErrNotExist) {
		return
	}

	if err := viper.ReadInConfig(); err != nil {

		l.Errorln("failed to read config file:", err)

		os.Exit(1)
	}

}

func persistConfig() {

	if !cfgSave {
		return
	}

	if err := viper.WriteConfig(); err != nil {

		l.Errorln("failed to save config file:", err)

		os.Exit(1)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "ghia",
		Short: "PoST Freedom.",
		Long:  ghiaTxt,
	}
)

func main() {

	var err error

	if err = rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
