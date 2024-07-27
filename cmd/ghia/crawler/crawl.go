package crawler

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/protocol/network"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	l = log.WithField("component", "main")
)

var crawlCommand = &cobra.Command{
	Use:   "crawl",
	Short: "Runs an instance of the crawler",
	Long:  `Runs an instance of the crawler.`,
	Run: func(cmd *cobra.Command, args []string) {

		l.Println("-- ghia (" + viper.GetString("network") + ") - " + pkg.SemVer + " - PoST Freedom. --")

		var err error
		var net *network.Network

		if net, err = network.Select(viper.GetString("network")); err != nil {

			l.Fatal(err)

			os.Exit(1)
		}

		spew.Dump(net)

		l.Debugln("DEBUG!")

		l.Println("-- fin --")
	},
}
