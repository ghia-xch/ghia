package crawler

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/peer"
	"github.com/ghia-xch/ghia/pkg/protocol"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var crawlCommand = &cobra.Command{
	Use:   "crawl",
	Short: "Runs an instance of the crawler",
	Long:  `Runs an instance of the crawler.`,
	Run: func(cmd *cobra.Command, args []string) {

		l.Println("-- ghia (" + viper.GetString("network") + ") - " + pkg.SemVer + " - PoST Freedom. --")

		var err error
		var client *node.Client

		client = node.NewClient(peer.NewPeerInfo("203.184.53.208", 8444))

		client.Handle(
			protocol.Handler(
				protocol.NewPeak,
				func(em protocol.EncodedMessage) (err error) {

					l.Infoln("New Peak found!")

					spew.Dump(em)

					return err
				},
			),
		)

		if err = client.Open(context.Background(), 10*time.Second); err != nil {
			l.Fatalln(err)
			return
		}

		//client.SendWith()

		l.Println("-- fin --")
	},
}
