package crawler

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	introducersFlag = "introducer"
	maxCrawlersFlag = "crawlers-max"
)

var (
	introducers []string
	maxCrawlers uint8
)

func InitFlags(cmd *cobra.Command) {

	cmd.PersistentFlags().StringSliceVarP(
		&introducers,
		introducersFlag,
		"",
		[]string{
			"/dns4/dns-introducer.chia.net/tcp/8445",
			"/dns6/dns-introducer.chia.net/tcp/8445",
		},
		"adds an additional introducer connection",
	)

	viper.BindPFlag(introducersFlag, cmd.PersistentFlags().Lookup(introducersFlag))

	cmd.PersistentFlags().Uint8VarP(
		&maxCrawlers,
		maxCrawlersFlag,
		"",
		3,
		"maximum number of crawling threads",
	)

	viper.BindPFlag(maxCrawlersFlag, cmd.PersistentFlags().Lookup(maxCrawlersFlag))
}
