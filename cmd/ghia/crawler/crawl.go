package crawler

import (
	"github.com/ghia-xch/ghia/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var crawlCommand = &cobra.Command{
	Use:   "crawl",
	Short: "Runs an instance of the crawler",
	Long:  `Runs an instance of the crawler.`,
	Run: func(cmd *cobra.Command, args []string) {

		l.Println("-- ghia (" + viper.GetString("network") + ") - " + pkg.SemVer + " - PoST Freedom. --")

		l.Debugln("DEBUG!")

		l.Println("-- fin --")
	},
}