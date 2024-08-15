package crawler

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/node/capability"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/full_node"
	"github.com/ghia-xch/ghia/pkg/peer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
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
				func(dec *protocol.MessageDecoder) (err error) {

					var newPeak full_node.NewPeak

					if err = newPeak.Decode(dec); err != nil {
						return err
					}

					l.Info("new peak found: ", newPeak.Height, " [", newPeak.HeaderHash.String(), "] ")
					l.Info("-- weight: ", newPeak.Weight.String())
					l.Info("-- fork: ", newPeak.ForkPointWithPreviousPeak)
					l.Info("-- unfinished block hash: ", newPeak.UnfinishedRewardBlockHash.String())

					return err
				},
			),
			protocol.Handler(
				protocol.NewTransaction,
				func(dec *protocol.MessageDecoder) (err error) {

					var newTransaction full_node.NewTransaction

					if err = newTransaction.Decode(dec); err != nil {
						return err
					}

					l.Info("new transaction found: [", newTransaction.TransactionId.String(), "]")
					l.Info("-- cost: ", newTransaction.Cost)
					l.Info("-- fees: ", newTransaction.Fees)

					var tx protocol.EncodedMessage
					var enc = protocol.NewMessageEncoder(1024)

					if tx, err = full_node.CreateRequestTransaction(newTransaction.TransactionId).Encode(enc); err != nil {
						return
					}

					l.Info("requesting transaction: [", newTransaction.TransactionId.String(), "]")

					var cb = func(dec *protocol.MessageDecoder) (err error) {

						spew.Dump(dec.Bytes())

						return err
					}

					if err = client.SendWith(tx, cb); err != nil {
						return err
					}

					return nil
				},
			),
		)

		if err = client.Open(context.Background(), 10*time.Second); err != nil {
			l.Fatalln(err)
			return
		}

		spew.Dump(client.IsCapableOf(capability.Base))

		interrupt := make(chan os.Signal, 1)

		signal.Notify(interrupt, os.Interrupt)

	CLOSER:
		for {

			select {

			case <-interrupt:

				l.Println("interrupt, closing websocket")

				if err = client.Close(); err != nil {
					l.Errorln("close:", err)
				}

			case <-client.IsClosed():

				l.Println("connection closed")

				break CLOSER
			}
		}

		l.Println("-- fin --")
	},
}
