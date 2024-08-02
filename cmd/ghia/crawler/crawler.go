package crawler

import (
	"github.com/ghia-xch/ghia/pkg/crawler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	l = log.WithField("component", "main")
)

func init() {

	crawler.InitFlags(crawlCommand)
}

func Init(c *cobra.Command) {

	crawlerCommand.AddCommand(crawlCommand)

	c.AddCommand(crawlerCommand)
}

var crawlerCommand = &cobra.Command{
	Use:   "crawler",
	Short: "Run and manage a network crawler",
	Long:  `Run and manage a network crawler.`,
}
