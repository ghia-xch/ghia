package crawler

import (
	"context"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/node/capability"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
	"github.com/ghia-xch/ghia/pkg/node/protocol/full_node"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
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

		var err error

		l.Println("-- ghia (" + viper.GetString("network") + ") - " + pkg.SemVer + " - PoST Freedom. --")

		a, _ := hex.DecodeString("16069bd1d0c581e0014f48aa828209a8a351d7dd069999766714fefdfc07fe95")

		b := protocol.Hash(a)

		nub := full_node.NewUnfinishedBlock2{
			UnfinishedRewardHash: protocol.Hash(a),
			FoliageHash:          &b,
		}

		spew.Dump(codec.Encode(nil, &nub))

		rq := full_node.RequestTransaction{
			TransactionId: protocol.Hash(a),
		}

		spew.Dump(codec.Encode(nil, &rq))

		os.Exit(0)

		var client *node.Client

		client = node.NewClient(peer.NewPeerInfo("192.168.8.117", 8444))

		client.Handle(
			protocol.Handler(
				protocol.NewPeak,
				func(em message.EncodedMessage) (err error) {

					var np full_node.NewPeak

					if err = codec.Decode(&np, em); err != nil {
						return err
					}

					l.Info("new peak found: ", np.Height, " [", np.HeaderHash.String(), "] ")
					l.Info("-- weight: ", np.Weight.String())
					l.Info("-- fork: ", np.ForkPointWithPreviousPeak)
					l.Info("-- unfinished block hash: ", np.UnfinishedRewardBlockHash.String())

					return err
				},
			),
			//protocol.Handler(
			//	protocol.NewTransaction,
			//	func(em message.EncodedMessage) (err error) {
			//
			//		var nt full_node.NewTransaction
			//
			//		if err = codec.Decode(&nt, em); err != nil {
			//			return err
			//		}
			//
			//		l.Info("new transaction found: [", nt.TransactionId.String(), "]")
			//		l.Info("-- cost: ", nt.Cost)
			//		l.Info("-- fees: ", nt.Fees)
			//
			//		l.Info("requesting transaction: [", nt.TransactionId.String(), "]")
			//
			//		var rq full_node.RequestTransaction
			//		rq.TransactionId = nt.TransactionId
			//
			//		if em, err = codec.Encode(nil, &rq); err != nil {
			//			return err
			//		}
			//
			//		if err = client.SendWith(em,
			//			func(em message.EncodedMessage) (err error) {
			//
			//				spew.Dump(em)
			//
			//				return err
			//			},
			//		); err != nil {
			//			return err
			//		}
			//
			//		return nil
			//	},
			//),
			//protocol.Handler(
			//	protocol.NewSignagePointOrEndOfSubSlot,
			//	func(dec *message.MessageDecoder) (err error) {
			//
			//		spew.Dump(dec.Type())
			//
			//		return nil
			//	},
			//),
		)

		if err = client.Open(context.Background(), 10*time.Second); err != nil {
			l.Fatalln(err)
			return
		}

		spew.Dump(client.IsCapableOf(capability.Base))

		var em message.EncodedMessage
		var peersReq full_node.RequestPeers

		if em, err = codec.Encode(nil, &peersReq); err != nil {
			l.Fatalln(err)
			os.Exit(1)
		}

		err = client.SendWith(em, func(rem message.EncodedMessage) (err error) {

			spew.Dump(rem)

			var respondPeers full_node.RespondPeers

			if err = codec.Decode(&respondPeers, rem); err != nil {
				l.Fatalln(err)
				os.Exit(1)
			}

			spew.Dump(respondPeers)

			return nil
		})

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
