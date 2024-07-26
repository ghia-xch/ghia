package crawler

import "github.com/spf13/cobra"

func init() {}

func Init(c *cobra.Command) {

	crawlerCommand.AddCommand(crawlCommand)

	c.AddCommand(crawlerCommand)
}

var crawlerCommand = &cobra.Command{
	Use:   "crawler",
	Short: "Run and manage a network crawler",
	Long:  `Run and manage a network crawler.`,
}
