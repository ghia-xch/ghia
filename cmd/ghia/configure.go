package main

import (
	"crypto/tls"
	"errors"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/protocol/network"
	"github.com/ghia-xch/ghia/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init() {

	viper.SetEnvPrefix("GHIA")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	cobra.OnInitialize(initBase)
	cobra.OnInitialize(initNetwork)
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogging)
	cobra.OnInitialize(initData)
	cobra.OnInitialize(initKeys)

	cobra.OnFinalize(persistConfig)
}

func initBase() {

	var err error
	var homeDir string

	if viper.GetString(baseDirFlag) == "" {

		if homeDir, err = os.UserHomeDir(); err != nil {
			cobra.CheckErr(err)
		}

		viper.Set(baseDirFlag, homeDir+"/.ghia")
	}

	if err = os.MkdirAll(viper.GetString(baseDirFlag), 0755); err != nil {
		cobra.CheckErr(err)
	}
}

var N *network.Network

func initNetwork() {

	var err error

	if N, err = network.Select(viper.GetString("network")); err != nil {

		l.Fatal(err)

		os.Exit(1)
	}
}

func initConfig() {

	var err error
	var configBase string

	if viper.GetString(configFileFlag) == "" {

		configBase = viper.GetString(baseDirFlag) + "/" + N.String.String()

		if err = os.MkdirAll(configBase, 0755); err != nil {
			cobra.CheckErr(err)
		}

		viper.Set(configFileFlag, configBase+"/config.toml")
	}

	viper.SetConfigFile(viper.GetString(configFileFlag))

	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		return
	}

	if err := viper.ReadInConfig(); err != nil {

		l.Errorln("failed to read config file:", err)

		os.Exit(1)
	}

}

func persistConfig() {

	if !viper.GetBool(configSaveFlag) {
		return
	}

	if err := viper.WriteConfig(); err != nil {

		l.Errorln("failed to save config file:", err)

		os.Exit(1)
	}
}

func initLogging() {

	log.SetLevel(
		util.GetLogLevel(
			viper.GetString(logsLevelFlag), log.DebugLevel,
		),
	)

	if log.GetLevel() == log.DebugLevel {
		log.SetReportCaller(true)
	}

	log.SetFormatter(
		&nested.Formatter{
			HideKeys:        true,
			FieldsOrder:     []string{"component", "category"},
			TimestampFormat: "2006-01-02 15:04:05",
			CallerFirst:     true,
		},
	)

	if viper.GetString(logsFormatFlag) == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if viper.GetBool(logsNoneFlag) {
		return
	}

	var err error

	if viper.GetString(logsDirFlag) == "" {
		viper.Set(logsDirFlag, viper.GetString(baseDirFlag)+"/"+N.String.String()+"/logs")
	}

	if err = os.MkdirAll(viper.GetString(logsDirFlag), 0755); err != nil {
		cobra.CheckErr(err)
	}

	var file *os.File

	if file, err = os.OpenFile(viper.GetString(logsDirFlag)+"/ghia.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
		cobra.CheckErr(err)
	}

	log.SetOutput(file)
}

func initData() {

	var err error

	if viper.GetString(dataDirFlag) == "" {
		viper.Set(dataDirFlag, viper.GetString(baseDirFlag)+"/"+N.String.String()+"/data")
	}

	if err = os.MkdirAll(viper.GetString(dataDirFlag), 0755); err != nil {
		cobra.CheckErr(err)
	}
}

func initKeys() {

	var err error
	var cert tls.Certificate

	if cert, err = tls.X509KeyPair([]byte(node.DefaultCACertificate), []byte(node.DefaultCAKey)); err != nil {
		cobra.CheckErr(err)
	}

	if viper.GetString(nodeCAKeyFileFlag) != "" {

		cert, err = tls.LoadX509KeyPair(viper.GetString(nodeCAKeyFileFlag), viper.GetString(nodeCACertFileFlag))

		if err != nil {
			cobra.CheckErr(err)
		}
	}

	spew.Dump(cert)

	//var err error

}
