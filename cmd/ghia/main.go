package main

import (
	"github.com/ghia-xch/ghia/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	l = log.WithField("component", "main")
)

var rootCmd = &cobra.Command{
	Use:   "ghia",
	Short: "PoST Freedom.",
	Long: `ghia (` + pkg.SemVer + `) - PoST Freedom.

Ghia is a tool suite for interacting with the chia net.`,
}

func main() {

	var err error

	if err = rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
