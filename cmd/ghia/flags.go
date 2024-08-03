package main

import (
	"github.com/ghia-xch/ghia/cmd/ghia/crawler"
	"github.com/spf13/viper"
)

const (
	baseDirFlag        = "base-dir"
	networkFlag        = "network"
	configFileFlag     = "config-file"
	configSaveFlag     = "config-save"
	logsDirFlag        = "logs-dir"
	logsLevelFlag      = "logs-level"
	logsFormatFlag     = "logs-format"
	logsNoneFlag       = "logs-none"
	dataDirFlag        = "data-dir"
	nodeCAKeyFileFlag  = "node-ca-key"
	nodeCACertFileFlag = "node-ca-cert"
	nodeKeyPathFlag    = "node-tls-key-path"
	nodeCertPathFlag   = "node-tls-cert-path"
)

var (
	baseDir        string
	net            string
	configFile     string
	configSave     bool
	logsDir        string
	logsLevel      string
	logsFormat     string
	logsNone       bool
	dataDir        string
	nodeCAKeyPath  string
	nodeCACertPath string
	nodeKeyPath    string
	nodeCertPath   string
)

func init() {

	// Base

	rootCmd.PersistentFlags().StringVarP(
		&baseDir, baseDirFlag, "B", "",
		"set the base directory (default is $HOME/.ghia)",
	)

	viper.BindPFlag(baseDirFlag, rootCmd.PersistentFlags().Lookup(baseDirFlag))

	// Network //

	rootCmd.PersistentFlags().StringVarP(
		&net, networkFlag, "N", "mainnet",
		"selects the network [mainnet|testnet|simnet]",
	)

	viper.BindPFlag(networkFlag, rootCmd.PersistentFlags().Lookup(networkFlag))

	// Config //

	rootCmd.PersistentFlags().StringVarP(
		&configFile, configFileFlag, "C", "",
		"config file (default is $GHIA_BASE_DIR/$GHIA_NETWORK/config.toml)",
	)

	viper.BindPFlag(configFileFlag, rootCmd.PersistentFlags().Lookup(configFileFlag))

	rootCmd.PersistentFlags().BoolVarP(
		&configSave, configSaveFlag, "", false,
		"saves the config file with any eligible envs/flags passed,",
	)

	viper.BindPFlag(configSaveFlag, rootCmd.PersistentFlags().Lookup(configSaveFlag))

	// Logs //

	rootCmd.PersistentFlags().StringVarP(
		&logsDir, logsDirFlag, "L", "",
		"logging directory (default is $GHIA_BASE_DIR/$GHIA_NETWORK/logs)",
	)

	viper.BindPFlag(logsDirFlag, rootCmd.PersistentFlags().Lookup(logsDirFlag))

	rootCmd.PersistentFlags().StringVarP(
		&logsLevel, logsLevelFlag, "", "info",
		"sets logging level [off|fatal|error|warn|info|check|debug|trace]",
	)

	viper.BindPFlag(logsLevelFlag, rootCmd.PersistentFlags().Lookup(logsLevelFlag))

	rootCmd.PersistentFlags().StringVarP(
		&logsFormat, logsFormatFlag, "", "text",
		"sets logging format [json|text]",
	)

	viper.BindPFlag(logsFormatFlag, rootCmd.PersistentFlags().Lookup(logsFormatFlag))

	rootCmd.PersistentFlags().BoolVarP(
		&logsNone, logsNoneFlag, "", false,
		"logs only to stdout (default false)",
	)

	viper.BindPFlag(logsNoneFlag, rootCmd.PersistentFlags().Lookup(logsNoneFlag))

	// Data //

	rootCmd.PersistentFlags().StringVarP(
		&dataDir, dataDirFlag, "D", "",
		"data directory (default is $GHIA_BASE_DIR/$GHIA_NETWORK/data)",
	)

	viper.BindPFlag(dataDirFlag, rootCmd.PersistentFlags().Lookup(dataDirFlag))

	// Certificate Authority

	rootCmd.PersistentFlags().StringVarP(
		&nodeCAKeyPath, nodeCAKeyFileFlag, "", "",
		"specifies the certificate authority key used to sign the node cert (defaults to embedded key)",
	)

	viper.BindPFlag(nodeCAKeyFileFlag, rootCmd.PersistentFlags().Lookup(nodeCAKeyFileFlag))

	rootCmd.PersistentFlags().StringVarP(
		&nodeCACertPath, nodeCACertFileFlag, "", "",
		"specifies a custom certificate authority cert used to sign the node cert (defaults to embedded cert)",
	)

	viper.BindPFlag(nodeCACertFileFlag, rootCmd.PersistentFlags().Lookup(nodeCACertFileFlag))

	// Node Certificate

	rootCmd.PersistentFlags().StringVarP(
		&nodeKeyPath, nodeKeyPathFlag, "", "",
		"specifies a TLS key path for the node (defaults $GHIA_BASE_DIR/$GHIA_NETWORK/keys)",
	)

	viper.BindPFlag(nodeKeyPathFlag, rootCmd.PersistentFlags().Lookup(nodeKeyPathFlag))

	rootCmd.PersistentFlags().StringVarP(
		&nodeCertPath, nodeCertPathFlag, "", "",
		"specifies a TLS cert path for the node (defaults $GHIA_BASE_DIR/$GHIA_NETWORK/keys)",
	)

	viper.BindPFlag(nodeCertPathFlag, rootCmd.PersistentFlags().Lookup(nodeCertPathFlag))

	crawler.Init(rootCmd)
}
