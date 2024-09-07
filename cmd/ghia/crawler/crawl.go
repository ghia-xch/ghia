package crawler

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/node/capability"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
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

		client = node.NewClient(peer.NewPeerInfo("192.168.8.117", 8444))

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

					l.Info("requesting transaction: [", newTransaction.TransactionId.String(), "]")

					spew.Dump(codec.Encode(
						nil,
						&full_node.RequestTransaction{
							TransactionId: newTransaction.TransactionId,
						},
					))

					spew.Dump(codec.Encode(nil, &newTransaction))

					//if err = client.SendWith(
					//	full_node.CreateRequestTransaction(newTransaction.TransactionId),
					//	func(dec *protocol.MessageDecoder) (err error) {
					//
					//		spew.Dump(dec.Type())
					//
					//		return err
					//	},
					//); err != nil {
					//	return err
					//}

					return nil
				},
			),
			protocol.Handler(
				protocol.NewSignagePointOrEndOfSubSlot,
				func(dec *protocol.MessageDecoder) (err error) {

					spew.Dump(dec.Type())

					return nil
				},
			),
		)

		if err = client.Open(context.Background(), 10*time.Second); err != nil {
			l.Fatalln(err)
			return
		}

		spew.Dump(client.IsCapableOf(capability.Base))

		//h1 := sha256.New()
		//h1.Write([]byte("fdhjkdshfjkdshff"))
		//r1 := make([]byte, 32)
		//r1 = h1.Sum(nil)

		// 0x2511ac63199f675412ec2db94f8b89802950f358c9ab4b7b86003f8c7dd7ea38

		//h2 := sha256.New()
		//h2.Write([]byte("fdgsdbbgfbggf"))
		//r2 := make([]byte, 32)
		//r2 = h2.Sum(nil)
		//
		//r3, _ := hex.DecodeString("16069bd1d0c581e0014f48aa828209a8a351d7dd069999766714fefdfc07fe95")
		//
		//spew.Dump(r3)

		//c := coin.NewCoin(protocol.Hash(r3), protocol.Hash(r2), 1)
		//
		//spew.Dump(c.Id())

		//if err = client.SendWith(
		//	full_node.CreateRequestPeers(),
		//	func(dec *protocol.MessageDecoder) (err error) {
		//
		//		l.Infoln("lol")
		//
		//		var rp full_node.RespondPeers
		//
		//		if err = rp.Decode(dec); err != nil {
		//
		//			l.Fatalln(err)
		//
		//			return err
		//		}
		//
		//		spew.Dump(rp)
		//
		//		return nil
		//	},
		//); err != nil {
		//	l.Fatalln(err)
		//}

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
