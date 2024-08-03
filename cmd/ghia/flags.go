package main

import (
	"github.com/ghia-xch/ghia/cmd/ghia/crawler"
	"github.com/spf13/viper"
)

const (
	baseDirFlag      = "base-dir"
	networkFlag      = "network"
	configFileFlag   = "config-file"
	configSaveFlag   = "config-save"
	logsDirFlag      = "logs-dir"
	logsLevelFlag    = "logs-level"
	logsFormatFlag   = "logs-format"
	logsNoneFlag     = "logs-none"
	dataDirFlag      = "data-dir"
	nodeCAKeyFlag    = "node-ca-key"
	nodeCACertFlag   = "node-ca-cert"
	nodeModeFlag     = "node-tls-mode"
	nodeKeyPathFlag  = "node-tls-key-path"
	nodeCertPathFlag = "node-tls-cert-path"
)

var (
	baseDir       string
	net           string
	configFile    string
	configSave    bool
	logsDir       string
	logsLevel     string
	logsFormat    string
	logsNone      bool
	dataDir       string
	tlsCAKeyPath  string
	tlsCACertPath string
	tlsMode       string
	nodeKeyPath   string
	nodeCertPath  string
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
		&tlsCAKeyPath, nodeCAKeyFlag, "", "",
		"specifies the certificate authority key used to sign the node cert (defaults to embedded key)",
	)

	viper.BindPFlag(nodeCAKeyFlag, rootCmd.PersistentFlags().Lookup(nodeCAKeyFlag))

	rootCmd.PersistentFlags().StringVarP(
		&tlsCACertPath, nodeCACertFlag, "", "",
		"specifies a custom certificate authority cert used to sign the node cert (defaults to embedded cert)",
	)

	viper.BindPFlag(nodeCACertFlag, rootCmd.PersistentFlags().Lookup(nodeCACertFlag))

	// Node Certificate

	rootCmd.PersistentFlags().StringVarP(
		&tlsMode, nodeModeFlag, "", "public",
		"selects which embedded keypair to use [public|private]",
	)

	viper.BindPFlag(nodeModeFlag, rootCmd.PersistentFlags().Lookup(nodeModeFlag))

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
